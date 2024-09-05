package repository

import (
	"context"
	"log"
	"time"

	"github.com/Edu4rdoNeves/Client-Server-API/database"
	"github.com/Edu4rdoNeves/Client-Server-API/model"
	"github.com/gin-gonic/gin"
)

func GetLastExchangeRate(ctx *gin.Context, exchangeRate *model.ExchangeRate) (*model.ExchangeRate, error) {
	db := database.GetDatabase()

	createCtx, cancel := context.WithTimeout(ctx.Request.Context(), 1*time.Second)
	defer cancel()

	err := db.WithContext(createCtx).Create(exchangeRate).Error
	if err != nil {
		log.Printf("Context timeout exceeded 1 second. Error: % v", err)
		return nil, err
	}

	queryCtx, queryCancel := context.WithTimeout(ctx.Request.Context(), 500*time.Millisecond)
	defer queryCancel()

	var lastExchangeRate model.ExchangeRate
	err = db.WithContext(queryCtx).Order("created_at desc").First(&lastExchangeRate).Error
	if err != nil {
		log.Printf("Context timeout exceeded 500 ms. Error: % v", err)
		return nil, err
	}

	return &lastExchangeRate, nil
}
