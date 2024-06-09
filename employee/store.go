package employee

import (
	"sync"
)

// EmployeeStore struct with thread safety
type EmployeeStore struct {
	Mu        sync.Mutex
	Employees map[int]Employee
	nextID    int
}

// NewEmployeeStore creates a new EmployeeStore
func NewEmployeeStore() *EmployeeStore {
	return &EmployeeStore{
		Employees: make(map[int]Employee),
		nextID:    1,
	}
}

// CreateEmployee adds a new employee
func (store *EmployeeStore) CreateEmployee(name, position string, salary float64) Employee {
	store.Mu.Lock()
	defer store.Mu.Unlock()

	employee := Employee{
		ID:       store.nextID,
		Name:     name,
		Position: position,
		Salary:   salary,
	}
	store.Employees[store.nextID] = employee
	store.nextID++
	return employee
}

// GetEmployeeByID retrieves an employee by ID
func (store *EmployeeStore) GetEmployeeByID(id int) (Employee, bool) {
	store.Mu.Lock()
	defer store.Mu.Unlock()

	employee, exists := store.Employees[id]
	return employee, exists
}

// UpdateEmployee updates an existing employee
func (store *EmployeeStore) UpdateEmployee(id int, name, position string, salary float64) bool {
	store.Mu.Lock()
	defer store.Mu.Unlock()

	if _, exists := store.Employees[id]; !exists {
		return false
	}

	store.Employees[id] = Employee{
		ID:       id,
		Name:     name,
		Position: position,
		Salary:   salary,
	}
	return true
}

// DeleteEmployee deletes an employee by ID
func (store *EmployeeStore) DeleteEmployee(id int) bool {
	store.Mu.Lock()
	defer store.Mu.Unlock()

	if _, exists := store.Employees[id]; !exists {
		return false
	}

	delete(store.Employees, id)
	return true
}
