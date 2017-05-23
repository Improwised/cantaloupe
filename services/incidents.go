package services

import (
	"encoding/json"
	"time"
)

// IncidentInfo : incident insformation.
type IncidentInfo struct {
	Id                 int
	Component_id       int
	Serial_no          string
	Title              string
	Recorder           string
	Component          string
	Warranty_timestamp time.Time
	Warranty_till      string
	Machine            *string
	Description        string
	Status             string
	Resolved_at        *time.Time
}

// AddIncident : add new incident.
func AddIncident(data string) {
	sess := SetupDB()
	i := IncidentInfo{}
	err := json.Unmarshal([]byte(data), &i) //converting JSON Object to GO structure ...
	CheckErr(err)
	i.Status = "active"

	_, err2 := sess.InsertInto("incidents").
		Columns("component_id", "title", "recorder", "description", "status").
		Record(i).
		Exec()
	CheckErr(err2)

	_, err3 := sess.Update("components").
		Set("active", false).
		Where("id = ?", i.Component_id).
		Exec()
	CheckErr(err3)

	_, err4 := sess.Update("machine_components").
		Set("deleted_at", "NOW()").
		Where("component_id = ?", i.Component_id).
		Exec()
	CheckErr(err4)
}

// EditIncident : edit incident.
func EditIncident(incidentId int, componentId string, recorder string, title string, description string) {
	sess := SetupDB()
	query := sess.Update("incidents")
	if componentId != "" {
		query.Set("component_id", componentId)
	}
	query.Set("title", title).
		Set("description", description).
		Set("recorder", recorder).
		Where("id = ?", incidentId).
		Exec()
}

// DeleteIncident : remove incident.
func DeleteIncident(incidentId int) {
	sess := SetupDB()

	componentId, err := sess.Select("component_id").
		From("incidents").
		Where("id = ?", incidentId).
		ReturnInt64()
	CheckErr(err)

	_, err1 := sess.Update("components").
		Set("active", true).
		Where("id = ?", componentId).
		Exec()
	CheckErr(err1)

	_, err2 := sess.DeleteFrom("incidents").
		Where("id = ?", incidentId).
		Exec()
	CheckErr(err2)
}

// DisplayIncidents :  all incidents.
func DisplayIncidents(all string) []byte {
	sess := SetupDB()
	incidentInfo := []IncidentInfo{}

	query := sess.Select(`i.id,
                        i.title,
                        i.description,
                        i.status,
                        i.recorder,
                        components.name AS Component,
                        components.serial_no,
                        components.warranty_till AS Warranty_timestamp`).
		From("incidents i").
		LeftJoin("components", "i.component_id = components.id")

	if all == "false" {
		query.Where("i.status = ?", "active")
	}

	query.LoadStruct(&incidentInfo)

	//extract only date from timestamp============================================
	for i := 0; i < len(incidentInfo); i += 1 {
		t := incidentInfo[i].Warranty_timestamp
		incidentInfo[i].Warranty_till = t.Format("2006-01-02")
	}
	//============================================================================

	b, err := json.Marshal(incidentInfo)
	CheckErr(err)
	return b
}

// DisplayIncident : display one incident.
func DisplayIncident(incidentId int) []byte {
	sess := SetupDB()
	incidentInfo := IncidentInfo{}
	sess.Select("id, component_id, title, description, status, resolved_at").
		From("incidents").
		Where("id = ?", incidentId).
		LoadStruct(&incidentInfo)
	b, err := json.Marshal(incidentInfo)
	CheckErr(err)
	return b
}

// IncidentUpdate : updates on incident.
type IncidentUpdate struct {
	Id            int
	Updated_by    string
	Created_at    time.Time
	Resolved_Date string
	ComponentId   int64
	IncidentId    int
	Description   string
}

// IncidentUpdates : updates on incident.
func IncidentUpdates(incidentId int, resolvedBy string, description string, isResolved string) {
	sess := SetupDB()

	//select component id on which incident happen ....===========================
	id, err1 := sess.Select("component_id").
		From("incidents").
		Where("id = ?", incidentId).
		ReturnInt64()
	CheckErr(err1)
	//============================================================================

	//insert record into incident_updates table ...================================
	var incident IncidentUpdate
	incident.ComponentId = id
	incident.IncidentId = incidentId
	incident.Updated_by = resolvedBy
	incident.Description = description

	_, err := sess.InsertInto("incident_update").
		Columns("incident_id", "component_id", "description", "updated_by").
		Record(incident).
		Exec()
	CheckErr(err)
	//============================================================================

	if isResolved == "true" {
		//update status of incident to resolved ...===================================
		_, err3 := sess.Update("incidents").
			Set("status", "resolved").
			Set("resolved_at", "NOW()").
			Where("id = ?", incident.IncidentId).
			Exec()
		CheckErr(err3)
		//============================================================================
	}
}

// IncidentInformation : incident info.
type IncidentInformation struct {
	Status      string
	Recorder    string
	Machine     string
	Component   string
	ComponentId int
	Description string

	IncidentUpdates []IncidentUpdate
}

// IncidentInformations : incident info.
func IncidentInformations(incident_id int) []byte {
	sess := SetupDB()
	m := IncidentInformation{}
	err2 := sess.Select(`incidents.status,
                       incidents.recorder,
                       incidents.component_id,
                       components.name as Component,
                       incidents.Description`).
		From("incidents").
		LeftJoin("components", "components.id = incidents.component_id").
		Where("incidents.id = ? ", incident_id).
		LoadStruct(&m)
	CheckErr(err2)

	p := []IncidentUpdate{}
	sess.Select("id, description, updated_by, created_at").
		From("incident_update").
		Where("incident_id = ? ", incident_id).
		LoadStruct(&p)

	for i := 0; i < len(p); i += 1 {
		t := p[i].Created_at
		p[i].Resolved_Date = t.Format("2006-01-02")
	}

	m.IncidentUpdates = p
	b, err := json.Marshal(m)
	CheckErr(err)
	return b
}

// IncidentComponentInfo : incident information releated to component.
type IncidentComponentInfo struct {
	ComponentId int
	InvoiceId   int

	IncidentId    int
	UpdatedBy     string
	CategoryId    int
	Name          string
	SerialNo      string
	Warranty_till *time.Time
	Description   string
	Active        bool
}

// IncidentAddComponent : add new incident on component.
func IncidentAddComponent(incident_id int, resolvedBy string, categoryId int, component string, serialNo string, description string) {
	sess := SetupDB()

	//select component id and invoice id on which incident happen ....============
	c := IncidentComponentInfo{}
	sess.Select("incidents.component_id, components.invoice_id, components.warranty_till").
		From("incidents").
		Join("components", "components.id = incidents.component_id").
		Where("incidents.id = ?", incident_id).
		LoadStruct(&c)
	//============================================================================

	//Add replaces components after resolved ....=================================
	c.IncidentId = incident_id
	c.UpdatedBy = resolvedBy
	c.CategoryId = categoryId
	c.Name = component
	c.SerialNo = serialNo
	c.Description = description
	c.Active = false

	_, err3 := sess.InsertInto("components").
		Columns("name", "invoice_id", "serial_no", "warranty_till", "description", "category_id", "active").
		Record(c).
		Exec()
	CheckErr(err3)
	//============================================================================

	//===Add resolved date in incidents table after resolved component ===========
	_, err5 := sess.Update("incidents").
		Set("resolved_at", "NOW()").
		Set("status", "Resolved").
		Where("id = ?", incident_id).
		Exec()
	CheckErr(err5)
	//============================================================================

	//=== deccommite component on which incident was happened ====================
	_, err6 := sess.Update("components").
		Set("deleted_at", "NOW()").
		Where("id = ?", c.ComponentId).
		Exec()
	CheckErr(err6)
	//============================================================================

	//=== add update in incident update ==========================================
	c.Description = "Component Replaced by New Component"
	_, err7 := sess.InsertInto("incident_update").
		Columns("incident_id", "component_id", "description", "updated_by").
		Record(c).
		Exec()
	CheckErr(err7)
	//============================================================================

}
