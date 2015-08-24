package main

import (
	"encoding/json"
	"log"

	"github.com/tucnak/telebot"
)

import "C"

//export SendMessage
func SendMessage(uid int, msg string, options string) {
	var real telebot.SendOptions
	if err := json.Unmarshal([]byte(options), &real); err != nil {
		log.Printf("Cannot unmarshal %s: %s", options, err)
		real = *opt
	}
	log.Printf("Options: %#v", real)
	if err := bot.SendMessage(telebot.User{ID: uid}, msg, &real); err != nil {
		log.Printf("Error sending message: %s", err)
	}
}
