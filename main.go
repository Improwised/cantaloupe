package main

import (
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	_ "github.com/gocraft/dbr"
	"github.com/improwised/cantaloupe/services"
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
	"github.com/rs/cors"
	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"
	"github.com/zenazn/goji/web/middleware"
)

func usersHandler(c web.C, w http.ResponseWriter, r *http.Request) {
	all := r.URL.Query().Get("all") //either display all users or only active users ...
	if all != "true" {              //true for display all users ...
		all = "false" // false for display only active users ...
	} // default it will display active users only ...
	b := services.DisplayUsers(all)
	w.Write([]byte(b))
}

func addUserHandler(c web.C, w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	companyEmail := r.FormValue("companyEmail")
	machineId := r.FormValue("machineId")
	response := services.AddNewUser(name, companyEmail, machineId) //PATH : /services/users.go
	w.Write([]byte(response))
}

func editUserInfoHandler(c web.C, w http.ResponseWriter, r *http.Request) {
	userId, err := strconv.Atoi(c.URLParams["userId"]) // converting from string to int ...
	name := r.FormValue("name")
	comapnyEmail := r.FormValue("companyEmail")
	machineId := r.FormValue("machineId")
	services.CheckErr(err)
	services.EditUserInfo(userId, name, comapnyEmail, machineId) //PATH : /services/users.go
}

func deleteUserHandler(c web.C, w http.ResponseWriter, r *http.Request) {
	ids := c.URLParams["userId"]
	services.DeleteUser(ids) //PATH : /services/users.go
}

func displayOneUserHandler(c web.C, w http.ResponseWriter, r *http.Request) {
	userId, err := strconv.Atoi(c.URLParams["userId"]) // converting from string to int ...
	services.CheckErr(err)
	b := services.DisplayUser(userId) //PATH : /services/users.go
	w.Write([]byte(b))
}

func machinesHandler(c web.C, w http.ResponseWriter, r *http.Request) {
	all := r.URL.Query().Get("all") //either display all machines or only active machines ...
	if all != "true" {
		all = "false"
	}
	b := services.DisplayMachines(all) //PATH : /services/machines.go
	w.Write([]byte(b))
}

func addMachineHandler(c web.C, w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	response := services.AddNewMachine(name) //PATH : /services/machines.go
	w.Write([]byte(response))
}

func editMachineHandler(c web.C, w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(c.URLParams["machineId"]) // converting from string to int ...
	services.CheckErr(err)
	name := r.FormValue("name")
	services.EditMachineInfo(id, name) //PATH : /services/machines.go
}

func deleteMachineHandler(c web.C, w http.ResponseWriter, r *http.Request) {
	ids := c.URLParams["machineId"]
	services.DeleteMachine(ids) //PATH : /services/machines.go
}

func displayMachineHandler(c web.C, w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(c.URLParams["machineId"]) // converting from string to int ...
	services.CheckErr(err)
	services.DisplayMachine(id) //PATH : /services/machines.go
}

func machineComponentsHandler(c web.C, w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(c.URLParams["machineId"]) // converting from string to int ...
	services.CheckErr(err)
	all := r.URL.Query().Get("all") //either display all components or only active components ...
	if all != "true" {
		all = "false"
	}
	response := services.DisplayMachineComponents(id, all) //PATH : /services/machines.go
	w.Write([]byte(response))
}

func addComponentsToMachine(c web.C, w http.ResponseWriter, r *http.Request) {
	machineId, err := strconv.Atoi(c.URLParams["machineId"]) // converting from string to int ...
	services.CheckErr(err)
	componentId, err := strconv.Atoi(r.FormValue("componentId")) // converting from string to int ...
	services.CheckErr(err)
	services.AddComponentsToMachine(machineId, componentId) //PATH : /services/machines.go
}

func removeComponentsFromMachine(c web.C, w http.ResponseWriter, r *http.Request) {
	machineId, err := strconv.Atoi(c.URLParams["machineId"]) // converting from string to int ...
	services.CheckErr(err)
	componentId, err2 := strconv.Atoi(c.URLParams["componentId"]) // converting from string to int ...
	services.CheckErr(err2)
	services.RemoveComponentsFromMachine(machineId, componentId) //PATH : /services/machines.go
}

func changeUserFromMachine(c web.C, w http.ResponseWriter, r *http.Request) {
	machineId, err := strconv.Atoi(c.URLParams["machineId"]) // converting from string to int ...
	services.CheckErr(err)
	userId, err2 := strconv.Atoi(c.URLParams["userId"]) // converting from string to int ...
	services.CheckErr(err2)
	services.ChangeUserFromMachine(machineId, userId) //PATH : /services/machines.go
}

func invoiceHandler(c web.C, w http.ResponseWriter, r *http.Request) {
	b := services.DisplayInvoices() //PATH : /services/invoices.go
	w.Write([]byte(b))
}

func addInvoiceHandler(c web.C, w http.ResponseWriter, r *http.Request) {
	body, err3 := ioutil.ReadAll(r.Body) // read all dara of form ...
	services.CheckErr(err3)
	services.AddInvoice(string(body)) //PATH : /services/invoices.go
}

func editInvoiceHandler(c web.C, w http.ResponseWriter, r *http.Request) {
	invoiceId, err := strconv.ParseInt(c.URLParams["invoiceId"], 10, 64) // converting from string to int64 ...
	services.CheckErr(err)
	body, err3 := ioutil.ReadAll(r.Body) // read all dara of form ...
	services.CheckErr(err3)
	services.EditInvoice(invoiceId, string(body)) //PATH : /services/invoices.go
}

func oneInvoiceDetailsHandler(c web.C, w http.ResponseWriter, r *http.Request) {
	invoiceId, err := strconv.Atoi(c.URLParams["invoiceId"]) // converting from string to int ...
	services.CheckErr(err)
	b := services.DisplayOneInvoice(invoiceId) //PATH : /services/invoices.go
	w.Write([]byte(b))
}

func incidentsHandler(c web.C, w http.ResponseWriter, r *http.Request) {
	all := r.URL.Query().Get("all") //either display all users or only active users ...
	if all != "true" {              //true for display all incidents ...
		all = "false" // false for display only active incidents ...
	} // default it will display active incidents only ...
	response := services.DisplayIncidents(all) //PATH : /services/incidents.go
	w.Write([]byte(response))
}

func addIncidentHandler(c web.C, w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body) // read all dara of form ...
	services.CheckErr(err)
	services.AddIncident(string(body)) //PATH : /services/incidents.go
}

func editIncidentHandler(c web.C, w http.ResponseWriter, r *http.Request) {
	incidentId, err := strconv.Atoi(c.URLParams["incidentId"]) // converting from string to int ...
	services.CheckErr(err)
	componentId := r.FormValue("componentId")
	title := r.FormValue("title")
	description := r.FormValue("description")
	recorder := r.FormValue("recorder")
	services.EditIncident(incidentId, componentId, recorder, title, description) //PATH : /services/incidents.go
}

func deleteIncidentHandler(c web.C, w http.ResponseWriter, r *http.Request) {
	incidentId, err := strconv.Atoi(c.URLParams["incidentId"]) // converting from string to int ...
	services.CheckErr(err)
	services.DeleteIncident(incidentId) //PATH : /services/incidents.go
}

func displayIncidentHandler(c web.C, w http.ResponseWriter, r *http.Request) {
	incidentId, err := strconv.Atoi(c.URLParams["incidentId"]) // converting from string to int ...
	services.CheckErr(err)
	response := services.DisplayIncident(incidentId) //PATH : /services/incidents.go
	w.Write([]byte(response))
}

func incidentsUpdateHandler(c web.C, w http.ResponseWriter, r *http.Request) {
	isResolved := r.URL.Query().Get("resolved")
	incidentId, err := strconv.Atoi(c.URLParams["incidentId"]) // converting from string to int ...
	services.CheckErr(err)
	description := r.FormValue("description")
	resolvedBy := r.FormValue("resolvedBy")
	services.IncidentUpdates(incidentId, resolvedBy, description, isResolved) //PATH : /services/incidents.go
}

func addComponentHandler(c web.C, w http.ResponseWriter, r *http.Request) {
	incidentId, err := strconv.Atoi(c.URLParams["incidentId"]) // converting from string to int ...
	services.CheckErr(err)
	description := r.FormValue("description")
	resolvedBy := r.FormValue("resolvedBy")
	serialNo := r.FormValue("serialNo")
	component := r.FormValue("component")
	categoryId, err2 := strconv.Atoi(r.FormValue("category"))
	services.CheckErr(err2)
	services.IncidentAddComponent(incidentId, resolvedBy, categoryId, component, serialNo, description) //PATH : /services/incidents.go
}

func incidentInfoHandler(c web.C, w http.ResponseWriter, r *http.Request) {
	incidentId, err := strconv.Atoi(c.URLParams["incidentId"]) // converting from string to int ...
	services.CheckErr(err)
	response := services.IncidentInformations(incidentId) //PATH : /services/incidents.go
	w.Write([]byte(response))
}

func componentsHandler(c web.C, w http.ResponseWriter, r *http.Request) {
	all := r.URL.Query().Get("all") //either display all components or only active components ...
	if all != "true" {
		all = "false"
	}
	response := services.DisplayComponents(all) //PATH : /services/components.go
	w.Write([]byte(response))
}

func setActiveComponentHandler(c web.C, w http.ResponseWriter, r *http.Request) {
	componentId, err := strconv.Atoi(c.URLParams["componentId"]) // converting from string to int ...
	services.CheckErr(err)
	services.ActiveComponent(componentId) //PATH : /services/components.go
}

func setDeactivecomponentHandler(c web.C, w http.ResponseWriter, r *http.Request) {
	componentId, err := strconv.Atoi(c.URLParams["componentId"]) // converting from string to int ...
	services.CheckErr(err)
	services.DeactiveComponent(componentId) //PATH : /services/components.go
}

func filterComponentsHandler(c web.C, w http.ResponseWriter, r *http.Request) {
	category, err := strconv.Atoi(r.URL.Query().Get("categoryId")) // converting from string to int ...
	services.CheckErr(err)
	response := services.FilterComponents(category) //PATH : /services/components.go
	w.Write([]byte(response))
}

func componentInfoHandler(c web.C, w http.ResponseWriter, r *http.Request) {
	componentId, err := strconv.Atoi(c.URLParams["componentId"]) // converting from string to int ...
	services.CheckErr(err)
	response := services.DisplayComponentInformation(componentId) //PATH : /services/components.go
	w.Write([]byte(response))
}

func categoriesHandler(c web.C, w http.ResponseWriter, r *http.Request) {
	all := r.URL.Query().Get("all") //either display all users or only active users ...
	if all != "true" {              //true for display all users ...
		all = "false" // false for display only active users ...
	} // default it will display active users only ...
	b := services.DisplayCategories(all)
	w.Write([]byte(b))
}

func addCategoryHandler(c web.C, w http.ResponseWriter, r *http.Request) {
	category := r.FormValue("category")
	description := r.FormValue("description")
	response := services.AddNewCategory(category, description) //PATH : /services/category.go
	w.Write([]byte(response))
}

func editCategoryInfoHandler(c web.C, w http.ResponseWriter, r *http.Request) {
	categoryId, err := strconv.Atoi(c.URLParams["categoryId"]) // converting from string to int ...
	category := r.FormValue("category")
	description := r.FormValue("description")
	services.CheckErr(err)
	services.EditCategoryInfo(categoryId, category, description) //PATH : /services/users.go
}

func deleteCategoryHandler(c web.C, w http.ResponseWriter, r *http.Request) {
	ids := c.URLParams["categoryId"]
	services.DeleteCategory(ids) //PATH : /services/users.go
}

//middleware for set Content-Type in applpication/json.
func setJSONResponse(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		h.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}

func main() {

	err := godotenv.Load("example_development.env")
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}

	//======= Middleware for AJAX Requsests ========================================
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "PATCH", "DELETE"},
		AllowedHeaders:   []string{"*"},
	})

	goji.Use(c.Handler)
	//==============================================================================

	APIs := web.New()
	goji.Handle("/*", APIs)
	APIs.Use(middleware.SubRouter)
	APIs.Use(setJSONResponse)

	//dealing with users ...
	APIs.Get("/users", usersHandler)
	APIs.Post("/users", addUserHandler)
	APIs.Patch("/users/:userId", editUserInfoHandler)
	APIs.Delete("/users/:userId", deleteUserHandler)
	APIs.Get("/users/:userId", displayOneUserHandler)

	//dealing with machines, it's user and its components ...
	APIs.Get("/machines", machinesHandler)
	APIs.Post("/machines", addMachineHandler)
	APIs.Patch("/machines/:machineId", editMachineHandler)
	APIs.Delete("/machines/:machineId", deleteMachineHandler)
	APIs.Get("/machines/:machineId", displayMachineHandler)
	APIs.Get("/machines/:machineId/components", machineComponentsHandler)
	APIs.Post("/machines/:machineId/components", addComponentsToMachine)
	APIs.Delete("/machines/:machineId/components/:componentId", removeComponentsFromMachine)
	APIs.Delete("/machines/:machineId/users/:userId", changeUserFromMachine)

	//dealing with invoices ...
	APIs.Post("/invoices", addInvoiceHandler)
	APIs.Get("/invoices", invoiceHandler)
	APIs.Get("/invoices/:invoiceId", oneInvoiceDetailsHandler)
	APIs.Patch("/invoices/:invoiceId", editInvoiceHandler)

	//dealing with incidents and its updates ...
	APIs.Get("/incidents", incidentsHandler)
	APIs.Post("/incidents", addIncidentHandler)
	APIs.Patch("/incidents/:incidentId", editIncidentHandler)
	APIs.Delete("/incidents/:incidentId", deleteIncidentHandler)
	APIs.Get("/incidents/:incidentId", displayIncidentHandler)
	APIs.Get("/incidents/:incidentId/incidentInfo", incidentInfoHandler)
	APIs.Post("/incidents/:incidentId/update", incidentsUpdateHandler)
	APIs.Post("/incidents/:incidentId/addComponent", addComponentHandler)

	//dealing with components ...
	APIs.Get("/components", componentsHandler)
	APIs.Get("/components/categories", categoriesHandler)
	APIs.Get("/components/filter", filterComponentsHandler)
	APIs.Get("/components/:componentId", componentInfoHandler)
	APIs.Get("/components/:componentId/activeIt", setActiveComponentHandler)
	APIs.Get("/components/:componentId/deactiveIt", setDeactivecomponentHandler)
	APIs.Post("/components/categories", addCategoryHandler)
	APIs.Patch("/components/categories/:categoryId", editCategoryInfoHandler)
	APIs.Delete("/components/categories/:categoryId", deleteCategoryHandler)

	port := os.Getenv("PORT") // if port set to different port.

	if port != "" {
		flag.Set("bind", ":"+port)
	}
	goji.Serve()
}
