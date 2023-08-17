package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/big"
	"net/http"
)

type Message struct {
	ChatID      int64  `json:"chat_id"`
	Text        string `json:"text"`
	ParseMode   string `json:"parse_mode"`
	ReplyMarkup any    `json:"reply_markup"`
}

func SendTelegramEndTx(telegram_key string, chat_id int64, userAddress string, keys *big.Int, amount *big.Int, createdTime uint64, PotAmount *big.Int, roundEndTime *big.Int) {
	url := "https://api.telegram.org/bot" + telegram_key + "/sendMessage"
	message := Message{
		ChatID: chat_id,
		Text: fmt.Sprintf("🔑  %v Keys Bought!\n\n🎰 Pot: %v MNT\n⏳ Countdown: %v\n🤑 Buyer: %v\n💰 Cost: %v MNT\n\n",
			new(big.Float).Quo(new(big.Float).SetInt(keys), new(big.Float).SetInt(divider)).String(),
			new(big.Float).Quo(new(big.Float).SetInt(PotAmount), new(big.Float).SetInt(divider)).String(),
			DurationToTimeFormat(roundEndTime),
			ShortenAddress(userAddress),
			new(big.Float).Quo(new(big.Float).SetInt(amount), new(big.Float).SetInt(divider)).String()),
		ParseMode: "Markdown",
		ReplyMarkup: map[string]any{"inline_keyboard": []any{
			[]any{
				map[string]string{
					"text": "Buy now to escape with pot!",
					"url":  "https://mantleasylum.app/",
				},
			},
		},
		},
	}
	payload, err := json.Marshal(message)
	if err != nil {
		log.Println(err)
		return
	}
	response, err := http.Post(url, "application/json", bytes.NewBuffer(payload))
	if err != nil {
		log.Println(err)
		return
	}
	defer func(body io.ReadCloser) {
		if err := body.Close(); err != nil {
			log.Println("failed to close response body")
		}
	}(response.Body)
	if response.StatusCode != http.StatusOK {
		log.Printf("failed to send successful request. Status was %v", response.Status)
		return
	}
	return
}
