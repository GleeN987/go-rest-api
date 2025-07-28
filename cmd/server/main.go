package main

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/GleeN987/go-rest-api/internal/comment"
	"github.com/GleeN987/go-rest-api/internal/db"
	transportHttp "github.com/GleeN987/go-rest-api/internal/transport/http"
	"github.com/golang-migrate/migrate/v4"
)

// Instantiating and starting application
func Run() error {
	fmt.Println("Starting application")
	time.Sleep(time.Second * 2)
	db, err := db.NewDatabase()
	if err != nil {
		fmt.Println("Failed to connect to db")
		return err
	}

	if err := db.MigrateDB(); err != nil {
		if !errors.Is(err, migrate.ErrNoChange) {
			fmt.Println("Failed to migrate database")
			return err
		}
	}
	commentService := comment.NewService(db)
	handler := transportHttp.NewHandler(commentService)
	if err := handler.Serve(); err != nil {
		log.Fatalf("fatal: %v", err)
	}

	return nil
}

func main() {
	fmt.Println("Go Rest API")
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
