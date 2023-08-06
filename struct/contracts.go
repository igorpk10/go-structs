package structs

type Teenager struct {
	Age int
}

func (c *Teenager) AddAge() {
	c.Age += 1
}

type Child struct {
	Age int
}

func (c *Child) AddAge() {
	c.Age += 1
}

// In go, we can define implict contracts
// Its a little diferent when i can think about it in other languagens.
// We need to have the function on the struct and define a interface for him
// And we can just call the interface method sending the struct memory position
type AddPersonAge interface {
	AddAge()
}

func ContractTest() {
	person1 := Teenager{Age: 15}
	person2 := Child{Age: 8}

	AddPersonAge(&person1)
	AddPersonAge(&person2)
}
