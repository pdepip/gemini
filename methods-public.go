package gemini

import (
	"encoding/json"
	"strconv"
)

// Symbols
func (g *GeminiAPI) Symbols() ([]string, error) {

	url := g.url + SYMBOLS_URL

	body, err := request("GET", url, nil, nil)
	if err != nil {
		return nil, err
	}

	var res []string
	json.Unmarshal(body, &res)

	return res, nil
}

// Order Book
func (g *GeminiAPI) OrderBook(symbol string, limitBids int, limitAsks int) (Book, error) {

	url := g.url + BOOK_URL + "/" + symbol

	params := map[string]string{
		"limit_bids": strconv.Itoa(limitBids),
		"limit_asks": strconv.Itoa(limitAsks),
	}

	var book Book

	body, err := request("GET", url, nil, params)
	if err != nil {
		return book, err
	}

	json.Unmarshal(body, &book)

	return book, nil
}

// Trades
func (g *GeminiAPI) Trades(symbol string, since int64, limitTrades int, includeBreaks bool) ([]Trade, error) {

	url := g.url + TRADES_URL + "/" + symbol

	params := map[string]string{
		// "since":          strconv.Itoa(int(since)),
		"limit_trades":   strconv.Itoa(limitTrades),
		"include_breaks": strconv.FormatBool(includeBreaks),
	}

	var res []Trade

	body, err := request("GET", url, nil, params)
	if err != nil {
		return nil, err
	}

	json.Unmarshal(body, &res)

	return res, nil
}