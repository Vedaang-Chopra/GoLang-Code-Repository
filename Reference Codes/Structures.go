package main

import "fmt"

// Structures is a data structure that allows to compose together values of different types.
// Syntax for creating structures: -
// Aggregating different Types
type person struct{
	first_name string
	last_name string
	age int

}

// Embedding two different Structures: Struct within Struct

type secretAgent struct{
	person
	license_to_kill bool
}
func main(){
	fmt.Println("Understanding Structures.................................")
	p1 :=person{
		first_name:"James",
		last_name:"Bond",
		age:32,
	}
	p2 :=person{
		first_name:"Miss",
		last_name:"MoneyPenny",
		age:27,
	}
	sa1 :=secretAgent{
		person:p1
	}
	fmt.Println("Structure of Person is given as:")
	fmt.Print("Object-1:")
	fmt.Println(p1)
	fmt.Print("First and Last Name are:",p1.first_name," ",p1.last_name)
	fmt.Print("\n Age is:",p1.age)

	fmt.Println()
	fmt.Print("Object-2:")
	fmt.Println(p2)
	fmt.Print("First and Last Name are:",p2.first_name," ",p2.last_name)
	fmt.Print("\n Age is:",p2.age)



}