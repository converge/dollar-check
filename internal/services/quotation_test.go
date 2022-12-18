package services

import (
	"net/http"
	"reflect"
	"testing"
)

func TestNewQuotationService(t *testing.T) {

	var quotationService *QuotationService

	type args struct {
		httpClient *http.Client
	}
	tests := []struct {
		name string
		args args
		want *QuotationService
	}{
		{
			name: "ok1",
			args: args{},
			want: quotationService,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewQuotationService(tt.args.httpClient); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewQuotationService() = %v, want %v", got, tt.want)
			}
		})
	}
}
