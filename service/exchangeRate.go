package service

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Edu4rdoNeves/Client-Server-API/model"
)

func ExchangeRate() (*model.ExchangeRate, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	request, err := http.NewRequestWithContext(ctx, http.MethodGet, os.Getenv("EXCHANGE_RATE_URL"), nil)
	if err != nil {
		log.Printf("Context timeout exceeded 1 second. Error: % v", err)
		return nil, err
	}

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Println("Fail to do request.")
		return nil, err
	}
	defer response.Body.Close()

	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		log.Println("Fail to bind a response.")
		return nil, err
	}

	var rate model.ExchangeRate
	err = json.Unmarshal(responseData, &rate)
	if err != nil {
		log.Println("Fail to Unmarshal a response.")
		return nil, err
	}

	return &rate, nil
}
