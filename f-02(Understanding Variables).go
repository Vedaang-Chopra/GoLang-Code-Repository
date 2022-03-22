package main

import "fmt"

// Package level variable, can't be declared by Technique-3.
// Need full declarations(that is datatype)
var package_variable float32=32.0

//To Declare Multiple values Together the following technique can be used:
var(
	demo_1 int32=20
	demo_2 float32=20.0
	demo_3 string="Demo"
)

// The go language is C style based. So it needs the main function to start.
//The starting curly brackets need to be present just after the function name(), otherwise it would cause error.
func main(){

	// Variables declared inside a function is function level.

	fmt.Println("There are the following information about Variable Declaration:")
	// Declaring variable has 3 different kinds:
	// var(keyword) <name of variable> <datatype>(Data type of variable)
	// Note:- All unused variables(variables with no assigned values or assigned values doesn't matter) are cause of error in golang.
	// Golang has garbage collector, and releases variable out of scope.
	// Variables within the same scope can not be declared multiple times.


	// Technique:-1 : Declaration followed by assignment
	// Used for Looping, or places where initialization is needed after declaration
	var var_1 int
	var_1=1
	fmt.Println("Variable-1:",var_1)

	// Technique:-2 : Declaration and assignment simultaneously
	var var_2 int=2
	fmt.Println("Variable-2:",var_2)

	// Technique:-3
	// Here golang automatically figures out the data type
	var_3:=3.

	fmt.Println("Variable-3:",var_3)
	// To run
	fmt.Printf("%T %v",var_3,var_3)					// %T is to find type of variable, %v its value.
	fmt.Println()

//	Different Data Types include:
	var var_4 float32=4.0
	fmt.Println("Variable-4:",var_4)

	fmt.Println("Package Scope/Level Variable:",package_variable)
	fmt.Println("Package Scope/Level Variable declared together:",demo_1,demo_2,demo_3)
}