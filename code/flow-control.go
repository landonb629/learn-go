package main 

import "fmt"

func main() { 
	landonAge := 24 
	testScore := 50 

	if testScore < 0 { 
		fmt.Println("test score cannot be less than 0")
	} else if testScore > 0 && testScore < 50 { 
		fmt.Println("you failed")
	} else if testScore > 50 && testScore < 75 { 
		fmt.Println("you got a D")
	} else {
		fmt.Println("I am tired of writing if statements")
	}

	if landonAge > 25 { 
		fmt.Println("getting old nerd")
	} else if landonAge == 24 { 
		fmt.Println("almost at your quarter life crisis")
	} else { 
		fmt.Println("enjoy your youth")
	}
}