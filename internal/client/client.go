package client

import (
	"fmt"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/Joh4nnesHartl/solpricebot/pkg/log"
	coingecko "github.com/superoo7/go-gecko/v3"
)

var (
	currentData CoinData
	Changed     = make(chan struct{})
	mutex       sync.RWMutex
)

var (
	httpClient = &http.Client{
		Timeout: time.Second * 10,
	}

	cryptoClient = coingecko.NewClient(httpClient)

	SupportedCurrencies []string
)

func init() {
	go updateData()

	res, err := cryptoClient.SimpleSupportedVSCurrencies()
	if err != nil {
		log.Fatal(err.Error())
	}

	SupportedCurrencies = []string(*res)

	startDataUpdater()
}

func startDataUpdater() {
	go func() {
		for range time.Tick(time.Second * 21) {
			updateData()
			Changed <- struct{}{}
		}
	}()
}

func updateData() {
	coinData, err := cryptoClient.CoinsID("solana", false, false, true, false, true, false)
	if err != nil {
		log.Error(fmt.Sprintf("error fetching market data from coingecko: %s", err.Error()))
		return
	}

	mutex.Lock()
	currentData = CoinData{
		CurrentPriceMapping: coinData.MarketData.CurrentPrice,
		Change1DayMapping:   coinData.MarketData.PriceChangePercentage24hInCurrency,
		Change7DaysMapping:  coinData.MarketData.PriceChangePercentage7dInCurrency,
		Change30DaysMapping: coinData.MarketData.PriceChangePercentage30dInCurrency,
		VolumeMapping:       coinData.MarketData.TotalVolume,
		MarketCapMapping:    coinData.MarketData.MarketCap,
		AllTimeHighMapping:  coinData.MarketData.ATH,
		AllTimeHighDate:     coinData.MarketData.ATHDate["usd"],
	}
	mutex.Unlock()
}

// GetSolMarketData retrieves all the necessary market data for Solana
func GetSolMarketData(currency string) (CoinDataCurrency, error) {

	cur := strings.ToLower(currency)

	if !IsSupportedCurrency(cur) {
		return CoinDataCurrency{}, fmt.Errorf(`couldnt find data for currency code "%s"`, cur)
	}

	mutex.RLock()
	defer mutex.RUnlock()

	return currentData.ToCurrency(cur), nil

}

func IsSupportedCurrency(currencyCode string) bool {
	return contains(SupportedCurrencies, strings.ToLower(currencyCode))
}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}
