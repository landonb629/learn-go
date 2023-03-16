# Flow Control

## The If construct

 ```
 if true { 
     // run this code
 }
 ```

 ## multiple expressions 
 - you can use the && || operators 
 - the if conditions are not wrapped in parenthesis 

 example 

 ```
 landonAge := 24 
 if landonAge > 25 { 
     fmt.Println("getting old nerd")
 } else if landonAge = 24 {
     fmt.Println("almost to your quarter life crisis")
 } else { 
     fmt.Println("youth")
 }

 ```