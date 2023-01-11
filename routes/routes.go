package routes

import (
	barangcontroller "github.com/galihwicakshm/go-rest-api/controllers/BarangController"
	"github.com/gin-gonic/gin"
)

func SetupRoutes() {
	r := gin.Default()
	r.GET("/api/products", barangcontroller.Index)
	r.Run()
}
