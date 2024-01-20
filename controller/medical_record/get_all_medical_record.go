package medical_record

import (
	"net/http"

	"github.com/faizallmaullana/be_rsGundar/models"
	"github.com/gin-gonic/gin"
)

func GetAllMedicalRecord(c *gin.Context) {
	var medicalRecord []models.MedicalRecord
	if err := models.DB.Find(&medicalRecord); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database tidak bisa diakses"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"medical_record": medicalRecord})
}
