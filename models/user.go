package models

import (
	u "fuzzy-umbrella/utils"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

// Token struct for token
type Token struct {
	UserID uint
	jwt.StandardClaims
}

// User struct
type User struct {
	gorm.Model
	Email    string `json:"email"`
	Password string `json:"password"`
	Token    string `json:"token";sql:"-"`
}

// Validate user credentials
func (user *User) Validate() (map[string]interface{}, bool) {
	if !strings.Contains(user.Email, "@") {
		return u.Message(false, "Email address is required"), false
	}

	if len(user.Password) < 6 {
		return u.Message(false, "Password is required"), false
	}

	temp := &User{}

	err := GetDB().Table("users").Where("email = ?", user.Email).First(temp).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return u.Message(false, "Connection error. Please retry"), false
	}

	if temp.Email != "" {
		return u.Message(false, "Email address already userd"), false
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

	GetDB().Create(user)

	if user.ID <= 0 {
		return u.Message(false, "Connection error. Please retry")
	}

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
	err := GetDB().Table("users").Where("email = ?", email).First(user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return u.Message(false, "Invalid credentials")
		}
		return u.Message(false, "Connection error. Please retry")
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

// GetUser information
func GetUser(u uint) *User {
	user := &User{}
	GetDB().Table("users").Where("id = ?", u).First(user)
	if user.Email == "" {
		return nil
	}

	user.Password = ""
	return user
}
