package services

import (
	"encoding/json"
	"strings"
	"time"
)

// UserInformation : user's information.
type UserInformation struct {
	Id            int
	Name          string
	Machine_id    *int
	Machine_name  *string
	Company_email string
	Deleted_at    time.Time
}

// UsersMachine : user's machine.
type UsersMachine struct {
	User_id    int64
	Machine_id string
}

// userInfo : user's information.
type userInfo struct {
	Name          string
	Company_email string
	Machine_id    string
	Machine_name  string
	Id            int64
}

// AddNewUser : add new user.
func AddNewUser(name string, email string, machineId string) []byte {
	sess := SetupDB()
	var m userInfo
	m.Name = name
	m.Company_email = email
	m.Machine_id = machineId

	_, err := sess.InsertInto("users").
		Columns("name", "company_email").
		Record(m).
		Exec()
	CheckErr(err)

	//get id of inserted user ...
	lastInsertedId, err := sess.Select("MAX(id)").
		From("users").
		ReturnInt64()

	//get name of machine assigned to user ...
	machine, err2 := sess.Select("name").
		From("machines").
		Where("id= ? ", machineId).
		ReturnString()
	CheckErr(err2)
	m.Machine_name = machine

	addUserMachine := UsersMachine{}
	addUserMachine.User_id = lastInsertedId
	addUserMachine.Machine_id = m.Machine_id
	// Insert user's machine's information ...
	_, err1 := sess.InsertInto("users_machine").
		Columns("user_id", "machine_id").
		Record(addUserMachine).
		Exec()
	CheckErr(err1)

	m.Id = lastInsertedId
	b, err := json.Marshal(m)
	CheckErr(err)
	return b
}

// EditUserInfo : edit user's information.
func EditUserInfo(userId int, name string, company_email string, machine_id string) {
	sess := SetupDB()

	_, err := sess.Update("users").
		Set("name", name).
		Set("company_email", company_email).
		Set("modified_at", "NOW()").
		Where("id = ?", userId).
		Exec()
	CheckErr(err)

	_, err2 := sess.Update("users_machine").
		Set("machine_id", machine_id).
		Where("user_id = ?", userId).
		Exec()
	CheckErr(err2)
}

// DeleteUser : remove user (doft delete).
func DeleteUser(userIds string) {
	ids := strings.Split(userIds, ",")
	sess := SetupDB()

	//deleting multiple users ======
	for i := 0; i < len(ids); i += 1 {
		_, err := sess.Update("users").
			Set("deleted_at", "NOW()").
			Where("id = ?", ids[i]).
			Exec()
		CheckErr(err)
	}
	//==============================
}

// DisplayUser : display one user's information.
func DisplayUser(userId int) []byte {
	sess := SetupDB()
	userInfo := UserInformation{}
	err := sess.Select(`users.id,
											users.name,
											users.company_email,
											users_machine.machine_id,
											machines.name as Machine_name`).
		From("users").
		LeftJoin("users_machine", "users.id = users_machine.user_id").
		LeftJoin("machines", "users_machine.machine_id = machines.id").
		Where("users.id = ?", userId).LoadStruct(&userInfo)
	CheckErr(err)

	b, err := json.Marshal(userInfo)
	CheckErr(err)
	return b
}

// DisplayUsers : display all user's information.
func DisplayUsers(allUsers string) []byte { // Display one User's Information ..
	sess := SetupDB()
	usersInfo := []UserInformation{}
	query := sess.Select(`users.id,
												users.name,
												users.company_email,
												users_machine.machine_id,
												machines.name as Machine_name`).
		From("users").
		LeftJoin("users_machine", "users.id = users_machine.user_id").
		LeftJoin("machines", "users_machine.machine_id = machines.id")

	//display all users or active users only ...
	if allUsers == "false" {
		query.Where("users.deleted_at IS NULL").
			LoadStruct(&usersInfo)
	} else {
		query.LoadStruct(&usersInfo)
	}

	b, err := json.Marshal(usersInfo)
	CheckErr(err)

	return b
}
