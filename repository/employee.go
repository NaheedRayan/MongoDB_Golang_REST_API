package repository

import (
	"context"

	"github.com/naheedrayan/mongodb_golang_rest_api/model"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type EmployeeRepo struct{
	MongoCollection *mongo.Collection
}


func (e *EmployeeRepo) InsertEmployee(emp *model.Employee) (interface{}, error) {
	result , err := e.MongoCollection.InsertOne(context.Background(), emp)
	if err != nil {
		return nil, err
	}
	return result , nil
}

func (e *EmployeeRepo) FindEmployeeByID(id string) (*model.Employee, error) {
	var emp model.Employee
	err := e.MongoCollection.FindOne(context.Background(), bson.D{{"employee_id", id}}).Decode(&emp)
	if err != nil {
		return nil, err
	}
	return &emp, nil
}

func (e *EmployeeRepo) FindAllEmployee() ([]*model.Employee, error) {
	result, err := e.MongoCollection.Find(context.Background(), bson.D{})
	if err != nil {
		return nil, err
	}
	
	var employees []*model.Employee
	err = result.All(context.Background(), &employees)
	if err != nil {
		return nil, err
	}
	return employees, nil

}


func (e *EmployeeRepo) UpdateEmployeeByID(empID string, updatedEmp *model.Employee) (int64, error) {
	result, err := e.MongoCollection.UpdateOne(context.Background(), bson.D{{"employee_id", empID}}, bson.D{{"$set", updatedEmp}})
	if err != nil {
		return 0, err
	}
	return result.ModifiedCount, nil
}

func (e *EmployeeRepo) DeleteEmployeeByID(empID string) (int64, error) {
	result, err := e.MongoCollection.DeleteOne(context.Background(), bson.D{{"employee_id", empID}})
	if err != nil {
		return 0, err
	}
	return result.DeletedCount, nil
}

func (e *EmployeeRepo) DeleteAllEmployee() (int64, error) {
	result, err := e.MongoCollection.DeleteMany(context.Background(), bson.D{})
	if err != nil {
		return 0, err
	}
	return result.DeletedCount, nil
}
