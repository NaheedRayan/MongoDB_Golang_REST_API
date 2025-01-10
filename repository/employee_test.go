package repository

import (
	"context"
	"log"
	"testing"

	"github.com/google/uuid"
	"github.com/naheedrayan/mongodb_golang_rest_api/model"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)



func NewMongoClient() *mongo.Client {
	
	mongoTestClient, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017"))

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to MongoDB!")


	err = mongoTestClient.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Ping to MongoDB successful!")
	return mongoTestClient
}


func TestMongoOperations(t *testing.T){
	mongoClient := NewMongoClient()
	defer mongoClient.Disconnect(context.Background())

	// dummy data
	emp1 := uuid.New().String()
	emp2 := uuid.New().String()

	coll := mongoClient.Database("companydb").Collection("employee_test")

	empRepo := EmployeeRepo{MongoCollection: coll}

	// InsertEmployee1
	t.Run("InsertEmployee", func(t *testing.T) {
		emp := &model.Employee{
			EmployeeID: emp1,
			Name:       "Tony Stark",
			Department: "Engineering",
		}
		result , err := empRepo.InsertEmployee(emp)
		if err != nil {
			t.Fatal("Error inserting employee1")
		}
		t.Log("Inserted Employee1: ", result)
	})

	// InsertEmployee2
	t.Run("InsertEmployee", func(t *testing.T) {
		emp := &model.Employee{
			EmployeeID: emp2,
			Name:       "Steve Rogers",
			Department: "Engineering",
		}
		result , err := empRepo.InsertEmployee(emp)
		if err != nil {
			t.Fatal("Error inserting employee2")
		}
		t.Log("Inserted Employee2: ", result)
	})


	// FindEmployeeByID
	t.Run("FindEmployeeByID", func(t *testing.T) {
		result, err := empRepo.FindEmployeeByID(emp1)
		if err != nil {
			t.Fatal("Error finding employee by id")
		}
		t.Log("Found Employee1: ", result)
	})

	// FindAllEmployee
	t.Run("FindAllEmployee", func(t *testing.T) {
		result, err := empRepo.FindAllEmployee()
		if err != nil {
			t.Fatal("Error finding all employees")
		}
		t.Log("Found all employees: ", result)
	})

	// UpdateEmployeeByID
	t.Run("UpdateEmployeeByID", func(t *testing.T) {
		updatedEmp := &model.Employee{
			EmployeeID: emp1,
			Name:       "Tony Pink",
			Department: "Engineering",
		}
		result, err := empRepo.UpdateEmployeeByID(emp1, updatedEmp)
		if err != nil {
			t.Fatal("Error updating employee by id")
		}
		t.Log("Updated Employee1: ", result)
	})

	// DeleteEmployeeByID
	t.Run("DeleteEmployeeByID", func(t *testing.T) {
		result, err := empRepo.DeleteEmployeeByID(emp1)
		if err != nil {
			t.Fatal("Error deleting employee by id")
		}
		t.Log("Deleted Employee1: ", result)
	})

	// DeleteAllEmployee
	t.Run("DeleteAllEmployee", func(t *testing.T) {
		result, err := empRepo.DeleteAllEmployee()
		if err != nil {
			t.Fatal("Error deleting all employees")
		}
		t.Log("Deleted all employees: ", result)
	})

}