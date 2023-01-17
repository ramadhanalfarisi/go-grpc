package app

import (
	"database/sql"
	"fmt"
	"log"
	"net"
	"os"
	"time"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/ramadhanalfarisi/go-grpc/handlers"
	"github.com/ramadhanalfarisi/go-grpc/helpers"
	"github.com/ramadhanalfarisi/go-grpc/models"
	"google.golang.org/grpc"
)

var host, uname, password, port, dbname string

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		helpers.Error(err)
	}
	if env := os.Getenv("ENVIRONMMENT"); env == "production" {
		port = os.Getenv("DB_PORT")
		host = os.Getenv("DB_HOST")
		uname = os.Getenv("DB_USER")
		password = os.Getenv("DB_PASSWORD")
		dbname = os.Getenv("DB_NAME")
	} else if env == "development" {
		port = os.Getenv("DB_PORT_DEV")
		host = os.Getenv("DB_HOST_DEV")
		uname = os.Getenv("DB_USER_DEV")
		password = os.Getenv("DB_PASSWORD_DEV")
		dbname = os.Getenv("DB_NAME_DEV")
	} else {
		port = os.Getenv("DB_PORT_TEST")
		host = os.Getenv("DB_HOST_TEST")
		uname = os.Getenv("DB_USER_TEST")
		password = os.Getenv("DB_PASSWORD_TEST")
		dbname = os.Getenv("DB_NAME_TEST")
	}
}

type App struct {
	DB          *sql.DB
}

func (a *App) Run() {
	srv := grpc.NewServer()
	prodServer := handlers.RegisterProductHandler(a.DB)
	models.RegisterProductsServer(srv, prodServer)
	log.Println("gRPC run at port :8080")
	l, err:= net.Listen("tcp", ":8080")
	if err != nil {
		helpers.Error(err)
	}
	log.Fatal(srv.Serve(l))
}


func (a *App) Migrate() {
	driver, err := postgres.WithInstance(a.DB, &postgres.Config{})
	if err != nil {
		helpers.Error(err)
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file://./migrations",
		"postgres", driver)
	if err != nil {
		helpers.Error(err)
	}
	err2 := m.Up()
	if err2 != nil {
		helpers.Error(err2)
	}
}

func (a *App) ConnectDB() {
	strCon := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, uname, password, dbname)
	db, err := sql.Open("postgres", strCon)
	if err != nil {
		helpers.Error(err)
	}
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	a.DB = db
	a.Migrate()
}
