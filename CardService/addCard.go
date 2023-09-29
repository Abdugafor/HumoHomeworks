package CardService

import (
	"fmt"
	"regexp"
)

var humoList = []string{"5234-0222", "6234-5654", "7333-2222", "5222-1233", "7212-3211"}

const cardRegex = `^\d{4}-\d{4}$`

func AddCard(cardNum string) {
	cardsNumberMap := make(map[string]bool)

	for _, v := range humoList {
		cardsNumberMap[v] = false
	}

	match, err := regexp.MatchString(cardRegex, cardNum)

	if err != nil {
		fmt.Printf("Regex error: %v", err.Error())
		return
	}

	if !match {
		fmt.Println("Input does not match requirements!")
	}

	_, exist := cardsNumberMap[cardNum]
	if exist {
		fmt.Println("number is already exists!")
		return
	}
	if cardNum[0] >= 53 && match {
		fmt.Println("Humo card")

		if match {
			cardsNumberMap[cardNum] = false
			fmt.Println("New list:", cardsNumberMap)
		}

	} else {
		fmt.Println("Other bank")
		otherBanksList := addCardOtherBank(cardNum)
		fmt.Println(otherBanksList)
	}

}
