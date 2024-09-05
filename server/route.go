package server

import (
	"github.com/Edu4rdoNeves/Client-Server-API/controller"
	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {

	main := router.Group("api/v1")
	{
		exchangeRateQuote := main.Group("quote")
		{
			exchangeRateQuote.GET("/", controller.GetLastExchangeRate)
		}

	}
}
