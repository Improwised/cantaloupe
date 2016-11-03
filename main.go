package main

import (
		"fmt"
		"io/ioutil"
		"net/http"
		"strconv"
		_"github.com/lib/pq"
		"github.com/zenazn/goji"
		_"github.com/gocraft/dbr"
		"github.com/zenazn/goji/web"

		"./services"
 )

func usersHandler(c web.C, w http.ResponseWriter, r *http.Request) {
	all := r.URL.Query().Get("all")//either display all users or only active users ...
	if all != "true"{//true for display all users ...
		all = "false" // false for display only active users ...
	} // default it will display active users only ...
	b := services.DisplayUsers(all)
	w.Write([]byte(b))
}

func addUserHandler(c web.C, w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	company_email := r.FormValue("company_email")
	machine_id := r.FormValue("machine_id")
	fmt.Println(name, company_email, machine_id)
	services.AddNewUser(name, company_email, machine_id) //PATH : /services/users.go
}

func editUserInfoHandler(c web.C, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	user_id, err := strconv.Atoi(c.URLParams["user_id"]) // converting from string to int ...
	updated_data := r.FormValue("information")
	fmt.Println(user_id, "--->", updated_data)
	services.CheckErr(err)
	services.EditUserInfo(user_id, updated_data) //PATH : /services/users.go
}

func deleteUserHandler(c web.C, w http.ResponseWriter, r *http.Request) {
	user_id, err := strconv.Atoi(c.URLParams["user_id"]) // converting from string to int ...
	services.CheckErr(err)
	services.DeleteUser(user_id) //PATH : /services/users.go
}

func displayOneUserHandler(c web.C, w http.ResponseWriter, r *http.Request) {
	user_id, err := strconv.Atoi(c.URLParams["user_id"]) // converting from string to int ...
	services.CheckErr(err)
	b := services.DisplayUser(user_id) //PATH : /services/users.go
	w.Write([]byte(b))
}

func machinesHandler(c web.C, w http.ResponseWriter, r *http.Request) {
	all := r.URL.Query().Get("all") //either display all machines or only active machines ...
	if all != "true"{
		all = "false"
	}
	b := services.DisplayMachines(all) //PATH : /services/machines.go
	w.Write([]byte(b))
}

func addMachineHandler(c web.C, w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	services.AddNewMachine(name) //PATH : /services/machines.go
}

func editMachineHandler(c web.C, w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(c.URLParams["machine_id"]) // converting from string to int ...
	services.CheckErr(err)
	name := r.FormValue("name")
	services.EditMachineInfo(id, name) //PATH : /services/machines.go
}

func deleteMachineHandler(c web.C, w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(c.URLParams["machine_id"]) // converting from string to int ...
	services.CheckErr(err)
	services.DeleteMachine(id) //PATH : /services/machines.go
}

func displayMachineHandler(c web.C, w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(c.URLParams["machine_id"]) // converting from string to int ...
	services.CheckErr(err)
	services.DisplayMachine(id) //PATH : /services/machines.go
}

func machineComponentsHandler(c web.C, w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(c.URLParams["machine_id"]) // converting from string to int ...
	services.CheckErr(err)
	all := r.URL.Query().Get("all") //either display all components or only active components ...
	if all != "true"{
		all = "false"
	}
	response := services.DisplayMachineComponents(id, all) //PATH : /services/machines.go
	w.Write([]byte(response))
}

func addComponentsToMachine(c web.C, w http.ResponseWriter, r *http.Request) {
	machine_id, err := strconv.Atoi(c.URLParams["machine_id"]) // converting from string to int ...
	services.CheckErr(err)
	components_details := r.FormValue("information")
	services.AddComponentsToMachine(machine_id, components_details) //PATH : /services/machines.go
}

func removeComponentsFromMachine(c web.C, w http.ResponseWriter, r *http.Request) {
	machine_id, err := strconv.Atoi(c.URLParams["machine_id"]) // converting from string to int ...
	services.CheckErr(err)
	component_id, err2 := strconv.Atoi(c.URLParams["component_id"]) // converting from string to int ...
	services.CheckErr(err2)
	services.RemoveComponentsFromMachine(machine_id, component_id) //PATH : /services/machines.go
}

func changeUserFromMachine(c web.C, w http.ResponseWriter, r *http.Request) {
	machine_id, err := strconv.Atoi(c.URLParams["machine_id"]) // converting from string to int ...
	services.CheckErr(err)
	user_id, err2 := strconv.Atoi(c.URLParams["user_id"]) // converting from string to int ...
	services.CheckErr(err2)
	services.ChangeUserFromMachine(machine_id, user_id) //PATH : /services/machines.go
}

func InvoiceHandler(c web.C, w http.ResponseWriter, r *http.Request) {
	b := services.DisplayInvoices() //PATH : /services/invoices.go
	w.Write([]byte(b))
}

func addInvoiceHandler(c web.C, w http.ResponseWriter, r *http.Request) {
	body, err3 := ioutil.ReadAll(r.Body) // read all dara of form ...
	services.CheckErr(err3)
	services.AddInvoice(string(body)) //PATH : /services/invoices.go
}

func incidentsHandler(c web.C, w http.ResponseWriter, r *http.Request) {
	response := services.DisplayIncidents() //PATH : /services/incidents.go
	w.Write([]byte(response))
}

func addIncidentHandler(c web.C, w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body) // read all dara of form ...
	services.CheckErr(err)
	services.AddIncident(string(body)) //PATH : /services/incidents.go
}

func editIncidentHandler(c web.C, w http.ResponseWriter, r *http.Request) {
	incident_id, err := strconv.Atoi(c.URLParams["incidents_id"]) // converting from string to int ...
	services.CheckErr(err)
	componentId, err2 :=  strconv.Atoi(r.FormValue("component_id")) // converting from string to int ...
	services.CheckErr(err2)
	title := r.FormValue("title")
	description := r.FormValue("description")
	services.EditIncident(incident_id, componentId, title, description) //PATH : /services/incidents.go
}

func deleteIncidentHandler(c web.C, w http.ResponseWriter, r *http.Request) {
	incident_id, err := strconv.Atoi(c.URLParams["incidents_id"]) // converting from string to int ...
	services.CheckErr(err)
	services.DeleteIncident(incident_id) //PATH : /services/incidents.go
}

func displayIncidentHandler(c web.C, w http.ResponseWriter, r *http.Request) {
	incidents_id, err := strconv.Atoi(c.URLParams["incidents_id"]) // converting from string to int ...
	services.CheckErr(err)
	response := services.DisplayIncident(incidents_id) //PATH : /services/incidents.go
	w.Write([]byte(response))
}

func incidentsUpdateHandler(c web.C, w http.ResponseWriter, r *http.Request) {
	component_id, err := strconv.Atoi(r.FormValue("component_id")) // converting from string to int ...
	services.CheckErr(err)
	incident_id, err1 := strconv.Atoi(r.FormValue("incidents_id")) // converting from string to int ...
	services.CheckErr(err1)
	description := r.FormValue("description")
	services.IncidentUpdates(incident_id, component_id, description) //PATH : /services/incidents.go
}

func componentsHandler(c web.C, w http.ResponseWriter, r *http.Request) {
	all := r.URL.Query().Get("all") //either display all components or only active components ...
	if all != "true"{
		all = "false"
	}
	response := services.DisplayComponents(all) //PATH : /services/components.go
	w.Write([]byte(response))
}

func componentInfoHandler(c web.C, w http.ResponseWriter, r *http.Request) {
	component_id, err := strconv.Atoi(c.URLParams["component_id"]) // converting from string to int ...
	services.CheckErr(err)
	response := services.DisplayComponentInformation(component_id) //PATH : /services/components.go
	w.Write([]byte(response))
}

func saveHandler(w http.ResponseWriter, r *http.Request) {
    // allow cross domain AJAX requests
    w.Header().Set("Access-Control-Allow-Origin", "*")
}

func main() {
	//dealing with users ...
	goji.Get("/users", usersHandler)
	goji.Post("/users", addUserHandler)
	goji.Patch("/users/:user_id", editUserInfoHandler)
	goji.Delete("/users/:user_id", deleteUserHandler)
	goji.Get("/users/:user_id", displayOneUserHandler)

	//dealing with machines, it's user and its components ...
	goji.Get("/machines", machinesHandler)
	goji.Post("/machines", addMachineHandler)
	goji.Patch("/machines/:machine_id", editMachineHandler)
	goji.Delete("/machines/:machine_id", deleteMachineHandler)
	goji.Get("/machines/:machine_id", displayMachineHandler)
	goji.Get("/machines/:machine_id/components", machineComponentsHandler)
	goji.Post("/machines/:machine_id/components", addComponentsToMachine)
	goji.Delete("/machines/:machine_id/components/:component_id", removeComponentsFromMachine)
	goji.Delete("/machines/:machine_id/users/:user_id", changeUserFromMachine)

	//dealing with invoices ...
	goji.Post("/invoices", addInvoiceHandler)
	goji.Get("/invoices", InvoiceHandler)

	//dealing with incidents ...
	goji.Get("/incidents", incidentsHandler)
	goji.Post("/incidents", addIncidentHandler)
	goji.Patch("/incidents/:incidents_id", editIncidentHandler)
	goji.Delete("/incidents/:incidents_id", deleteIncidentHandler)
	goji.Get("/incidents/:incidents_id", displayIncidentHandler)
	goji.Post("/incidents/:incidents_id/update", incidentsUpdateHandler)

	//dealing with components ...
	goji.Get("/components", componentsHandler)
	goji.Get("/components/:component_id", componentInfoHandler)

	goji.Serve()
}
