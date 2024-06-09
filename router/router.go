package router

import (
	"example.com/greetings/server"

	"github.com/gorilla/mux"
)

// NewRouter creates a new router and sets up the routes
func NewRouter(srv *server.Server) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/employees", srv.ListEmployees).Methods("GET")
	r.HandleFunc("/addEmployee", srv.CreateEmployee).Methods("POST")
	r.HandleFunc("/getEmpDetails/{id:[0-9]+}", srv.GetEmployeeByID).Methods("GET")
	r.HandleFunc("/updateEmp/{id:[0-9]+}", srv.UpdateEmployee).Methods("PUT")
	r.HandleFunc("/removeEmp/{id:[0-9]+}", srv.DeleteEmployee).Methods("DELETE")
	return r
}
