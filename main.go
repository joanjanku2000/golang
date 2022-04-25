package main

import "fmt" // from go standard library:
// formatting strings and printing messages to the console
// GO - strongly typed compiled lang , variable type cannot be changed

func main(){
	// Cannot compile code with unused vars or imports
  var nameOne string = "joan" // var variableName type

		var nameTwo = "Luigi" // infers the type
 
		var name3 string; // default empty string

		fmt.Println("Name "+nameOne)

		fmt.Println("Second name "+nameTwo + " name 3 " + name3)

		name3 = nameOne;

			fmt.Println("Now name3 is " + name3)

			name4 := "yoshi" // not using var but using : , only done when initializing , only inside functions

			
				fmt.Println("Name4 is " + name4)

				// numbers - integers floats

				var age1 int = 22
				var age2 = 30
				age3 := 40

				fmt.Println(age1 , age2 , age3)

	 	//integer bit size (8,16,32,64) & memory

			var numOne int8 = 25 // compile error if given number exceeds int bit number
			var numTwo int64 = 45

			// unsigned int 

			var numThree uint = 12 //uint8 , uint16 , uint32 , uint64

			fmt.Println(numOne , numTwo , numThree)


			//floats

			var floatOne float32 = 1.4 // 32 or 64 even -2.3
		 floatSixtyFour := 123.6
		fmt.Println(floatOne ,floatSixtyFour)

}