package indicators

// Calcula as Bollinger Bands
func BollingerBands(prices []float64, period int, numStdDev float64) ([]float64, []float64) {
	sma, err := SMA(prices, period)
	if err != nil {
		panic(err)
	}
	stddev := StdDev(prices, period)
	upperBand := make([]float64, len(sma))
	lowerBand := make([]float64, len(sma))

	for i := range sma {
		upperBand[i] = sma[i] + numStdDev*stddev[i]
		lowerBand[i] = sma[i] - numStdDev*stddev[i]
	}
	return upperBand, lowerBand
}
