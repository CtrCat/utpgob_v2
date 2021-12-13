package main

import (
	"log"

	"encoding/json"
	"fmt"
	"os"

	utopiago "github.com/Sagleft/utopialib-go"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Config struct {
	TelegramBotToken string
	UtpToken         string
	UtpPort          int
}

func main() {

	//read congig
	file, _ := os.Open("config.json")
	decoder := json.NewDecoder(file)
	configuration := Config{}
	err := decoder.Decode(&configuration)
	if err != nil {
		log.Panic(err)
	}
	fmt.Println(configuration.TelegramBotToken)
	fmt.Println(configuration.UtpToken)
	fmt.Println(configuration.UtpPort)

	// bot-token

	bot, err := tgbotapi.NewBotAPI(configuration.TelegramBotToken)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	// ini channel
	var ucfg tgbotapi.UpdateConfig = tgbotapi.NewUpdate(0)
	ucfg.Timeout = 60
	updatesChann := bot.GetUpdatesChan(ucfg)

	//utp
	client := utopiago.UtopiaClient{
		Protocol: "http",
		Token:    configuration.UtpToken,
		Host:     "127.0.0.1",
		Port:     configuration.UtpPort,
	}

	fmt.Println(client.GetBalance())

	//json goutpbot

	// update
	for {
		select {
		case update := <-updatesChann:
			// User bot
			UserName := update.Message.From.UserName

			// ID chat.

			ChatID := update.Message.Chat.ID

			// Text massage user
			Text := update.Message.Text

			log.Printf("[%s] %d %s", UserName, ChatID, Text)

			// reply
			reply := "Ok"
			// create massage
			msg := tgbotapi.NewMessage(ChatID, reply)
			// send
			bot.Send(msg)

			switch Text {

			case "/GetBalance":

				//UUSD

				fmt.Println("/GetBalanceUSD")

				balusd, err := client.GetUUSDBalance()
				if err != nil {
					panic(err.Error())

				}

				/*	jsonBytes, err := json.Marshal(balcrp)

					if err != nil {
						log.Println(err)
						return
					}*/

				x1 := fmt.Sprintf("%f", balusd)

				//x := fmt.Sprintf("% s", balcrp)

				fmt.Println("chek")
				fmt.Println(client.GetUUSDBalance())
				fmt.Println("chek")

				reply := x1
				msg := tgbotapi.NewMessage(ChatID, reply+"-- UUSD balance")

				bot.Send(msg)

			case "/GetSystemInfo":

				fmt.Println("/GetSystemInfo")

				sustim, err := client.GetSystemInfo()
				if err != nil {
					panic(err.Error())

				}

				jsonBytes, err := json.Marshal(sustim)

				if err != nil {
					log.Println(err)
					return
				}

				x1 := fmt.Sprintf("% s", jsonBytes)

				fmt.Println("chek")
				fmt.Println(client.GetSystemInfo())
				fmt.Println("chek")

				reply := x1
				msg := tgbotapi.NewMessage(ChatID, "\n \n-- SYSTEM \n\n"+reply+"\n \n-- SYSTEM")

				bot.Send(msg)

			default:

				fmt.Println("commands")

				reply := "Commands:\n /GetBalance \n /GetSystemInfo"
				msg := tgbotapi.NewMessage(ChatID, reply)

				bot.Send(msg)

			}

		}

	}
}
