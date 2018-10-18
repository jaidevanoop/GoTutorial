package main

import "fmt"

type contactInfo struct {
	email   string
	zipCope int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	alex := person{firstName: "Alex", lastName: "Anderson"}
	fmt.Printf("%v\n", alex)

	var bob person
	bob.firstName = "Bob"
	bob.lastName = "Anderson"

	fmt.Printf("%+v\n", bob)

	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipCope: 94000,
		},
	}

	jim.print()
	jimPointer := &jim
	jimPointer.updateName("Jimmy")
	jim.print()
	jim.updateName("James")
	jim.print()
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}
func (p person) print() {
	fmt.Printf("%+v\n", p)
}
