package db

import (
	"fmt"

	"database/sql"

	"github.com/DionisiyGri/sample-go/internal/config"
	log "github.com/sirupsen/logrus"

	// Imported for migrations and database driver
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

// connectionTemplate postgres connection template
const connectionTemplate = "host=%s port=%d user=%s dbname=%s password=%s sslmode=disable"

// NewPostgresConnection creates new connection to postgres db with credentials provided in config
func NewPostgresConnection(conf *config.Config) (*sql.DB, error) {
	connectionStr := fmt.Sprintf(connectionTemplate,
		conf.DB.Host,
		conf.DB.Port,
		conf.DB.User,
		conf.DB.DBName,
		conf.DB.Password,
	)
	log.Debug("new postgres connection")

	return sql.Open("postgres", connectionStr)
}

//Ping is used to try to Ping database & return any error
func Ping(conf *config.Config) error {
	conn, err := NewPostgresConnection(conf)
	if err != nil {
		return err
	}

	defer conn.Close()

	return conn.Ping()
}
