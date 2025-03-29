package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"sync"
)

var (
	once     sync.Once
	database *Database
)

type Database struct {
	db *sql.DB
}

type Config struct {
	User            string
	Host            string
	Password        string
	DBName          string
	Port            string
	ApplicationName string
	SSLMode         string
}

func ConnectDatabase(config Config) (*Database, error) {
	var err error
	once.Do(func() {
		database, err = initDatabase(config)
	})

	return database, err
}

func initDatabase(config Config) (*Database, error) {
	ConnString := fmt.Sprintf("user=%s host=%s password=%s dbname=%s port=%s application_name=%s sslmode=%s",
		config.User,
		config.Host,
		config.Password,
		config.DBName,
		config.Port,
		config.ApplicationName,
		config.SSLMode,
	)

	db, err := sql.Open("postgres", ConnString)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	log.Print("Bank connection successful")

	return &Database{db: db}, nil
}

func (d *Database) DB() *sql.DB {
	return d.db
}

func (d *Database) Close() error {
	return d.db.Close()
}
