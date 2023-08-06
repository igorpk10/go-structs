package objectComparison

import (
	"fmt"
	objectTest "main/object"
)

// On go, we have access on a memory, so we need to be careful about it.
func OnConstructor() {
	account1 := objectTest.CurrencyAccount{UserName: "User", AccountNumber: 1, AgencyNumber: 2, Balance: 1.5}
	account2 := objectTest.CurrencyAccount{UserName: "User", AccountNumber: 1, AgencyNumber: 2, Balance: 1.5}

	fmt.Println("Testing match: ", account1 == account2)

	//False because de adress memory is diferent
	fmt.Println("Testing match address match: ", &account1 == &account2)
}

// Because we allocate the memory for the objects, the thinks work a little diferent.
func OnFields() {

	fmt.Println("--- Testing allocate on memory ---")

	var account1 = new(objectTest.CurrencyAccount)
	account1.UserName = "User"
	account1.AgencyNumber = 2
	account1.AccountNumber = 1
	account1.Balance = 100.2

	var account2 = new(objectTest.CurrencyAccount)
	account2.UserName = "User"
	account2.AgencyNumber = 2
	account2.AccountNumber = 1
	account2.Balance = 100.2

	fmt.Println("Object one on memory: ", &account1)
	fmt.Println("Object two on memory: ", &account2)

	fmt.Println("Object one pointer: ", *account1)
	fmt.Println("Object two pointer: ", *account2)

	//False, because these try to match memory address too
	fmt.Println("Testing match", account1 == account2)
	//False, because these try to match only the address
	fmt.Println("Testing match acessing on memory: ", &account1 == &account2)
	//True, because these try to match just the content
	fmt.Println("Testing match by pointers: ", *account1 == *account2)
}
