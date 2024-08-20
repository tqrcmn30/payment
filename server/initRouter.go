package server

import (
	"payment/controller"

	"github.com/gin-gonic/gin"
)

func InitRouter(handler *controller.ControllerManager) *gin.Engine {
	router := gin.Default()
	api := router.Group("/api")

	api.GET("/home", func(ctx *gin.Context) {
		ctx.String(200, "Hello Payment")
	})

	EntityRoute := api.Group("/entity")
	{
		EntityRoute.GET("/", handler.GetListEntity)
		EntityRoute.GET("", handler.GetListEntity)
		EntityRoute.GET("/:id", handler.GetEntityTypeByID)
	}

	UserAccountRoute := api.Group("/Account")
	{
		UserAccountRoute.GET("", handler.GetListUserAccounts)
		UserAccountRoute.GET("/", handler.GetListUserAccounts)
		UserAccountRoute.GET("/:id", handler.GetUserAccountByID)
		UserAccountRoute.POST("/", handler.CreateUserAccount)
		UserAccountRoute.PUT("/:id", handler.UpdateUserAccount)
		UserAccountRoute.DELETE("/:id", handler.DeleteUserAccount)
	}

	PaymentTransactionRoute := api.Group("/Transaction")
	{
		PaymentTransactionRoute.GET("", handler.GetListPaymentTransaction)
		PaymentTransactionRoute.GET("/", handler.GetListPaymentTransaction)
		PaymentTransactionRoute.GET("/:No", handler.GetPaymentTransactionTypeByNo)
		PaymentTransactionRoute.POST("/", handler.CreatePaymentTransaction)
		PaymentTransactionRoute.DELETE("/:No", handler.DeletePaymentTransaction)
	}

	return router

}
