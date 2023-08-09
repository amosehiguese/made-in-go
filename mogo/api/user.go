package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/amosehiguese/mogo/models"
	"github.com/go-chi/chi/v5"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Println(err)
		json.NewEncoder(w).Encode(NewError(false, http.StatusUnprocessableEntity, "Unprocessable" ))
		return
	}

	id, err := user.InsertOne()
	if err != nil {
		log.Println(err)
		json.NewEncoder(w).Encode(NewError(false, http.StatusInternalServerError, "Internal Server Error"))
		return
	}
	user.Id = id
	userResp := NewResponse(true, user)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(userResp)
}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := models.GetAllUsers()
	if err != nil {
		log.Println(err)
		json.NewEncoder(w).Encode(NewError(false, http.StatusInternalServerError, "Internal Server Error"))
		return
	}

	allUsersResp := NewResponse(true, users)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(allUsersResp)
}

func RetriveUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	user, err := models.RetrieveOne(id)
	if err != nil {
		log.Println(err)
		json.NewEncoder(w).Encode(NewError(false, http.StatusInternalServerError, "Internal Server Error"))
		return
	}

	userResp := NewResponse(true, user)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(userResp)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	var updateUser models.User
	err := json.NewDecoder(r.Body).Decode(&updateUser)
	if err != nil {
		log.Println(err)
		json.NewEncoder(w).Encode(NewError(false,http.StatusUnprocessableEntity, "Unprocessable" ))
		return
	}

	err = updateUser.UpdateOne(id)
	if err != nil {
		log.Println(err)
		json.NewEncoder(w).Encode(NewError(false, http.StatusInternalServerError, "Internal Server Error"))
		return
	}

	newUpdatedUser, err := models.RetrieveOne(id)
	if err != nil {
		log.Println(err)
		json.NewEncoder(w).Encode(NewError(false, http.StatusInternalServerError, "Internal Server Error"))
		return
	}

	UpdatedUserResp := NewResponse(true, newUpdatedUser)

	json.NewEncoder(w).Encode(UpdatedUserResp)
}


func DeleteUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	err := models.DeleteOne(id)
	if err != nil {
		log.Println(err)
		json.NewEncoder(w).Encode(NewError(false, http.StatusInternalServerError, "Internal Server Error"))
		return
	}
	dataResp := fmt.Sprintf("User %s is successfully deleted",id)
	resp := NewResponse(true, dataResp)
	json.NewEncoder(w).Encode(resp)
}
