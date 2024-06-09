package server

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/employee-store/employee"
)

// Server struct
type Server struct {
	Store *employee.EmployeeStore
}

// NewServer creates a new Server
func NewServer(store *employee.EmployeeStore) *Server {
	return &Server{Store: store}
}

// listEmployees handles the request to list employees with pagination
func (s *Server) listEmployees(w http.ResponseWriter, r *http.Request) {
	s.Store.mu.Lock()
	defer s.Store.mu.Unlock()

	pageStr := r.URL.Query().Get("page")
	if pageStr == "" {
		pageStr = "1"
	}
	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		http.Error(w, "Invalid page number", http.StatusBadRequest)
		return
	}

	perPageStr := r.URL.Query().Get("per_page")
	if perPageStr == "" {
		perPageStr = "10"
	}
	perPage, err := strconv.Atoi(perPageStr)
	if err != nil || perPage < 1 {
		http.Error(w, "Invalid per_page number", http.StatusBadRequest)
		return
	}

	employees := make([]employee.Employee, 0, len(s.Store.employees))
	for _, employee := range s.Store.employees {
		employees = append(employees, employee)
	}

	start := (page - 1) * perPage
	end := start + perPage

	if start > len(employees) {
		start = len(employees)
	}
	if end > len(employees) {
		end = len(employees)
	}

	response := employees[start:end]
	json.NewEncoder(w).Encode(response)
}
