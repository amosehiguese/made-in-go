package models

import (
	"time"

	"github.com/google/uuid"
)


type User struct {
	ID 				uuid.UUID 		`db:"id" json:"user_id" validate:"required,uuid"`
	FirstName		string			`db:"firstname" json:"firstname" validate:"required,lte=255"`
	LastName		string			`db:"lastname" json:"lastname" validate:"required,lte=255"`
	Email			string			`db:"email" json:"email" validate:"required,email"`
	Password		string			`db:"password" json:"password,omitempty" validate:"required,lte=255"`
	UserRole		string			`db:"user_role" json:"user_role" validate:"required,lte=30"`
	CreatedAt		time.Time		`db:"created_at" json:"created_at"`
	UpdatedAt		time.Time		`db:"updated_at" json:"updated_at"`
	Portfolios		[]Portfolio		`db:"portfolios" json:"portfolios"`
}


