package main

import "fmt"

func main(){
	fmt.Println(" Understanding Looping Constructs.......................")
	// Go doesn't have while loops

	//	Syntax for Loops:
	fmt.Println(" Basic Loop.......................")
	for i:=0; i<10; i++{
		fmt.Print("\t",i)
	}
	fmt.Println()
	fmt.Println(" Printing star Patterns.......................")
	//	Printing Star
	for i:=0;i<5;i++{
		for j:=0;j<5;j++{
			if i>=j{
				fmt.Print("*")
			}else{
				continue
			}
		}
		fmt.Println()
	}

}

