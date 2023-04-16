package command

import (
	"fmt"
	"strings"
	"time"

	embed "github.com/clinet/discordgo-embed"
	"github.com/dustin/go-humanize"

	"github.com/Joh4nnesHartl/solpricebot/internal/client"
)

const color = 0xdc1fff

func makeLastFieldInline(e *embed.Embed) {
	length := len(e.Fields) - 1
	e.Fields[length].Inline = true
}

func createEmbedStatistiks(currencyCode string) *embed.Embed {
	if !client.IsSupportedCurrency(currencyCode) {
		return createEmbedError("invalid currency code", fmt.Sprintf("currency code '%s' is not valid", currencyCode))
	}

	data, _ := client.GetSolMarketData(currencyCode)
	date, _ := time.Parse(time.RFC3339, data.AllTimeHighDate)

	e := embed.NewEmbed()

	e.SetColor(color)

	e.SetTitle(fmt.Sprintf("Solana Price: %.2f %s", data.CurrentPrice, strings.ToUpper(currencyCode)))

	makeLastFieldInline(e.AddField("Change 24h", fmt.Sprintf("%.2f", data.Change1Day)+"%"))
	makeLastFieldInline(e.AddField("Change 7d", fmt.Sprintf("%.2f", data.Change7Days)+"%"))

	makeLastFieldInline(e.AddField("Change 30d", fmt.Sprintf("%.2f", data.Change30Days)+"%"))
	makeLastFieldInline(e.AddField("Volume", humanize.CommafWithDigits(data.Volume, 2)))
	makeLastFieldInline(e.AddField("Market Cap", humanize.CommafWithDigits(data.MarketCap, 2)))

	e.AddField("All Time High", fmt.Sprintf("%.2f (%s)", data.AllTimeHigh, date.Format("01.02.2006")))

	return e
}

func createEmbedError(title string, errorMessage string) *embed.Embed {
	e := embed.NewEmbed()
	e.SetColor(color)
	e.SetTitle("Error")

	e.AddField(title, errorMessage)

	return e
}
