package main

import (
	"os"

	"github.com/DionisiyGri/sample-go/internal/config"
	"github.com/DionisiyGri/sample-go/internal/db"
	"github.com/DionisiyGri/sample-go/internal/db/migrations"
	"github.com/DionisiyGri/sample-go/internal/db/postgres"
	"github.com/DionisiyGri/sample-go/internal/db/repos"
	"github.com/DionisiyGri/sample-go/internal/sample"
	"github.com/DionisiyGri/sample-go/internal/server"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("No .env file found")
	}
}

const (
	defaultConfigFileName = "config.json"
	defaultPort           = "8000"
	pathToMigrations      = "internal/db/migrations/psql-migrations"
)

func main() {
	confFile := os.Getenv("CONFIG")
	if confFile == "" {
		confFile = defaultConfigFileName
	}

	conf, err := config.Load(confFile)
	if err != nil {
		log.Fatal(err)
	}

	loglvl, err := log.ParseLevel(conf.Log.Level)
	if err != nil {
		log.Fatalf("failed to parse provided log lvl: `%s`", err.Error())
	}
	log.SetLevel(loglvl)

	dbConnect, err := db.NewPostgresConnection(conf)
	if err != nil {
		log.Fatalf("db.NewPostgresConnection err:%s", err)
	}
	defer dbConnect.Close()

	// Run migrations
	migrator, err := migrations.RunMigrationPostgres(dbConnect, pathToMigrations)
	if err != nil {
		log.Fatal(err)
	}

	_, err = migrator.Up()
	if err != nil {
		log.Fatal(err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	someServicePsql := postgres.NewSomeService(dbConnect)

	someRepo := repos.NewSomeRepository(someServicePsql)

	app := sample.NewSampleApp(someRepo)

	err = server.Init(
		port,
		app,
		*conf,
	)
	if err != nil {
		log.Fatalf("error starts new server: %s", err)
	}
}
