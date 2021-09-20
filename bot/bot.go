package bot

import (
	"fmt"
	"os"

	"github.com/Joh4nnesHartl/solpricebot/client"
	"github.com/Joh4nnesHartl/solpricebot/log"
	"github.com/bwmarrin/discordgo"

	"github.com/Joh4nnesHartl/solpricebot/bot/events"
)

var (
	// Session representates a discord session
	Session *discordgo.Session
)

func init() {
	Session = createSession()

	events.AssignSession(Session)

	startStatusUpdater(Session)
}

func startStatusUpdater(session *discordgo.Session) {
	go func() {
		for range client.Changed {
			coinData, _ := client.GetSolMarketData("usd")

			status := fmt.Sprintf("$%.2f", coinData.CurrentPrice)
			session.UpdateGameStatus(0, status)
			log.Infof("updated bot status to: %s", status)
		}
	}()
}

func createSession() *discordgo.Session {
	token := os.Getenv("BOT_TOKEN")
	if token == "" {
		log.Fatal("error no token set as env")
	}

	session, err := discordgo.New("Bot " + token)
	if err != nil {
		log.Fatal("error creating discord session out of Token\n")
	}

	return session
}
