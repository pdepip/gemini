package gemini

import (
	"encoding/json"
	"strconv"
)

// Symbols
func (api *Api) Symbols() ([]string, error) {

	url := api.url + SYMBOLS_URI

	var symbols []string

	body, err := api.request("GET", url, nil)
	if err != nil {
		return nil, err
	}

	json.Unmarshal(body, &symbols)

	return symbols, nil
}

// Ticker
func (api *Api) Ticker(symbol string) (Ticker, error) {

	url := api.url + TICKER_URI + symbol

	var ticker Ticker

	body, err := api.request("GET", url, nil)
	if err != nil {
		return ticker, err
	}

	json.Unmarshal(body, &ticker)

	return ticker, nil
}

// Order Book
func (api *Api) OrderBook(symbol string, limitBids, limitAsks int) (Book, error) {

	url := api.url + BOOK_URI + symbol
	params := map[string]interface{}{
		"limit_bids": strconv.Itoa(limitBids),
		"limit_asks": strconv.Itoa(limitAsks),
	}

	var book Book

	body, err := api.request("GET", url, params)
	if err != nil {
		return book, err
	}

	json.Unmarshal(body, &book)

	return book, nil
}

// Trades
func (api *Api) Trades(symbol string, since int64, limitTrades int, includeBreaks bool) ([]Trade, error) {

	url := api.url + TRADES_URI + symbol
	params := map[string]interface{}{
		"limit_trades":   strconv.Itoa(limitTrades),
		"include_breaks": strconv.FormatBool(includeBreaks),
	}

    if since >= 0 {
        params["since"] = strconv.Itoa(int(since))
    }

	var res []Trade

	body, err := api.request("GET", url, params)
	if err != nil {
		return nil, err
	}

	json.Unmarshal(body, &res)

	return res, nil
}

// Current Auction
func (api *Api) CurrentAuction(symbol string) (CurrentAuction, error) {

	url := api.url + AUCTION_URI + symbol

	var auction CurrentAuction

	body, err := api.request("GET", url, nil)
	if err != nil {
		return auction, err
	}

	json.Unmarshal(body, &auction)

	return auction, nil
}

// Auction History
func (api *Api) AuctionHistory(symbol string, since int64, limit int, includeIndicative bool) ([]Auction, error) {

	url := api.url + AUCTION_URI + symbol + "/history"
	params := map[string]interface{}{
		"since":                 strconv.Itoa(int(since)),
		"limit_auction_results": strconv.Itoa(limit),
		"include_indicative":    strconv.FormatBool(includeIndicative),
	}

	var auctions []Auction

	body, err := api.request("GET", url, params)
	if err != nil {
		return auctions, err
	}

	json.Unmarshal(body, &auctions)

	return auctions, nil
}
