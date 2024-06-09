package router

import (
	"github.com/employee-store/server"

	"github.com/gorilla/mux"
)

// NewRouter creates a new router and sets up the routes
func NewRouter(srv *server.Server) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/employees", srv.listEmployees).Methods("GET")
	return r
}
