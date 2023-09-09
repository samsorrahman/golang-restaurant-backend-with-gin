package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/samsorrahman/golang-rest-backend/controllers"
)

func InvoiceRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/invoices", controller.GetInvoices)
	incomingRoutes.GET("/invoices/:id", controller.GetInvoice)
	incomingRoutes.POST("/invoices", controller.CreateInvoice)
	incomingRoutes.PATCH("/invoices/:id", controller.UpdateInvoice)
	// incomingRoutes.DELETE("/invoices/:id", controller.DeleteInvoice)
}
