package objectComparison

import (
	"fmt"
	objectTest "main/object"
)

func TypesTest() {
	fmt.Println("=== Testing Types ===")
	onConstructor()
	onFields()
}

func onConstructor() {

	fmt.Println("--- Testing by constructor ---")

	account1 := objectTest.CurrencyAccount{UserName: "User", AccountNumber: 1, AgencyNumber: 2, Balance: 1.5}
	account2 := objectTest.CurrencyAccount{UserName: "User", AccountNumber: 1, AgencyNumber: 2, Balance: 1.5}

	fmt.Println("Testing match: ", account1 == account2)
	//False because de adress memory is diferent
	fmt.Println("Testing match address match: ", &account1 == &account2)

}

// Because we allocate the memory for the objects, the thinks work a little diferent.
func onFields() {

	fmt.Println("--- Testing allocate on memory ---")

	var account1 = new(objectTest.CurrencyAccount)
	account1.UserName = "Igor"
	account1.AgencyNumber = 2
	account1.AccountNumber = 1
	account1.Balance = 100.2

	var account2 = new(objectTest.CurrencyAccount)
	account2.UserName = "Igor"
	account2.AgencyNumber = 2
	account2.AccountNumber = 1
	account2.Balance = 100.2

	fmt.Println("Object one simple desmontration: ", account1)
	fmt.Println("Object two simple desmontration: ", account2)

	fmt.Println("Object one on memory: ", &account1)
	fmt.Println("Object two on memory: ", &account2)

	fmt.Println("Object one pointer: ", *account1)
	fmt.Println("Object two pointer: ", *account2)

	fmt.Println("Testing match", account1 == account2)
	fmt.Println("Testing match acessing on memory: ", &account1 == &account2)
	fmt.Println("Testing match by pointers: ", *account1 == *account2)
}
