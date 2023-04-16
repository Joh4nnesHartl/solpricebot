package client

type CoinData struct {
	CurrentPriceMapping map[string]float64
	Change1DayMapping   map[string]float64
	Change7DaysMapping  map[string]float64
	Change30DaysMapping map[string]float64
	VolumeMapping       map[string]float64
	MarketCapMapping    map[string]float64
	AllTimeHighMapping  map[string]float64
	AllTimeHighDate     string
}

type CoinDataCurrency struct {
	CurrentPrice    float64
	Change1Day      float64
	Change7Days     float64
	Change30Days    float64
	Volume          float64
	MarketCap       float64
	AllTimeHigh     float64
	AllTimeHighDate string
}

func (c CoinData) ToCurrency(currencyCode string) CoinDataCurrency {
	return CoinDataCurrency{
		CurrentPrice:    c.CurrentPriceMapping[currencyCode],
		Change1Day:      c.Change1DayMapping[currencyCode],
		Change7Days:     c.Change7DaysMapping[currencyCode],
		Change30Days:    c.Change30DaysMapping[currencyCode],
		Volume:          c.VolumeMapping[currencyCode],
		MarketCap:       c.MarketCapMapping[currencyCode],
		AllTimeHigh:     c.AllTimeHighMapping[currencyCode],
		AllTimeHighDate: c.AllTimeHighDate,
	}
}
