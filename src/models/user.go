package models

import (
	"auth/src/helpers"
	"time"
)

type User struct {
	ID        uint64    `json:"id"`
	Name      string    `json:"name"       gorm:"not null"  validate:"required"`
	Email     string    `json:"email"      gorm:"not null"  validate:"required"`
	Password  string    `json:"password"   gorm:"not null"  validate:"required"`
	CreatedAt time.Time `json:"created_at" gorm:"not null"`
	UpdatedAt time.Time `json:"updated_at" gorm:"not null"`
}

type UserResponse struct {
	ID        uint64    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Format formats User fields for better presentation
func (u *User) Format() *User {
	u.Name = helpers.FormatString(u.Name)
	u.Email = helpers.FormatString(u.Email)
	u.Password = helpers.FormatString(u.Password)
	return u
}

// Valid returns custom validation error messages for User
func (u *User) Valid() helpers.ErrorMessages {
	customMessages := helpers.ErrorMessages{
		"User.Name.required":     "Name is required",
		"User.Email.required":    "Email is required",
		"User.Password.required": "Password is required",
	}
	return helpers.ValidateStruct(u, customMessages)
}

// ResponseFormat returns a map representation of User without sensitive fields
func (u *User) ResponseFormat() UserResponse {
	return UserResponse{
		ID:    u.ID,
		Name:  u.Name,
		Email: u.Email,
	}
}

// UsersToResponseFormat converts a slice of User to a slice of UserResponse
func UsersToResponseFormat(users []User) []UserResponse {
	formatted := make([]UserResponse, len(users))
	for i, user := range users {
		formatted[i] = user.ResponseFormat()
	}

	return formatted
}
