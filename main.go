package main

import (
	"net/http"

	"github.com/YoSoyRev/go-resapi/db"
	"github.com/YoSoyRev/go-resapi/models"
	"github.com/YoSoyRev/go-resapi/routes"
	"github.com/gorilla/mux"
)

func main() {

	db.DBConnection()

	db.DB.AutoMigrate(models.Task{})
	db.DB.AutoMigrate(models.User{})

	r := mux.NewRouter()

	r.HandleFunc("/", routes.HomeHandler)

	r.HandleFunc("/user", routes.GetUsersHandler).Methods("GET")
	r.HandleFunc("/user", routes.PostUserHandler).Methods("POST")
	r.HandleFunc("/user/{id}", routes.GetUserHandler).Methods("GET")
	r.HandleFunc("/user/{id}", routes.DeleteUserHandler).Methods("DELETE")

	// Task routes

	r.HandleFunc("/task", routes.GetTasksHandler).Methods("GET")
	r.HandleFunc("/task", routes.CreateTaskHandler).Methods("POST")
	r.HandleFunc("/task/{id}", routes.GetTaskHandler).Methods("GET")
	r.HandleFunc("/task", routes.DeleteTaskHandler).Methods("DELETE")


	http.ListenAndServe(":4000", r)
}
