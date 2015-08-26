package main

import (
	"encoding/json"
	"log"
	"os"
	"time"

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
		if _, err := os.Stat(iniFile); err == nil {
			script.IniFile(iniFile)
		}
		script.Startup()
		script.Var("message_json", string(data))
		script.Var("entry_file", phpFile)
		_ = script.Exec("message.php")
		script.Close()
	}
}

func handleTask() {
	for tk := range pipe {
		go func(t task) {
			time.Sleep(time.Duration(t.nanosec))
			script := NewContext()
			if _, err := os.Stat(iniFile); err == nil {
				script.IniFile(iniFile)
			}
			script.Startup()
			script.Var("data", t.data)
			script.Var("entry_file", phpFile)
			_ = script.Exec("timedtask.php")
			script.Close()
		}(tk)
	}
}
