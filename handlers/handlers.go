package handlers

import (
	"net/http"

	"addyCodes.com/ReceiptProcessor/models"
	"addyCodes.com/ReceiptProcessor/rewards"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CreateReceipt(c *gin.Context) {
	var receipt models.Receipt

	if err := c.ShouldBindJSON(&receipt); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Could not bind JSON payload to receipt object", "details": err.Error()})
		return
	}

	receipt.ID = uuid.New().String()
	receipt.CalculateTotal()

	models.SaveReceipt(receipt)

	c.JSON(http.StatusCreated, gin.H{
		"message": "Receipt created successfully!",
		"receipt": receipt,
	})
}

func GetReceipt(c *gin.Context) {
	id := c.Param("id")
	receipt, err := models.GetReceiptByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, receipt)
}

func CalculateRewardPoints(c *gin.Context) {
	id := c.Param("id")
	receipt, err := models.GetReceiptByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	points := rewards.RewardPoints(receipt)

	c.JSON(http.StatusOK, gin.H{"points": points})
}
