package main

import (
	"context"
	"fmt"
	"go-gomanager/db"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
    db *db.Postgres
}

func HttpServerConfig(db *db.Postgres) *Server {
    return &Server{
        db: db,
    }
}

var (
    // DATABASE_URL = "postgresql://postgres:password@172.25.208.1:5433/postgres"
    DATABASE_URL = os.Getenv("DATABASE_URL")
)

type SomeConfig struct {
    DATABASE_URL string
}

func LoadEnv() (*SomeConfig, error) {
    err := godotenv.Load()
    if err != nil {
        log.Printf("Error loading .env file: %v", err)
        return nil, err
    }

    dbUrl := os.Getenv("DATABASE_URL")
    return &SomeConfig{
        DATABASE_URL: dbUrl,
    }, nil
}

func main() {
    config, loadConfigErr := LoadEnv()
    if loadConfigErr != nil {
        log.Fatal(loadConfigErr)
    }

    app := echo.New()

    // * Init db connection
    fmt.Println(config.DATABASE_URL)
    pg, connErr := db.InitDBConnection(context.Background(), config.DATABASE_URL)
    if connErr != nil {
        log.Fatal("unable to connect to db")
    }

    pingErr := pg.Ping(context.Background())
    if pingErr != nil {
        log.Fatal("unable to connect to db")
    }
    log.Println("connected to db")

    // * Middleware
    app.Use(middleware.TimeoutWithConfig(middleware.TimeoutConfig{
        Skipper:      middleware.DefaultSkipper,
        ErrorMessage: "Timeout",
        Timeout:      30 * time.Second,
    }))

    app.GET("/", func(c echo.Context) error {
        return c.String(http.StatusOK, "sukses 200 ok")
    })

    app.Logger.Fatal(app.Start(":8000"))
}