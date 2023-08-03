package instance

import (
	"fmt"
	objectTest "main/object"
)

func Hello() {
	initType1()
	initType2()
	initType3()
}

// For me is not a good idea, because its a error margin
func initType1() {
	account := objectTest.CurrencyAccount{"user", 5, 1, 1.5}
	fmt.Println("Implicit constructor: ", account)
}

func initType2() {
	account := objectTest.CurrencyAccount{UserName: "user", AgencyNumber: 5, AccountNumber: 1, Balance: 1.5}
	fmt.Println("Explict constructor: ", account)
}

func initType3() {
	var account *objectTest.CurrencyAccount
	account = new(objectTest.CurrencyAccount)
	account.AccountNumber = 5
	account.AgencyNumber = 5
	account.UserName = "User"
	account.Balance = 200.00

	// To acess with this type you need to acess with pointer
	// *account
	fmt.Println("Set itens account:", *account)
}
