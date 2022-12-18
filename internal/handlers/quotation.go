package handlers

import (
	"dolar-check/internal/services"
)

type QuotationHandler struct {
	QuotationService services.QuotationService
}

func NewQuotation(quotationService *services.QuotationService) *QuotationHandler {
	return &QuotationHandler{
		QuotationService: *quotationService,
	}
}

func (quotationHandler QuotationHandler) GetQuote(symbol string) (*float32, error) {
	quotation, err := quotationHandler.QuotationService.GetQuote(symbol)
	if err != nil {
		return nil, err
	}
	return quotation, nil
}
