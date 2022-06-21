package psql_connector

import (
	"fmt"
	"log"

	"github.com/codingconcepts/env"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func check(err error) {
	if err != nil {
		log.Fatalf("Fatal error during connecting to DB: %v", err)
	}
}

type postgresConfig struct {
	Host     string `env:"POSTGRES_HOST"`
	Port     int16  `env:"POSTGRES_PORT"`
	User     string `env:"POSTGRES_USER" required:"true"`
	Password string `env:"POSTGRES_PASSWORD" required:"true"`
	DBname   string `env:"POSTGRES_DB" required:"true"`
}

func createPostgresConfig() postgresConfig {
	pc := postgresConfig{Host: "localhost", Port: 5432}
	return pc
}

func getConfig() (postgresConfig, error) {
	config := createPostgresConfig()
	err := env.Set(&config)
	return config, err
}

func createConnection(config postgresConfig) (*gorm.DB, error) {
	dsn := "" +
		"host=%[1]v " +
		"port=%[2]v " +
		"user=%[3]v " +
		"password=%[4]v " +
		"dbname=%[5]v " +
		"sslmode=disable"

	dsn = fmt.Sprintf(dsn,
		config.Host,     // 0
		config.Port,     // 1
		config.User,     // 2
		config.Password, // 3
		config.DBname)   // 4

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	return db, err
}

func ConnectFromEnv() *gorm.DB {
	config, err := getConfig()
	check(err)

	db, err := createConnection(config)
	check(err)
	return db
}
