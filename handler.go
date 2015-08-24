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
		script.Var("message_json", string(data))
		script.Var("entry_file", phpFile)
		_ = script.Exec("lib.php")
		script.Close()
	}
}
