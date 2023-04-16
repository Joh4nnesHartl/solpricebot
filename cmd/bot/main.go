package main

import (
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/Joh4nnesHartl/solpricebot/internal/bot"
	"github.com/Joh4nnesHartl/solpricebot/internal/config"
	"github.com/Joh4nnesHartl/solpricebot/pkg/log"
)

func main() {
	configuration, err := config.Load()
	if err != nil {
		log.Fatal(err.Error())
	}

	dblClient, err := bot.NewDblClient(configuration.DBLToken)
	if err != nil {
		log.Fatal(err.Error())
	}

	manager, err := bot.NewSessionManager(configuration.BotToken)
	if err != nil {
		log.Fatal(err.Error())
	}

	manager.Gateway.GatewayBot()

	bot.StartDBLupdater(dblClient, manager)
	bot.StartStatusUpdater(manager)

	// healthcheck for gcp
	go func() {
		http.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
			w.WriteHeader(http.StatusOK)
		})
		http.ListenAndServe(":3000", nil)
	}()

	log.Info("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	log.Info("Stopping shard manager...")

	manager.Shutdown()

	log.Info("Shard manager stopped. Bot is shut down.")
}
