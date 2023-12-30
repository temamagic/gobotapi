# Golang Telegram Bot API 

This is a custom fork of go-telegram-bot-api.

This fork contains both core updates and patches specific to me.

The purpose of this fork is to have the necessary set of methods in my projects and not wait for the original author to integrate updates.

Based on:
- https://github.com/go-telegram-bot-api/telegram-bot-api
- https://github.com/OvyFlash/telegram-bot-api

Please, read original [readme](https://github.com/go-telegram-bot-api/telegram-bot-api/blob/master/README.md) first.

## ! Note !
The original library is updated rarely and little.
I do not guarantee that this fork will be updated more often.
I only update it when there is a need for it for me, or a good update has been submitted through a pull request and reviewed by someone.

As a basis, I also take another [repository](https://github.com/OvyFlash/telegram-bot-api), which is updated quite often (at the time of writing this note), thanks to him!

## Install & Update

`go get -u github.com/temamagic/gobotapi`.

## Example

```
package main

import (
	"log"

	tgbotapi "github.com/temamagic/gobotapi"
)

func main() {
	bot, err := tgbotapi.NewBotAPI("MyAwesomeBotToken")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil { // If we got a message
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
			msg.ReplyToMessageID = update.Message.MessageID

			bot.Send(msg)
		}
	}
}
```