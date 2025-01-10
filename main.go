package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
	"github.com/naheedrayan/mongodb_golang_rest_api/usecase"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)


var mongoClient *mongo.Client

func init() {
	// Load .env environment variables
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	log.Println("Environment variables loaded")

	mongoClient , err = mongo.Connect(context.Background(), options.Client().ApplyURI(os.Getenv("MONGO_URI")))
	if err != nil {
		panic(err)
	}
	log.Println("Connected to MongoDB")

	err = mongoClient.Ping(context.Background(), nil)
	if err != nil {
		panic(err)
	}
	log.Println("Ping to MongoDB successful")


}

func main() {

	// Close the connection to MongoDB
	defer mongoClient.Disconnect(context.Background())

	coll := mongoClient.Database(os.Getenv("DB_NAME")).Collection(os.Getenv("COLLECTION_NAME"))

	empService := usecase.EmployeeService{MongoCollection: coll}



	// Start the chi server
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Post("/employees", empService.CreateEmployee)
	r.Get("/employees/{id}", empService.GetEmployeeByID)
	r.Get("/employees", empService.GetAllEmployee)
	r.Put("/employees/{id}", empService.UpdateEmployeeByID)
	r.Delete("/employees/{id}", empService.DeleteEmployeeByID)
	r.Delete("/employees", empService.DeleteAllEmployee)


	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})

	log.Println("Server is running on port 3000")
	http.ListenAndServe(":3000", r)
	
}