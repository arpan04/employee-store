package main

import (
	"net/http"

	"example.com/greetings/router"
	"example.com/greetings/server"

	"example.com/greetings/employee"
)

func main() {
	store := employee.NewEmployeeStore()
	srv := server.NewServer(store)
	r := router.NewRouter(srv)

	port := ":8080"
	print("Listening on port ", port)
	http.ListenAndServe(port, r)
}
