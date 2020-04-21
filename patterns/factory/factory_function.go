package main

import "fmt"

type person struct {
	name     string
	age      int
	eyeCount int
}

type tiredPerson struct {
	name     string
	age      int
	eyeCount int
}

type Person interface {
	SayHello()
}

func (p *person) SayHello() {
	fmt.Printf("Hi, my name is %s, and I'm %d years old.\n", p.name, p.age)
}

func (p *tiredPerson) SayHello() {
	fmt.Printf("Sorry, I'm too tired to answer.\n")
}

// factory function
func NewPerson(name string, age int) Person {
	if age > 100 {
		return &tiredPerson{name, age, 2}
	}
	return &person{name, age, 2}
}

func main() {
	// use a constructor
	p := NewPerson("Jane", 21)
	// Can`t be done: p.Age = 30
	p.SayHello()

	p2 := NewPerson("Jane", 120)
	// Can`t be done: p.Age = 30
	p2.SayHello()
}
