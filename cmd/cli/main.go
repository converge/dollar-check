package main

import (
	"dolar-check/internal/handlers"
	"dolar-check/internal/services"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

// Config ok..
type Config struct {
	target float32
}

func main() {
	httpClient := http.DefaultClient
	quotationService := services.NewQuotationService(httpClient)
	quotationHandler := handlers.NewQuotation(quotationService)

	botToken := os.Getenv("BOT_TOKEN")
	chatId := os.Getenv("CHAT_ID")

	telegramService := services.NewTelegramService(httpClient, botToken, chatId)
	telegramHandler := handlers.NewTelegramHandler(telegramService)

	config := Config{
		target: 5.4,
	}

	counter := 0
	for {
		counter += 1
		quotation, err := quotationHandler.GetQuote("USDBRL=X")
		if err != nil {
			log.Println(err)
		}

		if *quotation > config.target {
			msg := fmt.Sprintf("(USDBRL) target reached! Current price: %f", *quotation)
			err = telegramHandler.SendMessage(msg)
			if err != nil {
				log.Println(err)
			}
		}
		fmt.Printf("current price is (%f)\n", *quotation)
		time.Sleep(5 * time.Second)
	}
}
