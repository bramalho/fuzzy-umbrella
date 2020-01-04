package handlers

import (
	"encoding/json"
	"fuzzy-umbrella/models"
	u "fuzzy-umbrella/utils"
	"net/http"
)

// CreateUser creates a new user
var CreateUser = func(w http.ResponseWriter, r *http.Request) {
	user := &models.User{}
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}

	resp := user.Create()
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

	resp := models.Login(user.Email, user.Password)
	u.Respond(w, resp)
}

// GetUser return user information
var GetUser = func(w http.ResponseWriter, r *http.Request) {
	resp := u.Message(true, "Logged In")
	resp["user"] = models.GetUser(r.Context().Value("user").(uint))

	u.Respond(w, resp)
}
