package opportunity

func ShouldBuyCall(sma50, ema20, macd, signal, rsi, lowerBand, price float64) bool {
	return sma50 > ema20 && macd > signal && rsi < 30 && price <= lowerBand
}

func ShouldBuyPut(sma50, ema20, macd, signal, rsi, upperBand, price float64) bool {
	return sma50 < ema20 && macd < signal && rsi > 70 && price >= upperBand
}
