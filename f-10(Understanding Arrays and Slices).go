package main

import "fmt"

func main(){
	fmt.Println("Understanding Arrays Syntax..............")
	fmt.Println()
	// Note: - 1. An array has fixed size, which is needed during declaration.
	//		   2. An array is homogeneous in nature.
	fmt.Println("Understanding 1-D Arrays ..............")

	//	To create an array use the following syntax.
	//	var(keyword) <array name> =[<length of array>]<data type> {initializing with values}
	//	Example-1:- General Array Declaration

	var array_1 =[3]int {1,2,3}
	fmt.Print("Type and Value of Array is:")
	fmt.Printf("%T %v",array_1,array_1)
	fmt.Println()


	//	Example-2:- Creating An Empty Array
	var array_2 [3]int
	fmt.Println("Empty Array Declared(with Default Values):",array_2)

	//	Example-3:- Creating an array without giving length as parameter(figuring out length during initialization)
	var array_3 =[...]int {1,2,3,4,5,6,7,8,9,10}
	fmt.Println("Length of Array is:",len(array_3))
	fmt.Println("Value of Array is:",array_3)

	//	Example-4:- Creating an array of Strings
	var array_4 =[3]string {"1","2","3"}

	fmt.Printf("Array of Strings has value %v and is with type  %T:",array_4,array_4)

	fmt.Println()
	fmt.Println()
	fmt.Println("Understanding Array Iteration ..............")

	// To access values in array, it is similar to accessing the index. Example:
	array_4[0]="3"
	fmt.Printf("Array of Strings has value %v and is with type  %T:",array_4,array_4)


	//	Creating 2-D Arrays:
	//	To create a 2-D arrays, we use the following syntax:

	fmt.Println()
	fmt.Println()
	fmt.Println("Understanding 2-D Arrays ..............")

	//	Example-6:- Creating An Empty 2-D Array of 2*2 size
	var array_2D_1 =[2][2]int{ [2]int{0,1}, [2]int{2,3}}
	fmt.Println("Empty 2-D Array Declared(with Default Values):",array_2D_1)


	//	Example-6:- Creating An Empty 2-D Array of 3*3 size
	var array_2D_2 [3][3]int
	fmt.Println("Empty 2-D Array Declared(with Default Values):",array_2D_2)

	fmt.Println()
	fmt.Println("Understanding Array reference with pointers ..............")

	var array_ptr_1 = [4]int{1,2,3,4}
	fmt.Printf("Intially declared array has value  %v:",array_ptr_1)
	fmt.Println()

	array_ptr_2:=array_ptr_1
	array_ptr_2[0]=6
	fmt.Printf("After change original array has value %v:",array_ptr_1)
	fmt.Println()

	fmt.Printf("After change duplicate array has value %v:",array_ptr_2)
	fmt.Println()

	array_ptr_3:= &array_ptr_1
	array_ptr_3[0]=6
	fmt.Printf("After change original array has value %v:",array_ptr_1)
	fmt.Println()

	fmt.Printf("After change duplicate array(2) has value %v:",array_ptr_3)
	fmt.Println()


}