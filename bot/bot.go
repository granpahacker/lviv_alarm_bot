// Copyright (C) Oleg Lysiak - All Rights Reserved
package bot

import (
	"fmt"
	"log"
	"lviv-alarm-bot/config"
	"net/http"
	"time"

	// discord bot library

	"github.com/PuerkitoBio/goquery"
	"github.com/bwmarrin/discordgo"
)

// main bot variables

var BotId string
var goBot *discordgo.Session

// alert status info

var (
	currentStatus string
	title         string
)

// start func

func Start() {
	goBot, err := discordgo.New("Bot " + config.Token)

	// check for errors

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	u, err := goBot.User("@me")

	// check for errors

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	BotId = u.ID

	// handler for messages

	goBot.AddHandler(messageHandler)

	err = goBot.Open()

	// check for errors

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Lviv Alarm Bot is online!")
}

func telegramParser(tick time.Time) {
	// link for parser

	webPage := "https://t.me/s/LvivAlarm"
	resp, err := http.Get(webPage)

	// check for errors

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	// check for errors

	if resp.StatusCode != 200 {
		log.Fatalf("failed to fetch data: %d %s", resp.StatusCode, resp.Status)
	}

	// check for errors

	doc, err := goquery.NewDocumentFromReader(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	title = doc.Find(".tgme_widget_message_text").Last().Text()

	time.Sleep(30 * time.Second)

}

// bot messages handler

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == BotId {
		return
	}
	var messageText = func(text string) string {
		_, _ = s.ChannelMessageSend(m.ChannelID, text)
		return text
	}

	// start command

	if m.Content == "Старт" {
		// send succes text

		messageText("Вітання! Бот запустивсі!")
		// loop for parser

		for t := range time.Tick(30 * time.Second) {
			go telegramParser(t)

			// checker for updates

			if title != currentStatus {
				messageText("@everyone @everyone @everyone")
				messageText(title)
				currentStatus = title
			}
		}
	} else if m.Content == "Статус" {
		messageText("Статус повітряної тривоги на даний момент:")
		messageText(title)
	}
}
