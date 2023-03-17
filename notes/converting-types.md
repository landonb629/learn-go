# converting types in GO
the main package used for converting types is: strconv 


Use case: Command line arguments 
- os.Args = array representing the command line arguments that you pass

How to find the type? 
``` Printf("%T, os.Args[1]) ``` 

- the reflection package 
``` fmt.Println(reflect.TypeOf(arg)) ``` 
This will output the data type of the argument 

_ is the don't care symbol, this means Atoi will return an error and the value, we are ignoring the error 

How do we handle the error in go?
- change the _ to err 
- do an if check to see if err is not nil, then we know that we have an error

Integer to String 
- using the Itoa method of strconv