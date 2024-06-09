package employee

import (
	"sync"
)

// EmployeeStore struct with thread safety
type EmployeeStore struct {
	mu        sync.Mutex
	employees map[int]Employee
	nextID    int
}

// NewEmployeeStore creates a new EmployeeStore
func NewEmployeeStore() *EmployeeStore {
	return &EmployeeStore{
		employees: make(map[int]Employee),
		nextID:    1,
	}
}

// CreateEmployee adds a new employee
func (store *EmployeeStore) CreateEmployee(name, position string, salary float64) Employee {
	store.mu.Lock()
	defer store.mu.Unlock()

	employee := Employee{
		ID:       store.nextID,
		Name:     name,
		Position: position,
		Salary:   salary,
	}
	store.employees[store.nextID] = employee
	store.nextID++
	return employee
}

// GetEmployeeByID retrieves an employee by ID
func (store *EmployeeStore) GetEmployeeByID(id int) (Employee, bool) {
	store.mu.Lock()
	defer store.mu.Unlock()

	employee, exists := store.employees[id]
	return employee, exists
}

// UpdateEmployee updates an existing employee
func (store *EmployeeStore) UpdateEmployee(id int, name, position string, salary float64) bool {
	store.mu.Lock()
	defer store.mu.Unlock()

	if _, exists := store.employees[id]; !exists {
		return false
	}

	store.employees[id] = Employee{
		ID:       id,
		Name:     name,
		Position: position,
		Salary:   salary,
	}
	return true
}

// DeleteEmployee deletes an employee by ID
func (store *EmployeeStore) DeleteEmployee(id int) bool {
	store.mu.Lock()
	defer store.mu.Unlock()

	if _, exists := store.employees[id]; !exists {
		return false
	}

	delete(store.employees, id)
	return true
}
