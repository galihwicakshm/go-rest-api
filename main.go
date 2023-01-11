package main

import (
	"github.com/galihwicakshm/go-rest-api/initializers"
	"github.com/galihwicakshm/go-rest-api/routes"
)

func init() {
	initializers.LoadEnvVariables()

}
func main() {
	routes.SetupRoutes()
}
