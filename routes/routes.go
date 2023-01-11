package routes

import (
	barangcontroller "github.com/galihwicakshm/go-rest-api/controllers/BarangController"
	kategoricontroller "github.com/galihwicakshm/go-rest-api/controllers/KategoriController"
	"github.com/gin-gonic/gin"
)

func SetupRoutes() {
	r := gin.Default()
	r.GET("/api/products", barangcontroller.Index)
	r.GET("/api/categories", kategoricontroller.Index)
	r.POST("/api/products", barangcontroller.Store)
	r.POST("/api/categories", kategoricontroller.Store)
	r.Run()
}
