package main

import "fmt"

type contactInfo struct {
	email   string
	zipCope int
}

type person struct {
	firstName string
	lastName  string
	cInfo     contactInfo
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
		cInfo: contactInfo{
			email:   "jim@gmail.com",
			zipCope: 94000,
		},
	}

	fmt.Printf("%+v\n", jim)
}
