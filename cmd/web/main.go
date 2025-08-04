package main

import (
	"fmt"
	"golang_tailwinds/internal/setup"
	"golang_tailwinds/routes"
	"net/http"
	"os"
)


func main() {
	// Define the port to listen on
	setup.SetUpIt()

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}

	address := fmt.Sprintf("0.0.0.0:%s", port)
	srv := &http.Server{
		Addr:    address,
		Handler: routes.NewRouter(),
	}
	fmt.Println(address)

	err := srv.ListenAndServe()
	if err != nil {
		fmt.Printf("Error ListenAndServe, %s", err)
	}
}
