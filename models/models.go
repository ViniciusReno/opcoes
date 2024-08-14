package models

type StockData struct {
	Chart Chart `json:"chart"`
}

type Chart struct {
	Result []Result `json:"result"`
}

type Result struct {
	Indicators Indicators `json:"indicators"`
	Meta       Meta       `json:"meta"`
	Timestamp  []int64    `json:"timestamp"`
}

type Indicators struct {
	Quote []Quote `json:"quote"`
}

type Quote struct {
	Close  []float64 `json:"close"`
	Volume []int64   `json:"volume"`
}

type Meta struct {
	Symbol             string  `json:"symbol"`
	RegularMarketPrice float64 `json:"regularMarketPrice"`
}
