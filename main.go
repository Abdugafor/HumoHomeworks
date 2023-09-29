package main

import (
	"GoHome/CardService"
	"fmt"
)

func main() {
	var cardNum string
	fmt.Println("Write your card number")
	fmt.Scan(&cardNum)

	CardService.AddCard(cardNum)
}
