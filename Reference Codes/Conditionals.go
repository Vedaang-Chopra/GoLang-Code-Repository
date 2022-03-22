package main

import "fmt"

func main() {
	fmt.Println("Understanding Conditionals........................................")
	//	Syntax for Conditional:
	var_1, var_2 := 1, 2
	if true {
		fmt.Println("The value is", var_1)
	}
	if false {
		fmt.Println("The value is", var_2)
	}
	// !false not working
	if !(2 == 2) {
		fmt.Println("The value is", var_2)
	}

	// Adding Initialization statements in if statements
	if x := 42; x == 42 {
		fmt.Println("Value is Initialized")
	}
	//	Syntax for if else
	y:=4
	if y == 1 {
		fmt.Println("Print Value", var_1)

	} else {
		fmt.Println("Print Value", y)
	}
	fmt.Println()
	//demo()
	//if (y == 1) {
	//	fmt.Println("Print Value", var_1)
	//} else if (y == 2) {
	//	fmt.Println("Print Value", var_2)
	//} else {
	//	fmt.Println("Print Value", y)
	//}
}


func demo() {

	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}

	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
}