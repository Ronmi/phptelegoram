package main

import (
	"encoding/json"
	"fmt"
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

//export TimedTask
func TimedTask(milisec int, data string) {
	var nanosec int64 = int64(milisec) * 1000000
	d := fmt.Sprintf("%s", data)
	pipe <- task{nanosec, d}
}
