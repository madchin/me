package main

import (
	"fmt"
	"strings"
)

var (
	ADDITIVE_NOTATION = map[string]int{
		"M": 1000,
		"D": 500,
		"C": 100,
		"L": 50,
		"X": 10,
		"V": 5,
		"I": 1,
	}
	SUBTRACTIVE_NOTATION = map[string]int{
		"IV": 4,
		"IX": 9,
		"XL": 40,
		"XC": 90,
		"CD": 400,
		"CM": 900,
	}
)

func Decode(roman string) (result int) {
	for key, value := range SUBTRACTIVE_NOTATION {
		for strings.Contains(roman, key) {
			roman = strings.Replace(roman, key, "", 1)
			result += value
		}
	}

	for key, value := range ADDITIVE_NOTATION {
		for strings.Contains(roman, key) {
			roman = strings.Replace(roman, key, "", 1)
			result += value
		}
	}
	return
}

func main() {
	aa := ADDITIVE_NOTATION["elo"]
	fmt.Print(aa)
}

/*

Create a function that takes a Roman numeral as its argument and returns its value as a numeric decimal integer.
You don't need to validate the form of the Roman numeral.

Modern Roman numerals are written by expressing each decimal digit of the number to be encoded separately,
starting with the leftmost digit and skipping any 0s. So 1990 is rendered "MCMXC" (1000 = M, 900 = CM, 90 = XC)
and 2008 is rendered "MMVIII" (2000 = MM, 8 = VIII). The Roman numeral for 1666, "MDCLXVI",
uses each letter in descending order.

*/
