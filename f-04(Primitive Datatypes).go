package main

import "fmt"

	func string_variables() {
		fmt.Println("Understanding String Datatypes..........................")

		var var_1 string = "This is a string variable declared using Technique-1 of Variable Declaration."
		fmt.Print("Type of Variable-1 is:")
		fmt.Printf("%T", var_1)
		fmt.Println()
		fmt.Print("Value of Variable-1 is:")
		fmt.Printf("%v", var_1)
		fmt.Println()

		var var_2 string
		var_2 = " This is a string variable declared using Technique-2 of Variable Declaration."
		fmt.Print("Type of Variable-2 is:")
		fmt.Printf("%T", var_2)
		fmt.Println()
		fmt.Print("Value of Variable-2 is:")
		fmt.Printf("%v", var_2)
		fmt.Println()

		var_3 := "This is a string variable declared using Technique-3 of Variable Declaration."
		fmt.Print("Type of Variable-3 is:")
		fmt.Printf("%T", var_3)
		fmt.Println()
		fmt.Print("Value of Variable-3 is:")
		fmt.Printf("%v", var_3)
		fmt.Println()

		var var_4 string
		fmt.Print("Type of Variable-4 is:")
		fmt.Printf("%T", var_4)
		fmt.Println()
		// Default value: Empty String
		fmt.Print("Default value of String Variable is:")
		fmt.Printf("%v", var_4)
		fmt.Println()
	}
	func boolean_variables(){

		fmt.Println("Understanding Boolean Data types..........................")

		var var_1 bool=false
		fmt.Print("Type of Variable-1 is:")
		fmt.Printf("%T", var_1)
		fmt.Println()
		fmt.Print("Value of Variable-1 is:")
		fmt.Printf("%v", var_1)
		fmt.Println()

		var var_2 bool
		var_2=true
		fmt.Print("Type of Variable-2 is:")
		fmt.Printf("%T", var_2)
		fmt.Println()
		fmt.Print("Value of Variable-2 is:")
		fmt.Printf("%v", var_2)
		fmt.Println()

		var_3:= 1==1
		fmt.Print("Type of Variable-3 is:")
		fmt.Printf("%T", var_3)
		fmt.Println()
		fmt.Print("Value of Variable-3 is:")
		fmt.Printf("%v", var_3)
		fmt.Println()

		var var_4 bool
		fmt.Print("Type of Variable-4 is:")
		fmt.Printf("%T", var_4)
		fmt.Println()
		fmt.Print("Default value of String Variable is:")
		// Since default value of bool variables is 0 and 0 is considered as false
		fmt.Printf("%v", var_4)
		fmt.Println()


	}
func integer_variables(){

	fmt.Println("Understanding Integer Data types and its Different Types..........................")

	var var_1 int64=1
	fmt.Print("Type of Variable-1 is:")
	fmt.Printf("%T", var_1)
	fmt.Println()
	fmt.Print("Value of Variable-1 is:")
	fmt.Printf("%v", var_1)
	fmt.Println()

	var var_2 int32
	var_2=2
	fmt.Print("Type of Variable-2 is:")
	fmt.Printf("%T", var_2)
	fmt.Println()
	fmt.Print("Value of Variable-2 is:")
	fmt.Printf("%v", var_2)
	fmt.Println()

	var_3:= 3
	fmt.Print("Type of Variable-3 is:")
	fmt.Printf("%T", var_3)
	fmt.Println()
	fmt.Print("Value of Variable-3 is:")
	fmt.Printf("%v", var_3)
	fmt.Println()

	var var_4 int8
	fmt.Print("Type of Variable-4 is:")
	fmt.Printf("%T", var_4)
	fmt.Println()
	fmt.Print("Default value of String Variable is:")
	// Since default value of bool variables is 0 and 0 is considered as false
	fmt.Printf("%v", var_4)
	fmt.Println()
	fmt.Println("Other Interger Types include int16, uint32,uint16 etc..")



}
	func main(){
		fmt.Println("Understanding Primitive Data Types of GoLang....")
		fmt.Println()
		string_variables()
		fmt.Println()
		boolean_variables()
		fmt.Println()
		integer_variables()
	}

