package showdata

import (
	"fmt"
	"time"

	"github.com/ViniciusReno/opcoes/models"
)

func ShowMe(data models.StockData, sma []float64, ema []float64, rsi []float64, support float64, resistance float64, macd []float64, signal []float64) {
	fmt.Println("Symbol:", data.Chart.Result[0].Meta.Symbol)
	fmt.Println("Market Price:", data.Chart.Result[0].Meta.RegularMarketPrice)

	timestamps := data.Chart.Result[0].Timestamp
	closes := data.Chart.Result[0].Indicators.Quote[0].Close
	volumes := data.Chart.Result[0].Indicators.Quote[0].Volume

	for i := 0; i < len(timestamps); i++ {
		date := time.Unix(timestamps[i], 0).Format("2006-01-02")
		fmt.Printf("Date: %s, Close: %.2f, Volume: %d\n", date, closes[i], volumes[i])
	}

	fmt.Println("")

	for i := 0; i < len(sma); i++ {
		fmt.Printf("SMA[%d]: %.2f\n", i, sma[i])
	}

	fmt.Println("")

	for i := 0; i < len(ema); i++ {
		fmt.Printf("EMA[%d]: %.2f\n", i, ema[i])
	}

	fmt.Println("")

	for i := 0; i < len(rsi); i++ {
		fmt.Printf("RSI[%d]: %.2f\n", i, rsi[i])
	}

	fmt.Println("")

	fmt.Printf("Support: %.2f\n", support)
	fmt.Printf("Resistance: %.2f\n", resistance)

	fmt.Println("")

	for i := 0; i < len(macd); i++ {
		fmt.Printf("MACD[%d]: %.2f\n", i, macd[i])
	}

	fmt.Println("")

	for i := 0; i < len(signal); i++ {
		fmt.Printf("Signal[%d]: %.2f\n", i, signal[i])
	}

}
