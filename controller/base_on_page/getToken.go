package base_on_page

import (
	"net/http"

	"github.com/faizallmaullana/be_rsGundar/models"
	"github.com/gin-gonic/gin"
)

func GetToken(c *gin.Context) {
	var token models.DBToken
	if err := models.DB.First(&token); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "token tidak ditemukan",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token.Token,
	})
}
