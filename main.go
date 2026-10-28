package main

import (
	"fmt"

	"github.com/Burlington1308/stringmod/quotes"
	"github.com/Burlington1308/stringmod/strings"
)

func main() {
	o, e := strings.CountOddEven("12345")
	fmt.Println(o, e) // 3 2

	fmt.Println(quotes.GetEmoji("turtle"))
}