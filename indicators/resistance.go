package indicators

// Encontra o menor e o maior preço em um período
func SupportResistance(prices []float64, period int) (support float64, resistance float64) {
	support, resistance = prices[0], prices[0]

	for i := 1; i < period; i++ {
		if prices[i] < support {
			support = prices[i]
		}
		if prices[i] > resistance {
			resistance = prices[i]
		}
	}
	return support, resistance
}
