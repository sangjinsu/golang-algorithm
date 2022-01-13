package main

import "fmt"

type Person struct {
	name string
	age  int
}

func NewPerson(name string, age int) *Person {
	return &Person{name: name, age: age}
}

func (p *Person) printName() {
	fmt.Println("name: ", p.name)
}

func (p *Person) setName(name string) {
	p.name = name
}

func (p *Person) printAge() {
	fmt.Println("name: ", p.name)
}

func (p *Person) setAge(age int) {
	p.age = age
}

type Singer01 struct {
	Person
}

type Singer02 struct {
	p Person
}

func main() {
	s01 := Singer01{Person{
		name: "jinsu",
		age:  27,
	}}

	// 승격
	s01.printAge()
	s01.Person.printAge()

	s02 := Singer02{p: Person{
		name: "jinsu",
		age:  27,
	}}

	// 승격 X
	s02.p.printAge()
}
