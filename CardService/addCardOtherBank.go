package CardService

var otherBanksList = []string{"1234-0222", "1234-5654", "2333-2222", "1222-1233", "3212-3211"}

func addCardOtherBank(cardNum string) []string {
	otherBanksList = append(otherBanksList, cardNum)
	return otherBanksList
}
