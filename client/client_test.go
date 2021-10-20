package client

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func init() {
	updateData()
}

func TestGetSolMarketData(t *testing.T) {

	coinData, err := GetSolMarketData("usd")
	require.NoError(t, err)

	assert.NotZero(t, coinData.Change1Day)
	assert.NotZero(t, coinData.Change7Days)
	assert.NotZero(t, coinData.Change30Days)
	assert.NotZero(t, coinData.Volume)
	assert.NotZero(t, coinData.MarketCap)
	assert.NotZero(t, coinData.AllTimeHigh)
	assert.NotZero(t, coinData.AllTimeHighDate)
}

func TestGetSolMarketDataUpperCaseCur(t *testing.T) {
	coinData, err := GetSolMarketData("USD")
	require.NoError(t, err)

	assert.NotZero(t, coinData.Change1Day)
	assert.NotZero(t, coinData.Change7Days)
	assert.NotZero(t, coinData.Change30Days)
	assert.NotZero(t, coinData.Volume)
	assert.NotZero(t, coinData.MarketCap)
	assert.NotZero(t, coinData.AllTimeHigh)
	assert.NotZero(t, coinData.AllTimeHighDate)
}

func TestGetSolMarketDataNotSupportedCur(t *testing.T) {
	_, err := GetSolMarketData("sub")
	require.Error(t, err)

	assert.Equal(t, `couldnt find data for currency code "sub"`, err.Error())
}
