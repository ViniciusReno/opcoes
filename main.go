package main

import (
	"fmt"

	"github.com/ViniciusReno/opcoes/data"
	"github.com/ViniciusReno/opcoes/indicators"
	"github.com/ViniciusReno/opcoes/opportunity"
)

func main() {
	symbols := []string{
		"ABEV3.SA", // ABEV3.SA é o símbolo da Ambev na B3
		"B3SA3.SA", // B3SA3.SA é o símbolo da B3 na B3
		"BBAS3.SA", // BBAS3.SA é o símbolo do Banco do Brasil na B3
		"BBDC4.SA", // BBDC4.SA é o símbolo do Bradesco na B3
		"ITUB4.SA", // ITUB4.SA é o símbolo do Itaú Unibanco na B3
		"JBSS3.SA", // JBSS3.SA é o símbolo da JBS na B3
		"MGLU3.SA", // MGLU3.SA é o símbolo da Magazine Luiza na B3
		"PETR4.SA", // PETR4.SA é o símbolo da Petrobras na B3
		"SUZB3.SA", // SUZB3.SA é o símbolo da Suzano na B3
		"TOTS3.SA", // TOTS3.SA é o símbolo da Totvs na B3
		"VALE3.SA", // VALE3.SA é o símbolo da Vale na B3
		"WEGE3.SA", // WEGE3.SA é o símbolo da Weg na B3
	}

	for _, symbol := range symbols {

		historicalData := data.FetchHistoricalData(symbol, "1mo", "1d")

		// Exemplo de extração de preços de fechamento
		prices := data.ExtractClosingPrices(historicalData)

		// Calcular indicadores
		sma50, err := indicators.SMA(prices, 50) // SMA de 50 dias
		if err != nil {
			panic(err)
		}

		ema20 := indicators.EMA(prices, 20)                              // EMA de 20 dias
		macdLine, signalLine := indicators.MACD(prices, 12, 26, 9)       // MACD com configuração padrão
		rsi14 := indicators.RSI(prices, 14)                              // RSI de 14 dias
		upperBand, lowerBand := indicators.BollingerBands(prices, 20, 2) // Bollinger Bands de 20 dias

		// Valores mais recentes
		recentSMA50 := sma50[len(sma50)-1]
		recentEMA20 := ema20[len(ema20)-1]
		recentMACD := macdLine[len(macdLine)-1]
		recentSignal := signalLine[len(signalLine)-1]
		recentRSI := rsi14[len(rsi14)-1]
		recentPrice := prices[len(prices)-1]
		recentUpperBand := upperBand[len(upperBand)-1]
		recentLowerBand := lowerBand[len(lowerBand)-1]

		// Decisão de compra ou venda
		if opportunity.ShouldBuyCall(recentSMA50, recentEMA20, recentMACD, recentSignal, recentRSI, recentLowerBand, recentPrice) {
			fmt.Println("Sinal de compra de CALL")
		} else if opportunity.ShouldBuyPut(recentSMA50, recentEMA20, recentMACD, recentSignal, recentRSI, recentUpperBand, recentPrice) {
			fmt.Println("Sinal de compra de PUT")
		} else {
			fmt.Println("Nenhum sinal claro")
		}
	}
}
