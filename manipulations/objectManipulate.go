package methods

import (
	"fmt"
	objectTest "main/object"
)

// Now, we go try to understand how to manipulate objects
func Methods() {
	fmt.Println("=== Object Funtions Manipulations ===")
	withdraw()
	verifyUserName()
}

func withdraw() {
	account1 := objectTest.CurrencyAccount{UserName: "User", AccountNumber: 1, AgencyNumber: 2, Balance: 1.5}
	account2 := objectTest.CurrencyAccount{UserName: "User", AccountNumber: 1, AgencyNumber: 2, Balance: 1.5}

	fmt.Println("--- Withdraw with custom function ---")
	account1.Withdraw(account2.Balance)
}

func verifyUserName() {
	account1 := objectTest.CurrencyAccount{UserName: "User", AccountNumber: 1, AgencyNumber: 2, Balance: 1.5}
	account2 := objectTest.CurrencyAccount{UserName: "User", AccountNumber: 1, AgencyNumber: 2, Balance: 1.5}
	account3 := objectTest.CurrencyAccount{UserName: "User", AccountNumber: 1, AgencyNumber: 2, Balance: 1.5}

	fmt.Println("Testing custom string comparation method")
	fmt.Println("Account 1 x Account 2: ", account1.VerifyUserName(account2.UserName))
	fmt.Println("Account 1 x Account 3: ", account1.VerifyUserName(account3.UserName))
	fmt.Println("Account 2 x Account 3: ", account2.VerifyUserName(account3.UserName))
}
