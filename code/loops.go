package main 

import "fmt"
import "strconv"
//import "os"

/*
example code: 
	- input a number and the program will run as long as the number you input is not 10
*/
func main() { 
  var i string
  var running = true

  for running { 
	fmt.Println("input a number: ")
  	fmt.Scan(&i) 
	converted := convertNum(i)
	if converted == 10 { 
		break;
	}
	fmt.Printf("your value was %d", converted)
	fmt.Println("\n")
  }
}

// converting the input to an integer
func convertNum(input string) int { 
  convert, err := strconv.Atoi(input)
  if err != nil { 
	panic(err)
  }
  return convert
} 