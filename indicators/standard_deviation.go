package indicators

import "math"

// Calcula o Desvio Padrão
func StdDev(prices []float64, period int) []float64 {
	sma, err := SMA(prices, period)
	if err != nil {
		panic(err)
	}

	stddev := make([]float64, len(sma))
	for i := range sma {
		variance := 0.0
		for j := i; j < i+period; j++ {
			variance += math.Pow(prices[j]-sma[i], 2)
		}
		stddev[i] = math.Sqrt(variance / float64(period))
	}
	return stddev
}
