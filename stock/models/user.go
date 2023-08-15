package models

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)


type User struct {
	ID 				uuid.UUID 		`db:"id" json:"user_id" validate:"required,uuid"`
	FirstName		string			`db:"firstname" json:"firstname" validate:"required,lte=255"`
	LastName		string			`db:"lastname" json:"lastname" validate:"required,lte=255"`
	Email			string			`db:"email" json:"email" validate:"required,email"`
	Password		string			`db:"password" json:"password,omitempty" validate:"required,lte=255"`
	Role			Role				`db:"role" json:"role" validate:"required,lte=30"`
	CreatedAt		time.Time		`db:"created_at" json:"created_at"`
	UpdatedAt		time.Time		`db:"updated_at" json:"updated_at"`
	Portfolios		[]Portfolio		`db:"portfolios" json:"portfolios"`
}


type SignIn struct {
	Email		string		`json:"email" validate:"required,email,lte=255"`
	Password    string		`json:"password" validate:"required,lte=255"`
}


type SignUp struct {
	FirstName		string			`json:"firstname" validate:"required,lte=255"`
	LastName		string			`db:"lastname" json:"lastname" validate:"required,lte=255"`
	Email			string			`db:"email" json:"email" validate:"required,email"`
	Password		string			`db:"password" json:"password,omitempty" validate:"required,lte=255"`
	Role			Role				`db:"role" json:"role" validate:"required,lte=30"`
}


type UserQueries struct {
	*sqlx.DB
}

func (uq *UserQueries) CreateUser(ctx context.Context, user *User) error {
	query := `INSERT INTO users VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`

	_, err := uq.ExecContext(ctx, query, user.ID, user.FirstName, user.LastName, user.Email, user.Password, user.Role, user.CreatedAt)

	if err != nil {
		return err
	}

	return nil
}
func (uq *UserQueries) RetrieveUser(ctx context.Context, id uuid.UUID) (*User, error) {
	var user User
	query := `SELECT * FROM users WHERE id = $1`

	err := uq.GetContext(ctx, &user, query, id)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (uq *UserQueries) RetrieveUserByEmail(ctx context.Context, email string) (*User, error) {
	var user User
	query := `SELECT * FROM users WHERE email=$1`

	err := uq.GetContext(ctx, &user, query, email)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (uq *UserQueries) GetUsers(ctx context.Context) (*[]User, error) {
	var users []User
	query := `SELECT * FROM users`

	err := uq.SelectContext(ctx, &users, query)
	if err != nil {
		return nil, err
	}

	return &users, nil
}
func (uq *UserQueries) UpdateUser(ctx context.Context, id uuid.UUID) error {
	query := `UPDATE users SET firstname=$2 lastname=$3 WHERE id=$1`

	_, err := uq.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}
func (uq *UserQueries) DeleteUser(ctx context.Context, id uuid.UUID) error {
	query := `DELETE from users WHERE id=$1`

	_, err := uq.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}

func (uq *UserQueries) GetUserPortfolio(ctx context.Context, id uuid.UUID) (*[]Portfolio, error) {
	var portfolios []Portfolio
	query := `SELECT * FROM portfolios WHERE id=$1`

	err := uq.SelectContext(ctx, &portfolios, query, id)
	if err != nil {
		return nil, err
	}

	return &portfolios, nil
}


