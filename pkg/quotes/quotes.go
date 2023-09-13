// https://github.com/Atheer2104/MotivationalQuotesGolang/blob/master/main.go

package quotes

import (
	_ "embed"
	"encoding/json"
	"math/rand"
)

//go:embed quotes.json
var quotesBytes []byte

var quotes Quotes

type Quotes struct {
	Quotes []Quote
}

type Quote struct {
	QuoteText   string
	QuoteAuthor string
}

func init() {
	err := json.Unmarshal(quotesBytes, &quotes)
	if err != nil {
		panic(err.Error())
	}
}

func GetQuote() Quote {
	return quotes.Quotes[rand.Int31n(int32(len(quotes.Quotes)))]
}
