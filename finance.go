package main

import (
	"fmt"
	"net/http"
	"time"
)

type Symbol string

type Candle struct {
	Timestamp int64
	Volume    float64
	Open      float64
	Close     float64
	High      float64
	Low       float64
}

type Result struct {
	Symbol       Symbol
	Currency     string
	Timezone     string
	CurrentPrice float64
	HighestPrice float64
	LowerPrice   float64
	Candles      []Candle
}

type Query struct {
	Symbol   Symbol
	Start    int64
	End      int64
	Interval string
}

type Client interface {
	SetHttpClient(http http.Client)
	Get(symbol Symbol) (Result, error)
}

type Yahoofinance struct {
	httpClient http.Client
}

func (y *Yahoofinance) SetHttpClient(httpClient http.Client) {
	y.httpClient = httpClient
}

func (y *Yahoofinance) Get(symbol Symbol) (Result, error) {
	fmt.Println(symbol)
	return Result{}, nil
}

type Finance struct {
	HttpClient http.Client
	Client     Client
}

func (f *Finance) SetupWithYahooFinance() {
	f.Client = &Yahoofinance{}
	f.Client.SetHttpClient(f.HttpClient)
}

func main() {
	httpClient := http.Client{Timeout: 30 * time.Second}
	finance := Finance{HttpClient: httpClient}
	finance.SetupWithYahooFinance()
	finance.Client.Get("PETR4")
}
