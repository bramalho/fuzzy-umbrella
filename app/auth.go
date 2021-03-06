package app

import (
	"context"
	"fuzzy-umbrella/models"
	u "fuzzy-umbrella/utils"
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

// JwtAuthentication handels JWT tokens
var JwtAuthentication = func(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		notAuth := []string{"/", "/register", "/login"}
		requestPath := r.URL.Path

		for _, value := range notAuth {
			if value == requestPath {
				next.ServeHTTP(w, r)
				return
			}
		}

		response := make(map[string]interface{})
		tokenHeader := r.Header.Get("Authorization")

		if tokenHeader == "" {
			response = u.Message(false, "Missing Authorization token")
			w.WriteHeader(http.StatusForbidden)
			u.Respond(w, response)
			return
		}

		splitted := strings.Split(tokenHeader, " ")
		if len(splitted) != 2 {
			response = u.Message(false, "Invalid/Malformed Authorization token")
			w.WriteHeader(http.StatusForbidden)
			u.Respond(w, response)
			return
		}

		tokenPart := splitted[1]
		tk := &models.Token{}

		token, err := jwt.ParseWithClaims(tokenPart, tk, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("token_password")), nil
		})

		if err != nil {
			response = u.Message(false, "Malformed Authorization token")
			w.WriteHeader(http.StatusForbidden)
			u.Respond(w, response)
			return
		}

		if !token.Valid {
			response = u.Message(false, "Token not valid.")
			w.WriteHeader(http.StatusForbidden)
			u.Respond(w, response)
			return
		}

		ctx := context.WithValue(r.Context(), "user", tk.UserID)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}
