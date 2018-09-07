package mop

import (
	"fmt"
	"testing"

	od "github.com/barchart/barchart-ondemand-client-golang"
)

func Test(t *testing.T) {
	od := od.New("FREE_API_KEY", false)
	od.BaseURL = "https://marketdata.websol.barchart.com/"
	m := NewMarket(od)

	fmt.Println("MM", m.Fetch())
}
