package main

import "fmt"

type contactInfo struct {
	email string
	zipCode int
}

type person struct {
	firstName string
	lastName string
	contact contactInfo
}

func main() {
	// viet := {"Viet", "Nguyen"}
	// viet := person{firstName: "Viet", lastName: "Nguyen"}
	// var viet person
	// viet = {"Viet", "Nguyen"}
	// viet.firstName = "Viet"
	// viet.lastName = "Nguyen"
	
	viet := person{
		firstName: "viet",
		lastName: "nguyen",
		contact: contactInfo{
			email: "viet@gmail.ca",
			zipCode: 10000,
		},
	}
	vietPointer := &viet

	// fmt.Println(vietPointer)
	vietPointer.updateName("nam")
	viet.print()
	// fmt.Println(viet)
	
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}