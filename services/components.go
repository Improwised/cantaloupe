package services

import (
	"encoding/json"
	"time"
)

// DisplayAllComponents : all components.
type DisplayAllComponents struct {
	Id                 int
	Serial_no          string
	Name               string
	Machine            *string
	Invoice_id         int
	Category           *string
	Warranty_timestamp *time.Time
	Warranty_till      string
	Description        string
	Active             bool
	Deleted_at         *time.Time
}

func removeDuplicates(elements []DisplayAllComponents) []DisplayAllComponents {
	// Use map to record duplicates as we find them.
	encountered := map[int]bool{}
	result := []DisplayAllComponents{}

	for v := range elements {
		if encountered[elements[v].Id] == true {
			// Do not add duplicate.
		} else {
			// Record this element as an encountered element.
			encountered[elements[v].Id] = true
			// Append to result slice.
			result = append(result, elements[v])
		}
	}
	// Return the new slice.
	return result
}

// DisplayComponents : all components.
func DisplayComponents(all string) []byte {
	sess := SetupDB()
	components := []DisplayAllComponents{}

	query := sess.Select(`c.id,
                        c.invoice_id,
                        c.serial_no,
                        c.name,
                        c.active,
                        c.description,
                        c.warranty_till as Warranty_timestamp,
                        machines.name as Machine,
                        categories.category,
                        machine_components.component_id,
                        machine_components.created_at`).
		From("components c").
		LeftJoin("machine_components", "c.id = machine_components.component_id").
		LeftJoin("machines", "machines.id = machine_components.machine_id").
		LeftJoin("categories", "c.category_id = categories.id").
		OrderDir("machine_components.created_at", false)

	// display all components or active components only ...
	if all == "false" {
		query.Where("c.deleted_at IS Null").
			LoadStruct(&components)
	} else {
		query.LoadStruct(&components)
	}

	//extract only date from timestamp========
	for i := 0; i < len(components); i += 1 {
		t := components[i].Warranty_timestamp
		components[i].Warranty_till = t.Format("2006-01-02")
	}
	//================================
	result := removeDuplicates(components)

	b, err := json.Marshal(result)
	CheckErr(err)
	return b
}

// FilterComponents : filter component.
func FilterComponents(category_id int) []byte {
	sess := SetupDB()
	components := []DisplayAllComponents{}

	query := sess.Select(`c.id,
                        c.invoice_id,
                        c.serial_no,
                        c.name,
                        c.active,
                        c.description,
                        c.warranty_till as Warranty_timestamp,
                        machines.name as Machine,
                        categories.category`).
		From("components c").
		LeftJoin("machine_components", "c.id = machine_components.component_id").
		LeftJoin("machines", "machines.id = machine_components.machine_id").
		Join("categories", "c.category_id = categories.id")

	query.Where("c.deleted_at IS NULL AND c.category_id::text like '%%?%%'", category_id).
		LoadStruct(&components)

	//extract only date from timestamp========
	for i := 0; i < len(components); i += 1 {
		t := components[i].Warranty_timestamp
		components[i].Warranty_till = t.Format("2006-01-02")
	}
	//================================
	b, err := json.Marshal(components)
	CheckErr(err)
	return b
}

// ComponentLog : component history.
type ComponentLog struct {
	Id         int
	Machine    string
	Created_at NullTime
	Added_at   string
	Deleted_at NullTime
	Removed_at string
}

// DisplayComponentInfo : one component information.
type DisplayComponentInfo struct {
	Incidents    []Incidents
	ComponentLog []ComponentLog

	User                 User
	Machine              Machine
	Machine_component_id int
	Invoice_id           int
	Invoice_number       string
	Component            string
	Description          string
	Active               string
	Status               string
	Warranty_timestamp   time.Time
	Warranty_till        string
	Deleted_at           *time.Time
}

// Incidents : incident information.
type Incidents struct {
	Id          int
	Title       string
	Description string
	Recorder    string
	Status      string
}

// Machine : machine information.
type Machine struct {
	Id                   int
	Name                 string
	Machine_component_id int
}

// User : user's information.
type User struct {
	Id   int
	Name string
}

// DisplayComponentInformation : all information of component.
func DisplayComponentInformation(componentId int) []byte {
	sqlStmt := "invoices.id AS Invoice_id,"
	sqlStmt += "invoices.invoice_number,"
	sqlStmt += "components.name AS Component,"
	sqlStmt += "components.description,"
	sqlStmt += "components.active,"
	sqlStmt += "components.warranty_till As Warranty_timestamp,"
	sqlStmt += "components.deleted_at"

	sess := SetupDB()
	components := DisplayComponentInfo{}

	err := sess.Select(sqlStmt).
		From("invoices").
		RightJoin("components", "invoices.id = components.invoice_id").
		Where("components.id= ?", componentId).
		LoadStruct(&components)
	CheckErr(err)

	//extract only date from timestamp========
	t := components.Warranty_timestamp
	components.Warranty_till = t.Format("2006-01-02")
	//================================

	components.Status = "Active"
	if components.Active == "false" {
		components.Status = "Not Active"
	}

	//Get AllIncidents information happen with perticuler component ...
	incidents := []Incidents{}
	sess.Select("id, title, description, status, recorder").
		From("incidents").
		Where("Component_id = ?", componentId).
		LoadStruct(&incidents)

	//Get Machine Information to which component connected if connected ...
	machine := Machine{}
	sess.Select(`machines.id,
               machines.Name,
               machine_components.id AS Machine_component_id`).
		From("machines").
		Join("machine_components", "machine_components.machine_id = machines.id").
		Where("machine_components.Component_id = ? AND machine_components.deleted_at IS NULL", componentId).
		OrderDir("machine_components.id", false).
		Limit(1).
		LoadStruct(&machine)

	//Get User's Information owned machine to which component connected if owned...
	if machine.Name != "" {
		user := User{}
		sess.Select("users.id, users.name").
			From("users").
			LeftJoin("users_machine", "users.id = users_machine.user_id").
			Where("users_machine.machine_id = ?", machine.Id).
			LoadStruct(&user)
		components.User = user
	}

	//===== Will Give entire history of component ================================
	log := []ComponentLog{}
	sess.Select(`machine_components.id,
               machines.name AS Machine,
               machine_components.created_at,
               machine_components.deleted_at`).
		From("machine_components").
		LeftJoin("machines", "machine_components.machine_id = machines.id").
		Where("machine_components.component_id= ?", componentId).
		LoadStruct(&log)

	//extract only date from timestamp========
	for i := 0; i < len(log); i += 1 {
		t := log[i].Created_at.Time
		log[i].Added_at = t.Format("2006-01-02")

		t2 := log[i].Deleted_at.Time
		log[i].Removed_at = t2.Format("2006-01-02")
	}
	//================================
	//============================================================================

	//combine all information in one object and return it ...
	components.Incidents = incidents
	components.Machine = machine
	components.ComponentLog = log

	b, err5 := json.Marshal(components)
	CheckErr(err5)
	return b
}

// ActiveComponent : to active component.
func ActiveComponent(componentId int) {
	sess := SetupDB()
	_, err := sess.Update("components").
		Set("active", "true").
		Where("id = ?", componentId).
		Exec()
	CheckErr(err)
}

// DeactiveComponent : to deactive component.
func DeactiveComponent(componentId int) {
	sess := SetupDB()
	_, err := sess.Update("components").
		Set("active", "false").
		Where("id = ?", componentId).
		Exec()
	CheckErr(err)
}
