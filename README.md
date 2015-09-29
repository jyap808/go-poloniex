go-poloniex
==========

go-poloniex is an implementation of the Poloniex API (public and private) in Golang.

Based off of https://github.com/toorop/go-bittrex/

## Import
	import "github.com/jyap808/go-poloniex"
	
## Usage
~~~ go
package main

import (
	"fmt"
	"github.com/jyap808/go-poloniex"
)

const (
	API_KEY    = "YOUR_API_KEY"
	API_SECRET = "YOUR_API_SECRET"
)

func main() {
	// Poloniex client
	poloniex := poloniex.New(API_KEY, API_SECRET)

	// Get tickers
    tickers, err := poloniex.GetTickers()
	fmt.Println(err, tickers)
}
~~~	

See ["Examples" folder for more... examples](https://github.com/jyap808/go-poloniex/blob/master/examples/poloniex.go)

## Stay tuned
[Follow me on Twitter](https://twitter.com/jyap)

