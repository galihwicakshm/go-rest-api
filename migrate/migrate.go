package main

import (
	"github.com/galihwicakshm/go-rest-api/initializers"
	"github.com/galihwicakshm/go-rest-api/models"
)

func init() {

	initializers.LoadEnvVariables()
	initializers.ConnectDatabase()

}

func main() {
	initializers.DB.AutoMigrate(&models.Product{})
}
