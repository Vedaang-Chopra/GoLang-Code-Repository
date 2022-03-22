package main

import "fmt"

// Variadic Parameters:- Function Parameters that taken any number of parameters of any type
// In go you can return multiple parameters from a function

func main()  {
	// Underscore helps us to receive returns which do not want to use. Values which can be dropped, and not be used.
	// "_" helps to throw away returns.
	n,_:=fmt.Println("Hello World")
	fmt.Println(n)

}