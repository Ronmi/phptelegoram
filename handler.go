package main

import (
	"encoding/json"
	"log"

	"github.com/tucnak/telebot"
)

func handleMessages(ch chan telebot.Message) {
	for msg := range ch {
		data, err := json.Marshal(msg)
		if err != nil {
			log.Printf("Error encoding message to json: %s", err)
			continue
		}
		script := NewContext()
		script.Startup()
		script.Var("message", string(data))
		_ = script.Exec(phpFile)
		script.Close()
	}
}
