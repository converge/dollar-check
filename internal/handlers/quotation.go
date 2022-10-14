package handlers

import (
	"dolar-check/internal/services"
)

type QuotationHandler struct {
	symbol            string
	regularMarketTime float32
	QuotationService  services.QuotationService
}

func NewQuotation(quotationService *services.QuotationService) *QuotationHandler {
	return &QuotationHandler{
		QuotationService: *quotationService,
	}
}

func (quotationHandler QuotationHandler) GetQuote(symbol string) error {
	err := quotationHandler.QuotationService.GetQuote(symbol)
	if err != nil {
		return err
	}
	return nil
}
