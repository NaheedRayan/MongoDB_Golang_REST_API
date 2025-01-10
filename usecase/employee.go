package usecase

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/naheedrayan/mongodb_golang_rest_api/model"
	"github.com/naheedrayan/mongodb_golang_rest_api/repository"
	"go.mongodb.org/mongo-driver/mongo"
)

type EmployeeService struct {
	MongoCollection *mongo.Collection
}

type Response struct {
	Data interface{} `json:"data,omitempty"`
	Error string      `json:"error,omitempty"`
}


//------------------- Handler functions -------------------//
func (svc *EmployeeService) CreateEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	res := &Response{}
	defer json.NewEncoder(w).Encode(res)

	// Decode the request body
	var emp model.Employee
	err := json.NewDecoder(r.Body).Decode(&emp)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		res.Error = err.Error()
		return
	}

	// new imployee id
	emp.EmployeeID = uuid.New().String()

	repo := repository.EmployeeRepo{MongoCollection: svc.MongoCollection}

	// Insert the employee into the database
	insertID, err := repo.InsertEmployee(&emp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		res.Error = err.Error()
		return
	}

	res.Data = emp.EmployeeID
	w.WriteHeader(http.StatusCreated)

	log.Println("Employee inserted with ID ", insertID , emp)
}


func (svc *EmployeeService) GetEmployeeByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Initialize the response object
	res := &Response{}
	defer func() {
		if err := json.NewEncoder(w).Encode(res); err != nil {
			log.Println("Failed to encode response:", err)
		}
	}()

	// Extract employee ID from the URL using chi context
	empID := chi.URLParam(r, "id")
	if empID == "" {
		w.WriteHeader(http.StatusBadRequest)
		res.Error = "Employee ID is required."
		return
	}

	// Create a repository instance
	repo := repository.EmployeeRepo{MongoCollection: svc.MongoCollection}

	// Find the employee by ID
	emp, err := repo.FindEmployeeByID(empID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		res.Error = err.Error()
		return
	}

	// Set the response data
	res.Data = emp
	w.WriteHeader(http.StatusOK)

	log.Printf("Employee found with ID: %s, Data: %+v\n", empID, emp)
}


func (svc *EmployeeService) GetAllEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	res := &Response{}
	defer json.NewEncoder(w).Encode(res)

	repo := repository.EmployeeRepo{MongoCollection: svc.MongoCollection}

	// Find all the employees
	employees, err := repo.FindAllEmployee()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		res.Error = err.Error()
		return
	}

	res.Data = employees
	w.WriteHeader(http.StatusOK)

	log.Println("All employees found ", employees)

}

func (svc *EmployeeService) UpdateEmployeeByID(w http.ResponseWriter, r *http.Request) {
	
	w.Header().Add("Content-Type", "application/json")

	res := &Response{}
	defer json.NewEncoder(w).Encode(res)

	// Extract the employee ID from the URL
	empID := chi.URLParam(r, "id")
	if empID == "" {
		w.WriteHeader(http.StatusBadRequest)
		res.Error = "Employee ID is required"
		return
	}

	// Decode the request body
	var emp model.Employee
	err := json.NewDecoder(r.Body).Decode(&emp)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		res.Error = err.Error()
		return
	}

	repo := repository.EmployeeRepo{MongoCollection: svc.MongoCollection}

	// Update the employee by ID
	updatedCount, err := repo.UpdateEmployeeByID(empID, &emp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		res.Error = err.Error()
		return
	}

	res.Data = updatedCount
	w.WriteHeader(http.StatusOK)

	log.Println("Employee updated with ID ", empID, emp)
	
}

func (svc *EmployeeService) DeleteEmployeeByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	res := &Response{}
	defer json.NewEncoder(w).Encode(res)

	// Extract the employee ID from the URL
	empID := chi.URLParam(r, "id")
	if empID == "" {
		w.WriteHeader(http.StatusBadRequest)
		res.Error = "Employee ID is required"
		return
	}

	repo := repository.EmployeeRepo{MongoCollection: svc.MongoCollection}

	// Delete the employee by ID
	deletedCount, err := repo.DeleteEmployeeByID(empID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		res.Error = err.Error()
		return
	}

	res.Data = deletedCount
	w.WriteHeader(http.StatusOK)

	log.Println("Employee deleted with ID ", empID)
}

func (svc *EmployeeService) DeleteAllEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	res := &Response{}
	defer json.NewEncoder(w).Encode(res)

	repo := repository.EmployeeRepo{MongoCollection: svc.MongoCollection}

	// Delete all the employees
	deletedCount, err := repo.DeleteAllEmployee()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		res.Error = err.Error()
		return
	}

	res.Data = deletedCount
	w.WriteHeader(http.StatusOK)

	log.Println("All employees deleted")
}