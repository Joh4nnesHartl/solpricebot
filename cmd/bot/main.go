package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/Joh4nnesHartl/solpricebot/bot"
	"github.com/Joh4nnesHartl/solpricebot/log"
)

func main() {
	// opening websocket connection
	err := bot.Session.Open()
	if err != nil {
		log.Fatal("error opening connection,", err)
	}

	log.Info("Bot is now running.  Press CTRL-C to exit.")

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	// Cleanly close down the Discord session.
	bot.Session.Close()
}
