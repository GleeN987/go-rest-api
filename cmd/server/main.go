package main

import "fmt"

// Instantiating and starting application
func Run() error {
	fmt.Println("Starting application")
	return nil
}

func main() {
	fmt.Println("Go Rest API")
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
