package config

import (
	"fmt"

	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	"go.uber.org/fx"
)

type Database struct {
	Client *sqlx.DB
}

var client *sqlx.DB


func NewDatabase(config Config) (database Database, err error) {

	connInfo := fmt.Sprintf("host=%s:%s user=%s password=%s dbname=%s sslmode=disable",
		config.DbHost, config.DbPort, config.DbUsername, config.DbPassword, config.DbSchema)

	client, err = sqlx.Open("pgx", connInfo)
	if err != nil {
		fmt.Printf("Failed to open connection to database: %s", err.Error())
		return
	}
	if err = client.Ping(); err != nil {
		fmt.Printf("Failed to ping database: %s", err.Error())
		return
	}

	database.Client = client

	return
}

var DatabaseModule = fx.Provide(NewDatabase)
