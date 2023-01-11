package barangcontroller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "OK BARANG",
	})
}

func Store(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "OK BARANG",
	})
}
