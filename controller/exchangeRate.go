package controller

import (
	"net/http"

	"github.com/Edu4rdoNeves/Client-Server-API/business"
	"github.com/gin-gonic/gin"
)

func GetLastExchangeRate(context *gin.Context) {
	rateResponse, err := business.GetLastExchangeRate(context)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"Error": "Failed to get a exchange rate" + err.Error(),
		})
	}

	context.JSON(http.StatusOK, rateResponse.Bid)
}
