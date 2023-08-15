package store

import (
	"log"

	"github.com/amosehiguese/stock/models"
	"github.com/jmoiron/sqlx"
)

type queries struct {
	*models.StockQueries
	*models.UserQueries
}

func OpenDbConn() *queries {
	var (
		db *sqlx.DB
		err error
	)

	db, err = postgresSqlConn()
	if err != nil {
		log.Fatal("Unable to connect to db ->", err)
	}

	return &queries{StockQueries: &models.StockQueries{DB:db}, UserQueries: &models.UserQueries{DB:db}}

}
