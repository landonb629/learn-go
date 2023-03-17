package main 

import (
	"os"
	"fmt"
	//"reflect"
	"strconv"
)

func main() { 
 arg1 := os.Args[1]
 number := 8

 printType(arg1) // function that prints the type of argument that it's passed
 fmt.Printf("\n")
 convertInt(arg1)
 fmt.Printf("\n")
 convertString(number)
 fmt.Printf("\n")
}

func printType(input string) string { 
	fmt.Printf("%T", input)
	return " "
}

func convertInt(num string) int {
	number, err := strconv.Atoi(num)
	if err != nil { 
	  fmt.Println("cannot convert %s to an integer", num)
	}  
	fmt.Printf("%T", number)
    return number
}

func convertString(num int) string { 
	str := strconv.Itoa(num)
	/*
	if err != nil { 
		fmt.Println("cannot convert %d to a string", num)
	}
	*/
	fmt.Printf("%T", str)
	return str
}