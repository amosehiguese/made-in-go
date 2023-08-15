package models

import (
	"context"

	"github.com/google/uuid"
)

type StockLogic interface {
	GetStocks(ctx context.Context) ([]Stock, error)
	RetrieveStock(ctx context.Context, id uuid.UUID) (*Stock, error)
	CreateStock(ctx context.Context, stock *Stock) error
	UpdateStock(ctx context.Context, id uuid.UUID, stock *Stock) error
	DeleteStock(ctx context.Context, id uuid.UUID) error
}

type UserLogic interface {
	CreateUser(ctx context.Context, user *User) error
	RetrieveUser(ctx context.Context, id uuid.UUID) (*User, error)
	GetUsers(ctx context.Context) ([]User, error)
	UpdateUser(ctx context.Context, id uuid.UUID) error
	DeleteUser(ctx context.Context, id uuid.UUID) error
	GetUserPortfolio(ctx context.Context, id uuid.UUID) (*[]Portfolio, error)
	RetrieveUserByEmail(ctx context.Context, email string) (*User, error)
}
