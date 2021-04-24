package main

import "fmt"

// App contains pointers to database connections
type App struct{}

// Run starts the REST App server
func (app *App) Run() error {
	fmt.Println("Setting up App")
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
