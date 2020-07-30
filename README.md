# Tickr ![CI](https://github.com/rafatbiin/tickr/workflows/CI/badge.svg)
News CateGOrizer for public companies listed in stock exchanges in Bangladesh.


## Overview
Tickr provides a Go package which let you categorize
news data(title,article etc.) into TICKER of publicly traded
companies in stock exchanges here in Bangladesh.


It is based on the [bdstock.json](https://github.com/rafatbiin/tickr/blob/master/data/bdstock.json) dataset that I developed
which includes 3 datapoints for each public company listed in stock market of Bangladesh:
  1. Company **names**, the way they usually appear in newspaper.
  2. **TICKER** of that company.
  3. **Sector** of that company.
  
This package includes two API:
1. **Get()**: for extracting TICKER(s) from given news data.
2. **Sector()**: for extracting sector from given TICKER.

In *addition* this library provides a [CLI](#cli) to perform the above operations via console.

## Getting Started

### Installing

To start using Tickr, install Go and run `go get`:

```sh
$ go get github.com/rafatbiin/tickr/...
```

This will retrieve the library and install the `tickr` command line utility into
your `$GOBIN` path.

### Getting the TICKER(s)

The top-level object in Tickr is a `Ticker`.

To get TICKER(s) for your news data, simply create an instance of `Ticker` and use the `t.Get(news_data)` function:

```go
package main

import (
	"github.com/rafatbiin/tickr"
)

func main() {
	t, err := tickr.New()
	if err != nil {
		// Handle error
	}
	news_data := "রূপালী ব্যাংকের দ্বিতীয় প্রান্তিক প্রকাশ"
	tickers := t.Get(news_data) // map(TICKER->frequency) => {"RUPALIBANK":1}
	...
}
```

### Getting the Sector

To get Sector of a publicly traded Compnay in Bangladesh, use the `t.Sector(TICKER)` function:

```go
package main

import (
	"github.com/rafatbiin/tickr"
)

func main() {
	t, err := tickr.New()
	if err != nil {
		// Handle error
	}
	ticker := "RUPALIBANK"
	s := t.Sector(ticker) // bank
	...
}
```


## CLI
For TICKER(s):
```sh
$ tickr ticker "রূপালী ব্যাংকের দ্বিতীয় প্রান্তিক প্রকাশ"
{"RUPALIBANK":1}
```

For Sector:
```sh
$ tickr sector RUPALIBANK
bank
```