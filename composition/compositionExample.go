package composition

import "fmt"

//In go, like other languages we have OOP too.
//We dont have inheritance, but we have composition.
type User struct {
	Id        int
	FederalId int
	UserName  string
	Address   Address
}

type Address struct {
	Id           int
	city         string
	number       int
	street       string
	neighborhood string
}

func CompositionExample() {
	initComposition()
}

func initComposition() {
	user := User{
		Id:        0,
		FederalId: 000000000000,
		UserName:  "Franky",
		Address: Address{
			Id:           0,
			city:         "Water 7",
			number:       0,
			street:       "street",
			neighborhood: "ocean",
		},
	}

	fmt.Println(user)
}
