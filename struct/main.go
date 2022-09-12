package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	contactInfo
}

type contactInfo struct {
	email   string
	zipCode int
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func main() {
	alex := person{
		firstName: "Alex",
		lastName:  "Alexandar",
		contactInfo: contactInfo{
			email:   "alex@gmail.com",
			zipCode: 9800,
		},
	}
	alex.print()
	name := "bill"
	fmt.Println(*&name)
}
