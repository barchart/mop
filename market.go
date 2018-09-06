// Copyright (c) 2013-2016 by Michael Dvorkin. All Rights Reserved.
// Use of this source code is governed by a MIT-style license that can
// be found in the LICENSE file.

package mop

import (
	"regexp"

	od "github.com/barchart/barchart-ondemand-client-golang"
)

// Market stores current market information displayed in the top three lines of
// the screen. The market data is fetched and parsed from the HTML page above.
type Market struct {
	IsClosed  bool              // True when U.S. markets are closed.
	Dow       map[string]string // Hash of Dow Jones indicators.
	Nasdaq    map[string]string // Hash of NASDAQ indicators.
	Sp500     map[string]string // Hash of S&P 500 indicators.
	Tokyo     map[string]string
	HongKong  map[string]string
	London    map[string]string
	Frankfurt map[string]string
	Yield     map[string]string
	Oil       map[string]string
	Yen       map[string]string
	Euro      map[string]string
	Gold      map[string]string
	regex     *regexp.Regexp // Regex to parse market data from HTML.
	errors    string         // Error(s), if any.
	OnDemand  *od.OnDemand
}
