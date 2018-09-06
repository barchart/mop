// Copyright (c) 2013-2016 by Michael Dvorkin. All Rights Reserved.
// Use of this source code is governed by a MIT-style license that can
// be found in the LICENSE file.

package mop

// Stock stores quote information for the particular stock ticker. The data
// for all the fields except 'Advancing' is fetched using Yahoo market API.
type Stock struct {
	Ticker     string // Stock ticker.
	LastTrade  string // l1: last trade.
	Change     string // c6: change real time.
	ChangePct  string // k2: percent change real time.
	Open       string // o: market open price.
	Low        string // g: day's low.
	High       string // h: day's high.
	Low52      string // j: 52-weeks low.
	High52     string // k: 52-weeks high.
	Volume     string // v: volume.
	AvgVolume  string // a2: average volume.
	PeRatio    string // r2: P/E ration real time.
	PeRatioX   string // r: P/E ration (fallback when real time is N/A).
	Dividend   string // d: dividend.
	Yield      string // y: dividend yield.
	MarketCap  string // j3: market cap real time.
	MarketCapX string // j1: market cap (fallback when real time is N/A).
	Advancing  bool   // True when change is >= $0.
}

// Quotes stores relevant pointers as well as the array of stock quotes for
// the tickers we are tracking.
type Quotes struct {
	market  *Market  // Pointer to Market.
	profile *Profile // Pointer to Profile.
	stocks  []Stock  // Array of stock quote data.
	errors  string   // Error string if any.
}
