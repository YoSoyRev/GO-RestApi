package routes

import (
	"encoding/json"
	"net/http"

	"github.com/YoSoyRev/go-resapi/db"
	"github.com/YoSoyRev/go-resapi/models"
	"github.com/gorilla/mux"
)

func GetTasksHandler(w http.ResponseWriter, r *http.Request) {
	var task []models.Task
	db.DB.Find(&task)
	json.NewEncoder(w).Encode(&task)
}

func CreateTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	json.NewDecoder(r.Body).Decode(&task)
	createTask := db.DB.Create(&task)
	err := createTask.Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest) // error 400
		w.Write([]byte(err.Error()))
		return
	}

	json.NewEncoder(w).Encode(&task)
}

func GetTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	paras := mux.Vars(r)

	db.DB.First(&task, paras["id"])

	if task.ID == 0 {
		w.WriteHeader(http.StatusNotFound) // error 404
		w.Write([]byte("Task not found"))
		return
	}
	json.NewEncoder(w).Encode(&task)
}

func DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	paras := mux.Vars(r)

	db.DB.First(&task, paras["id"])

	if task.ID == 0 {
		w.WriteHeader(http.StatusNotFound) // error 404
		w.Write([]byte("Task not found"))
		return
	}
	db.DB.Unscoped().Delete(&task)
	w.WriteHeader(http.StatusNoContent) //error 204
}
