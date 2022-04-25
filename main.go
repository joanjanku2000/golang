package main

// from go standard library:
// formatting strings and printing messages to the console
// GO - strongly typed compiled lang , variable type cannot be changed
import "fmt"


func main(){
		
	fmt.Print("Hello , ")
fmt.Print("World  \n") // just like C and Java

fmt.Println() // adds \n

fmt.Print("My age is ",22 , " \n")

age:= 22;  name:= "Joan";

fmt.Printf("My age is %v and my name is %v \n",age,name);
fmt.Printf("My age is %q and my name is %q \n",age,name); // output quotes on strings
fmt.Printf("My age is %v and of type is %T \n",age,age);
fmt.Printf("You scored %0.2f points ! ",2.226) // limiting to 2 decimal points


// Sprintf (save formatted strings)

var str = fmt.Sprintf("My age is %v and my name is %v \n",age,name) // doesn't print it only returns it
fmt.Print("\n The saved string is " ,str)

}