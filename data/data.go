package data

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Exemplo de função para buscar dados históricos do Yahoo Finance
func FetchHistoricalData(symbol, period, interval string) []byte {
	url := fmt.Sprintf("https://query1.finance.yahoo.com/v8/finance/chart/%s?range=%s&interval=%s", symbol, period, interval)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Erro ao buscar dados históricos: %v", err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Erro ao ler a resposta da API: %v", err)
	}

	return body
}

// Exemplo de função para extrair preços de fechamento do JSON da API
func ExtractClosingPrices(data []byte) []float64 {
	var result map[string]interface{}
	json.Unmarshal(data, &result)

	prices := []float64{}
	timestamps := result["chart"].(map[string]interface{})["result"].([]interface{})[0].(map[string]interface{})["timestamp"].([]interface{})
	closes := result["chart"].(map[string]interface{})["result"].([]interface{})[0].(map[string]interface{})["indicators"].(map[string]interface{})["quote"].([]interface{})[0].(map[string]interface{})["close"].([]interface{})

	for i := range timestamps {
		if closes[i] != nil {
			prices = append(prices, closes[i].(float64))
		}
	}
	return prices
}
