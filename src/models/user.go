package models

import (
	"auth/src/helpers"
	"time"
)

type User struct {
	ID        uint64    `json:"id"`
	Name      string    `json:"name"       gorm:"not null"               validate:"required"`
	Email     string    `json:"email"       gorm:"not null"               validate:"required"`
	Password  string    `json:"password"       gorm:"not null"               validate:"required"`
	CreatedAt time.Time `json:"created_at" gorm:"not null"`
	UpdatedAt time.Time `json:"updated_at" gorm:"not null"`
}

// Format formats User fields for better presentation
func (a *User) Format() *User {
	// a.Name = helpers.FormatString(a.Name)
	// a.Biography = helpers.FormatOptionalString(a.Biography)
	return a
}

// Valid returns custom validation error messages for User
func (a *User) Valid() helpers.ErrorMessages {
	customMessages := helpers.ErrorMessages{
		// "User.Name.required":      "Name is required",
		// "User.Age.required":       "Age is required",
		// "User.Biography.required": "Biography is required",
	}
	return helpers.ValidateStruct(a, customMessages)
}
