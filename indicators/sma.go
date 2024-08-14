package indicators

import "fmt"

// Função para calcular a SMA com ajuste dinâmico do período
func SMA(prices []float64, period int) ([]float64, error) {
	if len(prices) < period {
		fmt.Printf("Aviso: O período %d é maior que o número de preços disponíveis (%d). Ajustando o período para %d.\n", period, len(prices), len(prices))
		period = len(prices)
	}

	sma := make([]float64, len(prices)-period+1)
	for i := 0; i <= len(prices)-period; i++ {
		sum := 0.0
		for j := 0; j < period; j++ {
			sum += prices[i+j]
		}
		sma[i] = sum / float64(period)
	}

	return sma, nil
}
