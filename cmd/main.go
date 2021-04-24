package main

import (
	"fmt"
	"net/http"

	transportHTTP "vess/go_rest_api/internal/transport/http"
)

// App contains pointers to database connections
type App struct{}

// Run starts the REST App server
func (app *App) Run() error {
	fmt.Println("Setting up App")
	handler := transportHTTP.NewHandler()
	handler.SetupRoutes()

	if err := http.ListenAndServe(":8080", handler.Router); err != nil {
		fmt.Println("Failed to setup server")
		return err
	}

	return nil
}

func main() {
	fmt.Println("Go REST API Course")
	app := App{}
	if err := app.Run(); err != nil {
		fmt.Println("Error starting the REST API.")
		fmt.Println(err)
	}
}
