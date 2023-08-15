package api

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/amosehiguese/stock/models"
	"github.com/google/uuid"
)

func SignUp(w http.ResponseWriter, r *http.Request) {
	signUp := models.SignUp{}
	signUp.Role = models.UserRole
	if err := json.NewDecoder(r.Body).Decode(&signUp); err != nil {
		log.Println("Unable to process the data ->",err)
		http.Error(w, "Unprocessable entity" ,http.StatusUnprocessableEntity)
		return
	}

	validate := newValidator()
	if err := validate.Struct(signUp); err != nil{
		log.Println("Data provided didn't pass validation ->",err)
		json.NewEncoder(w).Encode(NewError(false, http.StatusBadRequest, validatorErrors(err)))
		return
	}

	user := models.User{
		ID: uuid.New(),
		FirstName: signUp.FirstName,
		LastName: signUp.LastName,
		Email: signUp.Email,
		Password: generatePasswordHash(signUp.Password),
		Role: signUp.Role,
		CreatedAt: time.Now(),
	}

	if err := validate.Struct(user); err != nil {
		log.Println("User struct didn't pass validation ->",err)
		json.NewEncoder(w).Encode(NewError(false, http.StatusBadRequest, validatorErrors(err)))
		return
	}

	var uq models.UserQueries
	ctx, cancel := context.WithTimeout(context.Background(), 15 * time.Second)
	defer cancel()
	err := uq.CreateUser(ctx, &user)
	if err != nil {
		log.Println("Unable to create user in db ->", err)
		json.NewEncoder(w).Encode(NewError(false, http.StatusInternalServerError, "We're sorry, but there was an issue processing your request. Our team has been notified. " ))
		return
	}

	user.Password = ""

	json.NewEncoder(w).Encode(NewResponse(true, user))

}

func SignIn(w http.ResponseWriter, r *http.Request) {
	var signIn models.SignIn

	if err := json.NewDecoder(r.Body).Decode(&signIn); err != nil {
		log.Println("Unable to process data ->",err)
		http.Error(w, "Unprocessable entity", http.StatusUnprocessableEntity)
		return
	}

	var uq models.UserQueries
	ctx, cancel := context.WithTimeout(context.Background(), 15 * time.Second)
	defer cancel()
	user, err := uq.RetrieveUserByEmail(ctx, signIn.Email)
	if err != nil {
		log.Println("Unable to retrieve user for db ->",err)
		json.NewEncoder(w).Encode(NewError(false, http.StatusInternalServerError, "An unexpected error occurred. Please try again" ))
		return
	}

	cp := comparePasswords(user.Password, signIn.Password)
	if !cp && (signIn.Email != user.Email) {
		log.Println("Passwords not the same ->", err)
		json.NewEncoder(w).Encode(NewError(false, http.StatusUnauthorized, "Authentication failed: Invalid email and password"))
	}

}

func SignOut(w http.ResponseWriter, r *http.Request) {}
