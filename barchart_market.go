package mop

import (
	"fmt"
	"strconv"

	od "github.com/barchart/barchart-ondemand-client-golang"
)

func NewMarket(onDemand *od.OnDemand) *Market {

	market := &Market{}

	market.IsClosed = false
	market.Dow = make(map[string]string)
	market.Nasdaq = make(map[string]string)
	market.Sp500 = make(map[string]string)

	market.Tokyo = make(map[string]string)
	market.HongKong = make(map[string]string)
	market.London = make(map[string]string)
	market.Frankfurt = make(map[string]string)

	market.Yield = make(map[string]string)
	market.Oil = make(map[string]string)
	market.Yen = make(map[string]string)
	market.Euro = make(map[string]string)
	market.Gold = make(map[string]string)

	market.OnDemand = onDemand

	market.errors = ``

	return market
}

func ToString(value float64, per int) string {
	return strconv.FormatFloat(value, 'f', per, 64)
}

func IntToString(value int) string {
	return strconv.Itoa(value)
}

// Fetch uses the Barchart OnDemand go library to fetch quote data
func (market *Market) Fetch() (self *Market) {

	quotes, _ := market.OnDemand.Quote([]string{"YM*0", "NQ*0", "ES*0", "NY*0",
		"HS*0", "UV*0", "DY*0", "$TNX", "CL*0", "J6*0", "E6*0", "GC*0"}, []string{"bid", "ask"})

	for _, q := range quotes.Results {
		switch sym := q.Symbol[:2]; sym {
		case "YM":
			market.Dow[`change`] = ToString(q.NetChange, 2)
			market.Dow[`latest`] = ToString(q.LastPrice, 2)
			market.Dow[`percent`] = ToString(q.PercentChange, 2)
		case "NQ":
			market.Nasdaq[`change`] = ToString(q.NetChange, 2)
			market.Nasdaq[`latest`] = ToString(q.LastPrice, 2)
			market.Nasdaq[`percent`] = ToString(q.PercentChange, 2)
		case "ES":
			market.Sp500[`change`] = ToString(q.NetChange, 2)
			market.Sp500[`latest`] = ToString(q.LastPrice, 2)
			market.Sp500[`percent`] = ToString(q.PercentChange, 2)
		case "NY":
			market.Tokyo[`change`] = ToString(q.NetChange, 2)
			market.Tokyo[`latest`] = ToString(q.LastPrice, 2)
			market.Tokyo[`percent`] = ToString(q.PercentChange, 2)
		case "HS":
			market.HongKong[`change`] = ToString(q.NetChange, 2)
			market.HongKong[`latest`] = ToString(q.LastPrice, 2)
			market.HongKong[`percent`] = ToString(q.PercentChange, 2)
		case "UV":
			market.London[`change`] = ToString(q.NetChange, 2)
			market.London[`latest`] = ToString(q.LastPrice, 2)
			market.London[`percent`] = ToString(q.PercentChange, 2)
		case "DY":
			market.Frankfurt[`change`] = ToString(q.NetChange, 2)
			market.Frankfurt[`latest`] = ToString(q.LastPrice, 2)
			market.Frankfurt[`percent`] = ToString(q.PercentChange, 2)
		case "$T":
			market.Yield[`change`] = ToString(q.NetChange, 2)
			market.Yield[`latest`] = ToString(q.LastPrice, 2)
			market.Yield[`percent`] = ToString(q.PercentChange, 2)
		case "CL":
			market.Oil[`change`] = ToString(q.NetChange, 2)
			market.Oil[`latest`] = ToString(q.LastPrice, 2)
			market.Oil[`percent`] = ToString(q.PercentChange, 2)
		case "J6":
			market.Yen[`change`] = ToString(q.NetChange, 2)
			market.Yen[`latest`] = ToString(q.LastPrice, 2)
			market.Yen[`percent`] = ToString(q.PercentChange, 2)
		case "E6":
			market.Euro[`change`] = ToString(q.NetChange, 2)
			market.Euro[`latest`] = ToString(q.LastPrice, 2)
			market.Euro[`percent`] = ToString(q.PercentChange, 2)
		case "GC":
			market.Gold[`change`] = ToString(q.NetChange, 2)
			market.Gold[`latest`] = ToString(q.LastPrice, 2)
			market.Gold[`percent`] = ToString(q.PercentChange, 2)
		default:
			fmt.Println("Sym", sym)
		}
	}

	return market
}

func (market *Market) Ok() (bool, string) {
	return true, market.errors
}

func (market *Market) isMarketOpen(body []byte) []byte {
	return []byte{}
}
