package handlers

import (
	"encoding/json"
	"fuzzy-umbrella/models"
	u "fuzzy-umbrella/utils"
	"net/http"
)

// Register creates a new user
var Register = func(w http.ResponseWriter, r *http.Request) {
	user := &models.User{}
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}

	user, err = models.CreateUser(user)
	if err != nil {
		u.Respond(w, u.Message(false, err.Error()))
		return
	}

	resp := u.Message(true, "Registed")
	resp["user"] = user
	u.Respond(w, resp)
}

// Authenticate authenticates user
var Authenticate = func(w http.ResponseWriter, r *http.Request) {
	user := &models.User{}
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}

	user, err = models.Login(user.Email, user.Password)
	if err != nil {
		u.Respond(w, u.Message(false, err.Error()))
		return
	}

	resp := u.Message(true, "Logged In")
	resp["user"] = user
	u.Respond(w, resp)
}

// GetUser return user information
var GetUser = func(w http.ResponseWriter, r *http.Request) {
	resp := u.Message(true, "User Account")
	resp["user"], _ = models.GetUserByID(r.Context().Value("user").(string))

	u.Respond(w, resp)
}
