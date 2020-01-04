package models

import (
	"context"
	u "fuzzy-umbrella/utils"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

// Token struct for token
type Token struct {
	UserID primitive.ObjectID
	jwt.StandardClaims
}

// User struct
type User struct {
	ID       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Email    string             `json:"email,omitempty" bson:"email,omitempty"`
	Password string             `json:"password,omitempty" bson:"password,omitempty"`
	Token    string             `json:"token,omitempty" bson:"token,omitempty"`
}

// GetByID information
func GetByID(id string) (*User, error) {
	var user User
	collection := GetDB().Collection("users")
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)

	oid, _ := primitive.ObjectIDFromHex(id)
	err := collection.FindOne(ctx, User{ID: oid}).Decode(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// GetByEmail find one user by email
func GetByEmail(email string) (*User, error) {
	var user User
	collection := GetDB().Collection("users")
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)

	err := collection.FindOne(ctx, User{Email: email}).Decode(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// Validate user credentials
func (user *User) Validate() (map[string]interface{}, bool) {
	if !strings.Contains(user.Email, "@") {
		return u.Message(false, "Email address is required"), false
	}

	if len(user.Password) < 6 {
		return u.Message(false, "Password is required"), false
	}

	user, err := GetByEmail(user.Email)
	if err != nil {
		return u.Message(false, "Somethign went wrong"), false
	}

	if user.Email != "" {
		return u.Message(false, "Invalid Email"), false
	}

	return u.Message(false, "success"), true
}

// Create a new user
func (user *User) Create() map[string]interface{} {
	if resp, ok := user.Validate(); !ok {
		return resp
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashedPassword)

	collection := GetDB().Collection("users")
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)

	result, err := collection.InsertOne(ctx, user)
	if err != nil {
		return u.Message(false, "Somethign went wrong")
	}

	user.ID = result.InsertedID.(primitive.ObjectID)

	tk := &Token{UserID: user.ID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))
	user.Token = tokenString
	user.Password = ""

	response := u.Message(true, "User successfuly created")
	response["user"] = user
	return response
}

// Login user
func Login(email, password string) map[string]interface{} {
	user := &User{}
	collection := client.Database("app_db").Collection("blogs")
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)

	err := collection.FindOne(ctx, User{Email: email, Password: password}).Decode(&user)
	if err != nil {
		return u.Message(false, "Connection error. Please retry")
	}
	if user.Email != "" {
		return u.Message(false, "Invalid credentials")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return u.Message(false, "Invalid credentials")
	}

	tk := &Token{UserID: user.ID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))
	user.Token = tokenString
	user.Password = ""

	resp := u.Message(true, "Logged In")
	resp["user"] = user
	return resp
}
