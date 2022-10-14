package main

import (
	"dolar-check/internal/handlers"
	"dolar-check/internal/services"
	"log"
	"net/http"
	"os"
)

func main() {
	httpClient := http.DefaultClient
	quotationService := services.NewQuotationService(httpClient)
	//quotationHandler := handlers.NewQuotation(quotationService)
	_ = handlers.NewQuotation(quotationService)

	//counter := 0
	//for {
	//	counter += 1
	//	err := quotationHandler.GetQuote("USDBRL=X")
	//	if err != nil {
	//		log.Println(err)
	//	}
	//
	//	fmt.Println("sleeping...", counter)
	//	time.Sleep(1 * time.Second)
	//}

	botToken := os.Getenv("BOT_TOKEN")
	chatId := os.Getenv("CHAT_ID")

	telegramService := services.NewTelegramService(httpClient, botToken, chatId)
	telegramHandler := handlers.NewTelegramHandler(telegramService)
	err := telegramHandler.SendMessage("test123")
	if err != nil {
		log.Println(err)
	}
}
