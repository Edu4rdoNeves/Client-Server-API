package business

import (
	"log"
	"os"
	"text/template"

	"github.com/Edu4rdoNeves/Client-Server-API/model"
	"github.com/Edu4rdoNeves/Client-Server-API/repository"
	"github.com/Edu4rdoNeves/Client-Server-API/service"
	"github.com/gin-gonic/gin"
)

func GetLastExchangeRate(ctx *gin.Context) (*model.ExchangeRate, error) {
	exchangeRate, err := service.ExchangeRate()
	if err != nil {
		return nil, err
	}

	exchangeRateResponse, err := repository.GetLastExchangeRate(ctx, exchangeRate)
	if err != nil {
		return nil, err
	}

	err = SaveAtTxtFile(exchangeRate)
	if err != nil {
		return nil, err
	}

	return exchangeRateResponse, nil
}

func SaveAtTxtFile(exchangeRate *model.ExchangeRate) error {

	tmpl := `
	Bid: {{.Bid}}
	CreatedAt: {{.CreatedAt}}
	`

	t := template.Must(template.New("exchangeRateTemplate").Parse(tmpl))

	file, err := os.Create("rate.txt")
	if err != nil {
		log.Printf("Error to create file: %v", err)
		return nil
	}
	defer file.Close()

	err = t.Execute(file, exchangeRate)
	if err != nil {
		log.Printf("Error to execute a template: %v", err)
		return nil
	}

	log.Printf("Data saved in rate.txt file")
	return nil
}
