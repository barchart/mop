package mop

import (
	"fmt"
	"testing"

	od "github.com/barchart/barchart-ondemand-client-golang"
)

func Test(t *testing.T) {
	m := NewMarket(od.New("API_KEY", false))

	fmt.Println("MM", m.Fetch())
}
