package kategoricontroller

import (
	"net/http"

	"github.com/galihwicakshm/go-rest-api/initializers"
	"github.com/galihwicakshm/go-rest-api/models"
	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	var kategori []models.Kategori
	initializers.DB.Find(&kategori)
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "Kategori berhasil ditampilkan",
		"data":    kategori})
}

func Store(c *gin.Context) {

	var kategori models.Kategori

	if err := c.ShouldBindJSON(&kategori); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	initializers.DB.Create(&kategori)
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "Kategori berhasil ditambahkan",
		"data":    kategori})

}
