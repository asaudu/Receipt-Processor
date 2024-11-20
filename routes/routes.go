package routes

import (
	"addyCodes.com/ReceiptProcessor/handlers"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	receiptRoutes := server.Group("/receipts")
	{
		receiptRoutes.POST("/process", handlers.CreateReceipt)

		receiptRoutes.GET("/:id/points", handlers.CalculateRewardPoints)

	}
}
