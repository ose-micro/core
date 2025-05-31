package utils

import (
	"fmt"
)
// Currency type definition
type Currency string

// Define all possible currencies as constants
const (
	USD  Currency = "USD"
	EUR  Currency = "EUR"
	GBP  Currency = "GBP"
	JPY  Currency = "JPY"
	CNY  Currency = "CNY"
	INR  Currency = "INR"
	NGN  Currency = "NGN"
	SLL  Currency = "SLL"
	AUD  Currency = "AUD"
	CAD  Currency = "CAD"
	CHF  Currency = "CHF"
	ZAR  Currency = "ZAR"
	KRW  Currency = "KRW"
	BRL  Currency = "BRL"
	MXN  Currency = "MXN"
	SEK  Currency = "SEK"
	NOK  Currency = "NOK"
	DKK  Currency = "DKK"
	RUB  Currency = "RUB"
)

// currencySymbols maps ISO currency codes to their symbols
var currencySymbols = map[Currency]string{
	USD: "$", EUR: "€", GBP: "£", JPY: "¥", CNY: "¥", INR: "₹", NGN: "₦",
	SLL: "Le", AUD: "A$", CAD: "C$", CHF: "CHF", ZAR: "R", KRW: "₩",
	BRL: "R$", MXN: "$", SEK: "kr", NOK: "kr", DKK: "kr", RUB: "₽",
}

// Money formats the amount with the appropriate currency symbol
func MoneyWithCurrency(amount float64, currencyCode Currency) string {
	symbol, found := currencySymbols[currencyCode]
	if !found {
		// Fallback: use code as prefix if symbol not found
		symbol = string(currencyCode) + " "
	}

	return fmt.Sprintf("%s%s", symbol, Money(amount))
}

func Money(val float64) string {
	parts := fmt.Sprintf("%.2f", val)
	intPart, fracPart := parts[:len(parts)-3], parts[len(parts)-3:]

	var result []byte
	count := 0
	for i := len(intPart) - 1; i >= 0; i-- {
		if count == 3 {
			result = append([]byte{','}, result...)
			count = 0
		}
		result = append([]byte{intPart[i]}, result...)
		count++
	}
	return string(result) + fracPart
}
