package tickr

import (
	"github.com/rafatbiin/tickr/data"
	"regexp"
	"strings"
	"unicode"
)

// defaultGramSize is the highest number of n from ngram.
const defaultGramSize = 5

// Ticker serves two purpose:
// 1. match input string with the Company Trie
// 2. match the company ticker with the sector map
type Ticker struct {
	company *Company
	sectors map[string]string
}

func New() (*Ticker, error) {
	bdstock, err := data.LoadBDStock()
	if err != nil {
		return nil, err
	}

	trie := buildCompany(bdstock)
	sectors := buildSectors(bdstock)
	return &Ticker{company: trie, sectors: sectors}, nil
}

func buildCompany(bdstock *data.BDStock) *Company {
	company := NewCompany()
	for _, c := range bdstock.Data.Companies {
		for _, name := range c.Names {
			company.Put(sanitize(name, true, true), c.Ticker)
		}
	}
	return company
}

func buildSectors(bdstock *data.BDStock) map[string]string {
	sectors := make(map[string]string)
	for _, c := range bdstock.Data.Companies {
		sectors[c.Ticker] = c.Sector
	}
	return sectors
}

// Get will return a map of all the matched company tickers with their respective frequency
func (t *Ticker) Get(text string) map[string]int {
	tickers := make(map[string]int)

	text = strings.TrimSpace(text)
	text = sanitize(text, true, false)

	re, err := regexp.Compile(`\s+`)
	if err != nil {
		return nil
	}
	tokens := re.Split(text, -1)

	gramSize := min(defaultGramSize, len(tokens))
	i := 0

	for ; i < len(tokens)-gramSize+1; i++ {
		if ticker := t.find(tokens[i : i+gramSize]); ticker != "" {
			tickers[ticker]++
		}
	}
	for ; i < len(tokens); i++ {
		if ticker := t.find(tokens[i:]); ticker != "" {
			tickers[ticker]++
		}
	}
	return tickers
}

// find will try to match the candidate prefix with the Company Trie
func (t *Ticker) find(tokens []string) string {
	candidate := strings.Join(tokens, "")
	return t.company.Get(candidate)
}

func (t *Ticker) Sector(ticker string) string {
	return t.sectors[ticker]
}

// sanitize will remove all the spaces and punctuations(if set true) from the input string str.
func sanitize(str string, removePunct, removeSpace bool) string {
	var b strings.Builder
	b.Grow(len(str))
	for _, ch := range str {
		if removePunct && unicode.IsPunct(ch) {
			continue
		}
		if removeSpace && unicode.IsSpace(ch) {
			continue
		}
		b.WriteRune(unicode.ToLower(ch))
	}
	return b.String()
}

func min(left, right int) int {
	if left > right {
		return right
	}

	return left
}
