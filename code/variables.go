package main 

import "fmt"

func main() { 
	var (
		name = "landon"
		age = 24 
		occupation = "Devops Engineer"
	)

	fmt.Println(name, age, occupation)
	fmt.Printf("my name is %s, and I am %d years old, my occupation is %s", name, age, occupation)
	
}