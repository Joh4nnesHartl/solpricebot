package commands

import (
	"fmt"
	"strings"
	"time"

	"github.com/Joh4nnesHartl/solpricebot/client"
	embed "github.com/clinet/discordgo-embed"
	"github.com/dustin/go-humanize"
)

// #dc1fff
func makeLastFieldInline(e *embed.Embed) {
	length := len(e.Fields) - 1
	e.Fields[length].Inline = true
}

// #dc1fff
func createEmbedStatistiks(currencyCode string) *embed.Embed {
	data, _ := client.GetSolMarketData(currencyCode)
	date, _ := time.Parse(time.RFC3339, data.AllTimeHighDate)

	e := embed.NewEmbed()

	e.SetColor(0xdc1fff)

	e.SetTitle(fmt.Sprintf("Solana Price: %.2f %s", data.CurrentPrice, strings.ToUpper(currencyCode)))

	makeLastFieldInline(e.AddField("Change 24h", fmt.Sprintf("%.2f", data.Change1Day)+"%"))
	makeLastFieldInline(e.AddField("Change 7d", fmt.Sprintf("%.2f", data.Change7Days)+"%"))

	makeLastFieldInline(e.AddField("Change 30d", fmt.Sprintf("%.2f", data.Change30Days)+"%"))
	makeLastFieldInline(e.AddField("Volume", humanize.CommafWithDigits(data.Volume, 2)))
	makeLastFieldInline(e.AddField("Market Cap", humanize.CommafWithDigits(data.MarketCap, 2)))

	e.AddField("All Time High", fmt.Sprintf("%.2f (%s)", data.AllTimeHigh, date.Format("01.02.2006")))

	return e
}

func createEmbedHelp() *embed.Embed {
	e := embed.NewEmbed()
	e.SetColor(0xdc1fff)
	e.SetTitle("SolanaPriceBot Help")
	e.AddField("Usage", "!sol <currency>")
	e.AddField("Examples", "!sol | !sol eur")
	e.AddField("Solana Donation Address", "DcFptgidUnsjBXjxAwPCoGZTCA3mGjizLm56at9pLNLb")
	e.AddField("API", "CoinGecko")
	e.AddField("Source Code", "https://github.com/Joh4nnesHartl/solpricebot")
	e.SetFooter("by binary#8607")
	return e
}
