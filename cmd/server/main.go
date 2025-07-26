package main

import (
	"context"
	"fmt"

	"github.com/GleeN987/go-rest-api/internal/db"
)

// Instantiating and starting application
func Run() error {
	fmt.Println("Starting application")

	db, err := db.NewDatabase()
	if err != nil {
		fmt.Println("Failed to connect to db")
		return err
	}

	if err := db.Ping(context.Background()); err != nil {
		return err
	}

	fmt.Println("pinged")
	return nil
}

func main() {
	fmt.Println("Go Rest API")
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
