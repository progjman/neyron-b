package routes

import (
	"neyron-b/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine { // Настройка маршрутов api
	r := gin.Default()

	r.SetTrustedProxies(nil) //Отключаем доверее к прокси (безопаснее)

	r.POST("/accounts", handlers.CreateAccount)
	r.GET("/accounts", handlers.GetAccounts)
	r.GET("/accounts/:id", handlers.GetAccountByID)
	r.PUT("/accounts/:id", handlers.UpdateAccount)
	r.DELETE("/accounts/:id", handlers.DeleteAccount)

	return r
}
