package mop

func NewQuotes(market *Market, profile *Profile) *Quotes {
	return &Quotes{
		market:  market,
		profile: profile,
		errors:  ``,
	}
}

func (quotes *Quotes) Fetch() (self *Quotes) {

	qs, _ := quotes.market.OnDemand.Quote(quotes.profile.Tickers,
		[]string{"bid", "fiftyTwoWkHigh", "dividendRateAnnual", "dividendYieldAnnual", "fiftyTwoWkLow", "avgVolume", "ask"})

	quotes.stocks = make([]Stock, len(qs.Results))

	for i, q := range qs.Results {
		quotes.stocks[i].Ticker = q.Symbol
		quotes.stocks[i].LastTrade = ToString(q.LastPrice, 2)
		quotes.stocks[i].Change = ToString(q.NetChange, 2)
		quotes.stocks[i].ChangePct = ToString(q.PercentChange, 2)
		quotes.stocks[i].Open = ToString(q.Open, 2)
		quotes.stocks[i].High = ToString(q.High, 2)
		quotes.stocks[i].Low = ToString(q.Low, 2)
		quotes.stocks[i].Low52 = ToString(q.FiftyTwoWkLow, 2)
		quotes.stocks[i].High52 = ToString(q.FiftyTwoWkHigh, 2)
		quotes.stocks[i].Volume = IntToString(q.Volume)
		quotes.stocks[i].AvgVolume = IntToString(q.AvgVolume)
		quotes.stocks[i].PeRatio = ToString(q.LastPrice, 2)
		quotes.stocks[i].PeRatioX = ToString(q.LastPrice, 2)
		quotes.stocks[i].Dividend = NAString(q.DividendRateAnnual)
		quotes.stocks[i].Yield = NAString(q.DividendYieldAnnual)
		quotes.stocks[i].MarketCap = ToString(q.LastPrice, 2)
		quotes.stocks[i].MarketCapX = ToString(q.LastPrice, 2)
		quotes.stocks[i].Advancing = (q.NetChange > 0)

		//fmt.Println(q.DividendRateAnnual)
	}
	//fmt.Println("symbols", qs)

	return quotes
}

func NAString(field string) string {
	if len(field) == 0 {
		return "N/A"
	}
	return field

}

func (quotes *Quotes) Ok() (bool, string) {
	return quotes.errors == ``, quotes.errors
}

func (quotes *Quotes) AddTickers(tickers []string) (added int, err error) {
	if added, err = quotes.profile.AddTickers(tickers); err == nil && added > 0 {
		quotes.stocks = nil // Force fetch.
	}
	return
}

func (quotes *Quotes) RemoveTickers(tickers []string) (removed int, err error) {
	if removed, err = quotes.profile.RemoveTickers(tickers); err == nil && removed > 0 {
		quotes.stocks = nil // Force fetch.
	}
	return
}

func (quotes *Quotes) isReady() bool {
	return (quotes.stocks == nil || !quotes.market.IsClosed) && len(quotes.profile.Tickers) > 0
}
