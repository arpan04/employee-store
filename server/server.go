package server

import (
	"encoding/json"
	"net/http"
	"strconv"

	"example.com/greetings/employee"
	"github.com/gorilla/mux"
)

// Server struct
type Server struct {
	Store *employee.EmployeeStore
}

// NewServer creates a new Server
func NewServer(store *employee.EmployeeStore) *Server {
	return &Server{Store: store}
}

// listEmployees handles the request to list Employees with pagination
func (s *Server) ListEmployees(w http.ResponseWriter, r *http.Request) {
	s.Store.Mu.Lock()
	defer s.Store.Mu.Unlock()

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

	Employees := make([]employee.Employee, 0, len(s.Store.Employees))
	for _, employee := range s.Store.Employees {
		Employees = append(Employees, employee)
	}

	start := (page - 1) * perPage
	end := start + perPage

	if start > len(Employees) {
		start = len(Employees)
	}
	if end > len(Employees) {
		end = len(Employees)
	}

	response := Employees[start:end]
	json.NewEncoder(w).Encode(response)
}

// createEmployee handles the request to create a new employee
func (s *Server) CreateEmployee(w http.ResponseWriter, r *http.Request) {
	var newEmployee employee.Employee

	// payload validation
	if err := json.NewDecoder(r.Body).Decode(&newEmployee); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	createdEmployee := s.Store.CreateEmployee(newEmployee.Name, newEmployee.Position, newEmployee.Salary)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdEmployee)
}

// GetEmployeeByID handles the request to get an employee by ID
func (s *Server) GetEmployeeByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid employee ID", http.StatusBadRequest)
		return
	}

	employee, exists := s.Store.GetEmployeeByID(id)
	if !exists {
		http.Error(w, "Employee not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(employee)
}

// UpdateEmployee handles the request to update an employee by ID
func (s *Server) UpdateEmployee(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid employee ID", http.StatusBadRequest)
		return
	}

	var updatedEmployee employee.Employee
	if err := json.NewDecoder(r.Body).Decode(&updatedEmployee); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if !s.Store.UpdateEmployee(id, updatedEmployee.Name, updatedEmployee.Position, updatedEmployee.Salary) {
		http.Error(w, "Employee not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(updatedEmployee)
}

// DeleteEmployee handles the request to delete an employee by ID
func (s *Server) DeleteEmployee(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid employee ID", http.StatusBadRequest)
		return
	}

	if !s.Store.DeleteEmployee(id) {
		http.Error(w, "Employee not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
