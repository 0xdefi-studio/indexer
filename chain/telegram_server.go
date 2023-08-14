package chain

import (
	"fmt"
	models2 "github.com/0xdefi-studio/indexer/models"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strings"
)

func (m *Manager) TelegramServerRun(telegram_bot_key string) {

	bot, err := tgbotapi.NewBotAPI(telegram_bot_key)
	if err != nil {
		m.Sugar.Panic(err)
	}

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil { // If we got a message
			m.Sugar.Infof("[%s] %s", update.Message.From.UserName, update.Message.Text)

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
			//msg.ReplyToMessageID = update.Message.MessageID
			if !strings.HasPrefix(update.Message.Text, "/") {
				continue
			}

			command := update.Message.Command()
			switch command {
			case "subscribe":
				//user := &User{Id: user1.Id}
				//    err = db.Model(user).WherePK().Select()
				//    if err != nil {
				//        panic(err)
				//    }
				chatId := fmt.Sprintf("%v", update.Message.Chat.ID)
				subscription := models2.TelegramChat{ID: chatId}
				err := m.DB.Model(&subscription).WherePK().Select()
				m.Sugar.Debugf("Subscription: %v", subscription)
				if err != nil {
					subscription = models2.TelegramChat{
						ID:           chatId,
						UserName:     update.Message.Chat.UserName,
						Type:         update.Message.Chat.Type,
						IsSubscribed: true,
					}
					_, err := m.DB.Model(&subscription).Insert()
					if err != nil {
						m.Sugar.Panic(err)
						msg.Text = "Error while inserting subscription"
						continue
					}
					msg.Text = "🔔 You subscribed to updates"
				} else {
					subscription.IsSubscribed = true
					_, err := m.DB.Model(&subscription).Where("id = ?", chatId).Update()
					if err != nil {
						m.Sugar.Errorf("Error while updating subscription: %v", err)
						msg.Text = "Error while updating subscription"
						continue
					}
					msg.Text = "You are already subscribed to updates"
				}
			case "unsubscribe":
				chatId := fmt.Sprintf("%v", update.Message.Chat.ID)
				subscription := models2.TelegramChat{ID: chatId}
				err := m.DB.Model(&subscription).WherePK().Select()
				m.Sugar.Debugf("Subscription: %v", subscription)
				if err != nil {
					msg.Text = "You are not subscribed to updates"
					continue
				}
				subscription.IsSubscribed = false
				_, err = m.DB.Model(&subscription).Where("id = ?", chatId).Update()
				if err != nil {
					m.Sugar.Errorf("Error while updating subscription: %v", err)
					msg.Text = "Error while updating subscription"
					continue
				}
				msg.Text = "🚫 You are unsubscribed from updates"
			default:
				msg.Text = "/subscribe - subscribe to updates\n/unsubscribe - unsubscribe from updates"
			}
			if _, err := bot.Send(msg); err != nil {
				m.Sugar.Panic(err)
			}
		}
	}
}
