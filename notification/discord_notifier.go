package cmd

import (
	"fmt"
	"github.com/gtuk/discordwebhook"
	"log"
	"math/big"
	"time"
)

type WebhookData struct {
	UserAddress string `json:"UserAddress"`
	Keys        string `json:"Keys"`
	Amount      string `json:"Amount"`
}

func ShortenAddress(address string) string {
	if len(address) < 10 {
		return address // Not a valid address, return as is
	}

	// Get the first 6 characters
	prefix := address[0:6]
	// Get the last 4 characters
	suffix := address[len(address)-4:]

	return fmt.Sprintf("%s...%s", prefix, suffix)
}

var divider = new(big.Int).Exp(big.NewInt(10), big.NewInt(18), nil)

//var divider20 = new(big.Int).Exp(big.NewInt(10), big.NewInt(20), nil)

func SendDiscordEndTx(discord_web_hook string, userAddress string, keys *big.Int, amount *big.Int, createdTime uint64) {
	data := WebhookData{
		UserAddress: ShortenAddress(userAddress),
		Keys:        new(big.Int).Div(keys, divider).String(),
		Amount:      new(big.Float).Quo(new(big.Float).SetInt(amount), new(big.Float).SetInt(divider)).String(),
	}

	//make human readable time utc = 0
	t := time.Unix(int64(createdTime), 0)
	content := fmt.Sprintf(":moneybag: UserAddress: %s\nKeys: %s\nAmount: %s\nTime created: %v\n\n",
		data.UserAddress, data.Keys, data.Amount, t)

	var username = "mantleasylum:alert"
	message := discordwebhook.Message{
		Username: &username,
		Content:  &content,
	}

	err := discordwebhook.SendMessage(discord_web_hook, message)
	if err != nil {
		log.Println(err)
		return
	}
}
