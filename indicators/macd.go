package indicators

// Função para calcular o MACD e a linha de sinal
func MACD(prices []float64, shortPeriod, longPeriod, signalPeriod int) ([]float64, []float64) {
	shortEMA := EMA(prices, shortPeriod) // EMA de curto prazo
	longEMA := EMA(prices, longPeriod)   // EMA de longo prazo

	macdLine := make([]float64, len(prices))
	for i := range macdLine {
		macdLine[i] = shortEMA[i] - longEMA[i] // Linha MACD é a diferença entre as duas EMAs
	}

	signalLine := EMA(macdLine, signalPeriod) // Linha de sinal é a EMA da linha MACD

	return macdLine, signalLine
}
