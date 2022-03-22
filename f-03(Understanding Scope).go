package main

import "fmt"

// Golang has no private scope(similar to C/C++ class private close)

// Outside the Package Level Variable
// These variables can be used outside the package. They need to be be declared with capital letters.
// Such variables can be exported.
// e.g A can be used outside the package main
var A int32=40


// Package Level Variable
var a int32 =32

func main() {
	// Variables within the same scope can not be declared multiple times.

	fmt.Print("Type and Value of Variable a declared at package/global level is :")
	fmt.Printf("%T %v",a,a)
	fmt.Println()

	// Note: Shadowing is a concept which specifies that precedence of local variables over package level variables if they share the same name.
	// Local Level Variable
	var a int32=40
	fmt.Print("Type and Value of Variable a declared at function/local level is :")
	fmt.Printf("%T %v",a,a)

//	Declare variables names with the following format (Naming Convention):
//	1. Camel Case: theUrl
//	2. Acronym: EmployeeName
// 	3. Keep short variable names

}