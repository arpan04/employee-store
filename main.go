package main

import (
	"net/http"

	"github.com/employee-store/router"
	"github.com/employee-store/server"

	"github.com/employee-store/employee"
)

func main() {
	store := employee.NewEmployeeStore()
	srv := server.NewServer(store)
	r := router.NewRouter(srv)

	http.ListenAndServe(":8080", r)
}
