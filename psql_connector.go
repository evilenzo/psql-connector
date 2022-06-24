package psql_connector

import (
	"database/sql"
	"fmt"

	"github.com/codingconcepts/env"
	log "github.com/sirupsen/logrus"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

type postgresConfig struct {
	User     string `env:"POSTGRES_USER" required:"true"`
	Password string `env:"POSTGRES_PASSWORD" required:"true"`
	Host     string `env:"POSTGRES_HOST"`
	Port     int16  `env:"POSTGRES_PORT"`
	DB       string `env:"POSTGRES_DB" required:"true"`
}

func createPostgresConfig() postgresConfig {
	pc := postgresConfig{Host: "postgres", Port: 5432}
	return pc
}

func getConfig() (postgresConfig, error) {
	config := createPostgresConfig()
	err := env.Set(&config)
	return config, err
}

func createConnection(config postgresConfig) *bun.DB {
	dsn := fmt.Sprintf("postgres://%[1]v:%[2]v@%[3]v:%[4]v/%[5]v?sslmode=disable",
		config.User,     // 1
		config.Password, // 2
		config.Host,     // 3
		config.Port,     // 4
		config.DB)       // 5

	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))
	db := bun.NewDB(sqldb, pgdialect.New())
	return db
}

func ConnectFromEnv() *bun.DB {
	config, err := getConfig()
	if err != nil {
		log.Error("Error during getting env config: ", err)
	}

	db := createConnection(config)
	return db
}
