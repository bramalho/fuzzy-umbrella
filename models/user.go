package models

import (
	"context"
	"errors"
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

// GetUserByID information
func GetUserByID(id string) (*User, error) {
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

func (user *User) validate() error {
	if !strings.Contains(user.Email, "@") {
		return errors.New("Email is required")
	}

	if len(user.Password) < 6 {
		return errors.New("Password is required")
	}

	user, err := GetByEmail(user.Email)
	if err != nil {
		return errors.New("Something went wrong")
	}

	if user.Email != "" {
		return errors.New("Invalid Email")
	}

	return nil
}

// CreateUser validates and create a new user
func CreateUser(user *User) (*User, error) {
	err := user.validate()
	if err != nil {
		return nil, err
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashedPassword)

	collection := GetDB().Collection("users")
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)

	result, err := collection.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}

	user.ID = result.InsertedID.(primitive.ObjectID)

	tk := &Token{UserID: user.ID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))
	user.Token = tokenString
	user.Password = ""

	return user, nil
}

// Login user
func Login(email, password string) (*User, error) {
	user := &User{}
	collection := client.Database("app_db").Collection("blogs")
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)

	err := collection.FindOne(ctx, User{Email: email, Password: password}).Decode(&user)
	if err != nil {
		return nil, errors.New("Something went wrong")
	}
	if user.Email != "" {
		return nil, errors.New("Invalid credentials")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return nil, errors.New("Invalid credentials")
	}

	tk := &Token{UserID: user.ID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))
	user.Token = tokenString
	user.Password = ""

	return user, nil
}
