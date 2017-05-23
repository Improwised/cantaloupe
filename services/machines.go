package services

import (
	"database/sql/driver"
	"encoding/json"
	"strings"
	"time"
)

// MachineInfo : machine information
type MachineInfo struct {
	Id         int64
	Name       string
	User       *string
	Deleted_at *time.Time
}

// AddNewMachine to add new machine's information.
func AddNewMachine(name string) []byte {
	sess := SetupDB()
	addMachine := MachineInfo{}
	addMachine.Name = name

	// Insert new machine ...
	_, err := sess.InsertInto("machines").
		Columns("name").
		Record(addMachine).
		Exec()
	CheckErr(err)

	//get id of inserted machine ...
	lastInsertedId, err := sess.Select("MAX(id)").
		From("machines").
		ReturnInt64()
	addMachine.Id = lastInsertedId

	b, err := json.Marshal(addMachine)
	CheckErr(err)
	return b
}

// EditMachineInfo to edit machine's information.
func EditMachineInfo(id int, name string) {
	sess := SetupDB()
	_, err := sess.Update("machines").
		Set("name", name).
		Set("modified_at", "NOW()").
		Where("id = ?", id).
		Exec()
	CheckErr(err)
}

// DeleteMachine to remove machine's info(soft delete).
func DeleteMachine(machineIds string) {
	ids := strings.Split(machineIds, ",")
	sess := SetupDB()

	//deleting multiple users ======
	for i := 0; i < len(ids); i += 1 {
		_, err := sess.Update("machines").
			Set("deleted_at", "NOW()").
			Where("id = ?", ids[i]).
			Exec()
		CheckErr(err)
	}
	//==============================
}

func removeDuplicatesUser(elements []MachineInfo) []MachineInfo {
	// Use map to record duplicates as we find them.
	encountered := map[int64]bool{}
	result := []MachineInfo{}

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

// DisplayMachines to list down all the machines.
func DisplayMachines(allMachiens string) []byte {
	sess := SetupDB()
	machinesInfo := []MachineInfo{}

	query := sess.Select("machines.id, users.name as User, machines.name").
		From("machines").
		LeftJoin("users_machine", "machines.id = users_machine.machine_id").
		LeftJoin("users", "users.id = users_machine.user_id").
		OrderDir("users_machine.created_at", false)

	//display all machines or active machines only ...
	if allMachiens == "false" {
		query.Where("machines.deleted_at IS NULL").
			LoadStruct(&machinesInfo)
	} else {
		query.LoadStruct(&machinesInfo)
	}
	result := removeDuplicatesUser(machinesInfo)

	b, err := json.Marshal(result)
	CheckErr(err)
	return b
}

// DisplayMachine to display information of specific machine.
func DisplayMachine(id int) []byte { // Display Single machine's information ...
	sess := SetupDB()
	machineInfo := MachineInfo{}

	sess.Select("u.id, u.name, machines.name").
		From("users u").
		LeftJoin("users_machine", "u.id = users_machine.user_id").
		LeftJoin("machines", "machines.id = users_machine.machine_id").
		Where("u.id = 1").
		LoadStruct(&machineInfo)

	b, err := json.Marshal(machineInfo)
	CheckErr(err)
	return b
}

// AllComponents : information of all components.
type AllComponents struct {
	Id            *int
	Name          *string
	SerialNo      *string
	Description   *string
	Warranty      *time.Time
	Warranty_till string
	Created_at    *time.Time
	AddOn         string
}

// AllIncidents : all incidents info.
type AllIncidents struct {
	Id          *int
	LastUpdate  *string
	Title       *string
	Description *string
	Status      *string
	Recorder    *string
}

// PastUses : history of machine.
type PastUses struct {
	Begin     NullTime
	BeginDate string
	End       NullTime
	EndDate   string
	User      *string
}

// AllInfoOfMachine : all machines with its component and users.
type AllInfoOfMachine struct {
	Id         int
	Name       string
	Machine    string
	Created_at NullTime
	UsingSince string

	Components []AllComponents
	Incidents  []AllIncidents
	PastUses   []PastUses
}

// ========To Handle Null Time stamp from database...===========================

// NullTime : to handle empty timestamp.
type NullTime struct {
	Time  time.Time
	Valid bool // Valid is true if Time is not NULL
}

// Scan implements the Scanner interface.
func (nt *NullTime) Scan(value interface{}) error {
	nt.Time, nt.Valid = value.(time.Time)
	return nil
}

// Value implements the driver Valuer interface.
func (nt NullTime) Value() (driver.Value, error) {
	if !nt.Valid {
		return nil, nil
	}
	return nt.Time, nil
}

//==============================================================================
// DisplayMachineComponents : components releated to machine.
func DisplayMachineComponents(machineId int, allComponents string) []byte {
	sess := SetupDB()
	//all machine's information ....========
	machineInfo := AllInfoOfMachine{}
	query2 := sess.Select(`machines.id,
												machines.name AS Machine,
												users_machine.created_at,
												users.name,
												users.id AS UserId`).
		From("machines").
		LeftJoin("users_machine", "users_machine.machine_id = machines.id").
		LeftJoin("users", "users_machine.user_id = users.id").
		Where("machines.id = ? AND users_machine.deleted_at IS NULL", machineId)
	query2.LoadStruct(&machineInfo)

	//==================================

	t := machineInfo.Created_at.Time
	machineInfo.UsingSince = t.Format("2006-01-02")

	//all information of machine's components ===========
	MachineComponents := []AllComponents{}
	query := sess.Select(`components.id,
												components.name,
												components.serial_no,
												components.description,
												components.warranty_till AS Warranty,
												machine_components.created_at`).
		From("components").
		Join("machine_components", "machine_components.component_id = components.id")

	query.Where("machine_components.machine_id = ? AND machine_components.deleted_at is NULL", machineId).
		LoadStruct(&MachineComponents)

	for i := 0; i < len(MachineComponents); i += 1 {
		time := MachineComponents[i].Created_at
		MachineComponents[i].AddOn = time.Format("2006-01-02")

		time2 := MachineComponents[i].Warranty
		MachineComponents[i].Warranty_till = time2.Format("2006-01-02")
	}

	incidents := []AllIncidents{}
	sess.Select("incidents.id, incidents.title, incidents.description, incidents.status, incidents.recorder").
		From("incidents").
		LeftJoin("components", "incidents.component_id=components.id").
		LeftJoin("machine_components", "machine_components.component_id=components.id").
		Where("machine_components.machine_id = ?", machineId).
		LoadStruct(&incidents)

	pastUses := []PastUses{}
	sess.Select("users_machine.created_at AS Begin, users_machine.deleted_at AS End, users.name AS User").
		From("users_machine").
		Join("users", "users_machine.user_id = users.id").
		Where("users_machine.machine_id = ?", machineId).
		LoadStruct(&pastUses)

	for i := 0; i < len(pastUses); i += 1 {
		t2 := pastUses[i].Begin.Time
		pastUses[i].BeginDate = t2.Format("2006-01-02")

		t3 := pastUses[i].End.Time
		pastUses[i].EndDate = t3.Format("2006-01-02")
	}

	//merge all info of machine and it's components in single object=======
	machineInfo.PastUses = pastUses
	machineInfo.Incidents = incidents
	machineInfo.Components = MachineComponents
	//============================================

	b, err := json.Marshal(machineInfo)
	CheckErr(err)
	return b
}

// MachineComponents : machine's component.
type MachineComponents struct {
	MachineId   int
	ComponentId int
}

// AddComponentsToMachine : add new component to machine.
func AddComponentsToMachine(machineId int, componentId int) {
	sess := SetupDB()
	m := MachineComponents{}
	m.MachineId = machineId
	m.ComponentId = componentId
	_, err := sess.Update("machine_components").
		Set("deleted_at", "NOW()").
		Where("component_id = ?", componentId).
		Exec()
	CheckErr(err)

	_, err2 := sess.InsertInto("machine_components").
		Columns("machine_id", "component_id").
		Record(m).
		Exec()
	CheckErr(err2)

	_, err3 := sess.Update("components").
		Set("active", "true").
		Where("id = ?", componentId).
		Exec()
	CheckErr(err3)
}

// RemoveComponentsFromMachine : remove component from machine.
func RemoveComponentsFromMachine(machineId int, componentId int) {
	sess := SetupDB()
	_, err := sess.Update("machine_components").
		Set("deleted_at", "NOW()").
		Where("component_id = ?", componentId).
		Exec()
	CheckErr(err)

	_, err2 := sess.Update("components").
		Set("active", false).
		Set("deleted_at", "NOW()").
		Where("id = ?", componentId).
		Exec()
	CheckErr(err2)
}

// ChangeUserFromMachine : change machine's user.
func ChangeUserFromMachine(machineId int, userId int) {
	sess := SetupDB()
	_, err := sess.Update("users_machine").
		Set("deleted_at", "NOW()").
		Where("machine_id = ? AND user_id = ?", machineId, userId).
		Exec()
	CheckErr(err)
}
