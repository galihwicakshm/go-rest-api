package main

import (
	"github.com/galihwicakshm/go-rest-api/initializers"
	"github.com/galihwicakshm/go-rest-api/routes"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDatabase()

}
func main() {
	routes.SetupRoutes()
}
