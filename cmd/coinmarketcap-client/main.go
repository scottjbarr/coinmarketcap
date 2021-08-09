package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/scottjbarr/coinmarketcap"
)

func main() {
	c, err := coinmarketcap.New(os.Getenv("API_KEY"))
	if err != nil {
		panic(err)
	}

	symbols := ""
	convert := ""

	flag.StringVar(&symbols, "symbols", "", "eg. BSV,BTC,ETH")
	flag.StringVar(&convert, "convert", "USD", "eg. USD")
	flag.Parse()

	if len(symbols) == 0 {
		flag.Usage()
		os.Exit(1)
	}

	q := coinmarketcap.QuotesQuery{
		Symbol:  strings.Split(symbols, ","),
		Convert: strings.Split(convert, ","),
	}

	res, err := c.QuotesLatest(q)
	if err != nil {
		panic(err)
	}

	for symbol, q := range res.Data {
		for c, detail := range q.Quote {
			fmt.Printf("%s_%s = %+v\n", symbol, c, detail)
		}
	}
}
