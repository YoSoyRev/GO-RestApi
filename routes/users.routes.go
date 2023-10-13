package routes

import (
	"encoding/json"
	"net/http"

	"github.com/YoSoyRev/go-resapi/db"
	"github.com/YoSoyRev/go-resapi/models"
	"github.com/gorilla/mux"
)

func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	var user []models.User
	db.DB.Find(&user)
	json.NewEncoder(w).Encode(&user)
}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	paras := mux.Vars(r)
	db.DB.First(&user, paras["id"])

	if user.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("User not found"))
		return
	}

	db.DB.Model(&user).Association("Tasks").Find(&user.Tasks)

	json.NewEncoder(w).Encode(&user)
}

func PostUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)

	CreateUser := db.DB.Create(&user)
	err := CreateUser.Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}

	json.NewEncoder(w).Encode(&user)
}

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	paras := mux.Vars(r)
	db.DB.First(&user, paras["id"])

	if user.ID == 0 {
		w.WriteHeader(http.StatusNotFound) //error 404
		w.Write([]byte("User not found"))
		return
	}

	db.DB.Unscoped().Delete(&user)
	w.WriteHeader(http.StatusOK)
}
