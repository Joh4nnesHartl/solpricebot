package bot

import (
	"fmt"
	"os"
	"time"

	"github.com/Joh4nnesHartl/solpricebot/client"
	"github.com/Joh4nnesHartl/solpricebot/log"
	"github.com/bwmarrin/discordgo"
	"github.com/top-gg/go-dbl"

	"github.com/Joh4nnesHartl/solpricebot/bot/events"
)

const botID = 888052145734172732

var (
	// Session representates a discord session
	Session *discordgo.Session

	// dblClient representates an top.gg api client
	dblClient *dbl.Client
)

func init() {
	Session = createSession()
	dblClient = createDBLclient()

	events.AssignSession(Session)

	startStatusUpdater(Session)
	startDBLupdater(dblClient, Session)
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

func startDBLupdater(dblClient *dbl.Client, session *discordgo.Session) {
	time.AfterFunc(20*time.Second, func() { updateDBLdata(dblClient, session) })

	go func() {
		for range time.Tick(30 * time.Minute) {
			err := updateDBLdata(dblClient, session)
			if err != nil {
				log.Errorf("error updating top.gg data: %s", err.Error())
			}
		}
	}()
}

func updateDBLdata(dblClient *dbl.Client, session *discordgo.Session) error {
	session.State.RLock()
	serverCount := len(session.State.Ready.Guilds)
	session.State.RUnlock()

	log.Infof("Updating server top.gg server count to %d", serverCount)
	return dblClient.PostBotStats(fmt.Sprint(botID), &dbl.BotStatsPayload{
		Shards: []int{serverCount},
	})
}

func createSession() *discordgo.Session {
	token := os.Getenv("BOT_TOKEN")
	if token == "" {
		log.Fatal("error no bot token set as env")
	}

	session, err := discordgo.New("Bot " + token)
	if err != nil {
		log.Fatal("error creating discord session out of Token\n")
	}

	return session
}

func createDBLclient() *dbl.Client {
	token := os.Getenv("DBL_TOKEN")
	if token == "" {
		log.Fatal("error no dbl token set as env")
	}

	dblClient, err := dbl.NewClient(token)
	if err != nil {
		log.Fatal("error creating dbl client out of token\n")
	}

	return dblClient
}
