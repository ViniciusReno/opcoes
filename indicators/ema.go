package indicators

// Calcula a Média Móvel Exponencial
func EMA(prices []float64, period int) []float64 {
	ema := make([]float64, len(prices))
	multiplier := 2.0 / (float64(period) + 1.0)
	ema[0] = prices[0] // A primeira EMA é igual ao primeiro preço

	for i := 1; i < len(prices); i++ {
		ema[i] = ((prices[i] - ema[i-1]) * multiplier) + ema[i-1]
	}
	return ema
}
