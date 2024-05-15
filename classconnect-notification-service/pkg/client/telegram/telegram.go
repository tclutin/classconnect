package telegram

import (
	"fmt"
	"net/http"
	url2 "net/url"
)

type Client interface {
	Send(chatId uint64, message string) error
}

type Telegram struct {
	Token      string
	httpClient http.Client
}

func NewClient(token string) *Telegram {
	return &Telegram{
		httpClient: http.Client{},
		Token:      token,
	}
}

func (t *Telegram) Send(chatId uint64, message string) error {
	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage?chat_id=%v&text=%s", t.Token, chatId, url2.QueryEscape(message))

	resp, err := t.httpClient.Post(url, "", nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("telegram API returned non-200 status code: %d", resp.StatusCode)
	}

	return nil
}
