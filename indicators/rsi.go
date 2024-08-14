package indicators

// Calcula o Índice de Força Relativa
func RSI(prices []float64, period int) []float64 {
	rsi := make([]float64, len(prices)-period)
	gains, losses := 0.0, 0.0

	// Calcula os ganhos e perdas iniciais
	for i := 1; i <= period; i++ {
		change := prices[i] - prices[i-1]
		if change > 0 {
			gains += change
		} else {
			losses -= change
		}
	}

	avgGain := gains / float64(period)
	avgLoss := losses / float64(period)
	rs := avgGain / avgLoss
	rsi[0] = 100 - (100 / (1 + rs))

	// Calcula o RSI para os períodos subsequentes
	for i := period + 1; i < len(prices); i++ {
		change := prices[i] - prices[i-1]
		if change > 0 {
			gains = change
			losses = 0
		} else {
			gains = 0
			losses = -change
		}

		avgGain = ((avgGain * float64(period-1)) + gains) / float64(period)
		avgLoss = ((avgLoss * float64(period-1)) + losses) / float64(period)
		rs = avgGain / avgLoss
		rsi[i-period] = 100 - (100 / (1 + rs))
	}

	return rsi
}
