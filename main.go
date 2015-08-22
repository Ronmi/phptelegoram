package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"

	"github.com/tucnak/telebot"
)

var (
	tokenFile string
	workers   uint
	token     string
	phpFile   string
	bot       *telebot.Bot
	opt       *telebot.SendOptions
)

func init() {
	opt = new(telebot.SendOptions)
	flag.StringVar(&tokenFile, "t", "token", "File contains bot token")
	flag.StringVar(&phpFile, "php", "entry.php", "PHP entry file, you can access JSON encoded message data in $message")
	flag.UintVar(&workers, "w", 1, "Run `N` goroutines to process message, must greater than 0")
	flag.Parse()

	data, err := ioutil.ReadFile(tokenFile)
	if err != nil {
		log.Fatalf("Cannot read token from file[%s]: %s", tokenFile, err)
	}
	token = strings.TrimSpace(string(data))

	if workers < 1 {
		flag.PrintDefaults()
		os.Exit(1)
	}

	if _, err := os.Stat(phpFile); err != nil {
		log.Fatalf("PHP entry file %s error: %s", phpFile, err)
	}

	bot, err = telebot.NewBot(token)
	if err != nil {
		log.Fatalf("Cannot start telebot: %s", err)
	}
}

func main() {
	ch := make(chan telebot.Message)
	bot.Listen(ch, 30*time.Second)

	for i := uint(1); i < workers; i++ {
		go handleMessages(ch)
	}

	handleMessages(ch)
}
