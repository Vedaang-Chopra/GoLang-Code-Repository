package main

import "fmt"
import "strconv"

func main(){
	fmt.Println("Understanding Type Casting..............................................................")
	var var_1 int32
	var_1=1
	fmt.Print("Type and Value of Variable-1:")
	fmt.Printf("%T %v",var_1,var_1)
	fmt.Println()

	//	To convert an integer to float:
	//	Pass the value of variable to float constructor

	var var_2 float32= float32(var_1)
	fmt.Print("Type and Value of Variable-2(Float32 Typecasting from Int32) is:",var_2)
	fmt.Printf("%T %v",var_2,var_2)
	fmt.Println()

	//	To convert an integer to string

	var var_3 string= string(var_1)
	fmt.Print("Type and Value of Variable-3(String Typecasting from Int32) is:",var_3)
	fmt.Printf("%T %v",var_3,var_3)
	fmt.Println()

	// The conversion is not done properly. To do so an external package is needed strconv.
	// The strconv package has the function Itoa which takes an int datatype(not int32 or int64 only int!!!!) as input and returns it as string

	// Method-1:- Fast
	var var_4 string
	var_4= strconv.Itoa(int(var_1))
	fmt.Print("Type and Value of Variable-4(String Typecasting from Int32 but using strconv package and Itoa function)  is:")
	fmt.Printf("%T %v",var_4,var_4)
	fmt.Println()

	// Method-2:- Slow
	var_4= fmt.Sprint(var_1)
	fmt.Print("Type and Value of Variable-4(String Typecasting from Int32 but using fmt package and sprint function)  is:")
	fmt.Printf("%T %v",var_4,var_4)
	fmt.Println()

	// Method-3:- Faster
	// The strconv package has the function FormatInt which takes an int64 datatype(not int32 or int only int64!!!!) as input and returns it as string
	var_4= strconv.FormatInt(int64(var_1),10)
	fmt.Print("Type and Value of Variable-4(String Typecasting from Int32 but using strconv package and FormatInt function)  is:")
	fmt.Printf("%T %v",var_4,var_4)
	fmt.Println()

}