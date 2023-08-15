package models

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type Stock struct {
	StockID 		uuid.UUID 		`db:"id" json:"id" validate:"required,uuid"`
	Name    		string 			`db:"name" json:"name" validate:"required,alpha,lte=255"`
	CurrentPrice    float64			`db:"current_price" json:"current_price" validate:"required,numeric"`
	Company 		string 			`db:"company" json:"company" validate:"required,alpha,lte=255"`
	MarketCap		float64			`db:"market_cap" json:"market_cap" validate:"required,numeric"`
	Sector		string	`db:"sector" json:"sector,omitempty" validate:"required,lte=255"`
	Holdings		[]Portfolio		`db:"holdings" json:"holdings,omitempty"`
	Prices			[]StockPrice	`db:"prices" json:"prices,omitempty"`
	Listings		[]Exchange		`db:"listings" json:"listings,omitempty"`
}

type Portfolio struct {
	ID					uuid.UUID 	`db:"id" json:"id" validate:"required,uuid"`
	UserID      		uuid.UUID	`db:"user_id" json:"user_id" validate:"required,uuid"`
	StockSymbol			string		`db:"stock_symbol" json:"stock_symbol" validate:"required"`
	Quantity			float64		`db:"quantity" json:"quantity" validate:"required,numeric"`
	PurchasePrice		float64		`db:"purchase_price" json:"purchase_price" validate:"required"`
	PurchaseDate		time.Time	`db:"purchase_date" json:"purchase_date" validate:"required"`

}

type StockPrice struct {
	ID				uuid.UUID 		`db:"id" json:"id" validate:"required,uuid"`
	StockSymbol		string			`db:"stock_symbol" json:"stock_symbol" validate:"required,lte=255"`
	Opening			float64			`db:"opening" json:"opening" validate:"required,numeric"`
	Closing			float64			`db:"closing" json:"closing" validate:"required,numeric"`
	High			float64			`db:"high" json:"high" validate:"required,numeric"`
	Low				float64			`db:"low" json:"low" validate:"required,numeric"`
	Volume			uint64			`db:"volume" json:"volume" validate:"required"`
}

type Exchange struct {
	Id			uuid.UUID
	Code		string
	Name		string
	Country     string
	City		string
	TradingHours	string
	NumberOfListings		string
	Website			string
}

type StockQueries struct {
	*sqlx.DB
}

func (s *StockQueries) GetStocks(ctx context.Context) (*[]Stock, error) {
	var stocks []Stock

	query := `SELECT * FROM stocks`

	err := s.SelectContext(ctx, &stocks, query)
	if err != nil {
		return nil, err
	}

	return &stocks, nil
}

func (s *StockQueries) RetrieveStock(ctx context.Context, id uuid.UUID) (*Stock, error) {
	var stock Stock

	query := `Select * FROM stocks WHERE id = $1`

	err := s.GetContext(ctx, &stock, query, id)
	if err != nil {
		return nil, err
	}

	return &stock, nil
}

func (s *StockQueries) CreateStock(ctx context.Context, stock *Stock) error {
	query := `INSERT INTO stocks VALUES ($1, $2, $3, $4, $5)`

	_, err := s.ExecContext(ctx, query, stock.StockID, stock.Name, stock.CurrentPrice, stock.Company, stock.MarketCap)
	if err != nil {
		return err
	}

	return nil
}

func (s *StockQueries) UpdateStock(ctx context.Context,id uuid.UUID, stock *Stock) error  {
	query := `UPDATE stocks SET name=$2, company=$3, price=$4 market_cap=$5 WHERE id=$1`

	_, err := s.ExecContext(ctx, query, id, stock.Name, stock.Company, stock.CurrentPrice, stock.MarketCap)
	if err != nil {
		return err
	}
	return nil
}

func (s *StockQueries) DeleteStock(ctx context.Context,id uuid.UUID) error {
	query := `DELETE FROM stocks WHERE id=$1 `

	_, err := s.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}


