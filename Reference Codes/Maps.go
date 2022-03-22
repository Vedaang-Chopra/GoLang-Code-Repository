package main
import "fmt"

func main(){
	fmt.Println("Understanding Maps.....................")
	/*
	Map is a Data Structure that maps different keys to different values.
	So basically it is a dictionary(python).
	The keys of maps can have the following data types: strings,list,pointers, arrays, numeric types, structs, channels,booleans etc.
	The keys of maps can't have the following data types: slices, maps and oth	er functions etc.

	*/
	// The Syntax for Declaring Map is:
	// var(keyword) <name of map variable> = map[<data type of keys>]<data type of values>{initializing values}
	var map_1 = map[string]int{
		"a":1,
		"b":2,
		"c":3,
	}
	fmt.Println(map_1)
}