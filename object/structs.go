package objectTest

import "fmt"

//In go our class work like in old C language.
//So, we have structs
type CurrencyAccount struct {
	UserName      string
	AgencyNumber  int
	AccountNumber int
	Balance       float32
}

//We can create functions to our structs
//So, if in some place we need to call her
//Just call currencyAccount.withdraw()
func (c *CurrencyAccount) Withdraw(number float32) float32 {
	c.Balance = c.Balance - number
	return c.Balance
}

//We can create functions what not change properties, like:
//This function can't atribute other values to the respective object
//Because she dont have the pointer of memory.
func (u CurrencyAccount) VerifyUserName(userName string) bool {
	return u.UserName == userName
}

//We have some types for init the structs
func InitWithoutExplictFields() {
	account := CurrencyAccount{"user", 5, 1, 1.5}
	fmt.Println("Implicit constructor: ", account)
}

func InitWithExplicitFields() {
	account := CurrencyAccount{UserName: "user", AgencyNumber: 5, AccountNumber: 1, Balance: 1.5}
	fmt.Println("Explict constructor: ", account)
}

//I cant think a name for it
func InitAlongTheNecessary() {
	var account *CurrencyAccount
	account = new(CurrencyAccount)
	account.AccountNumber = 5
	account.AgencyNumber = 5
	account.UserName = "User"
	account.Balance = 200.00

	// To acess with this type you need to acess with pointer
	fmt.Println("Set itens account:", *account)
}
