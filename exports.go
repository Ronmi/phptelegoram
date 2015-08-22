package main

import (
	"log"

	"github.com/tucnak/telebot"
)

import "C"

//export SendMessage
func SendMessage(uid int, msg string) {
	if err := bot.SendMessage(telebot.User{ID: uid}, msg, opt); err != nil {
		log.Printf("Error sending message: %s", err)
	}
}
