package services

import (
	"net/http"
	"reflect"
	"testing"
)

func TestNewTelegramService(t *testing.T) {
	var telegramService TelegramService
	type args struct {
		httpClient *http.Client
		botToken   string
		chatId     string
	}
	tests := []struct {
		name string
		args args
		want *TelegramService
	}{
		{
			name: "ok",
			args: args{},
			want: &telegramService,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewTelegramService(tt.args.httpClient, tt.args.botToken, tt.args.chatId); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTelegramService() = %v, want %v", got, tt.want)
			}
		})
	}
}
