package main

import "fmt"

type contactInfo struct {
  email string
  zip int
}

type person struct {
  firstName string
  lastName string
  contact contactInfo
}

func main()  {
  var alex person
  alex.print()
  alex.firstName = "Alex"
  alex.lastName = "Andersen"
  alex.print()
  alex.updateName("John")
  alex.print()
}

func (p person) print() {
  fmt.Println(p)
}

func (pointerToPerson *person) updateName(newFirstName string) {
  (*pointerToPerson).firstName = newFirstName
}