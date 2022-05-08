package main

import (
	"fmt"
	"reflect"
)

type person struct {
	name string
	age  int
}

func newPerson(name string) *person {
	p := person{name: name}
	p.age = 42
	return &p
}

func main() {
	fmt.Println(person{"bob", 20})

	fmt.Println(person{name: "Alice", age: 30})

	fmt.Println(person{name: "fred"})

	fmt.Println(&person{name: "Ann", age: 40})

	fmt.Println(newPerson("Jon"))

	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)

	sp := &s
	fmt.Println(sp.age)

	fmt.Println(reflect.TypeOf(sp))
}
