package main


import "fmt" 


func main() {
	// Taking two local variable
	var v1 int = 400
	var v2 int = 700

	// using if statement
	if (v1 == 400) {

		// if condition is true then check the following
		if (v2 == 700) {

			// if condition true then display the following
			fmt.Printf("Value v1 is %d and value v2 is %d\n", v1, v2)
		}
	}

	// use of if..else..if ladder
	// taking a local variable
	var v int = 700

	// checking the condition
	if (v == 100) {

		// if condition is true then display the following
		fmt.Printf("Value of v is 100\n")
	} else if(v == 200) {

		fmt.Printf("Value of v is 200\n")

	} else if(v == 300) {

		fmt.Printf("Value of v is 300\n")

	} else {

		// if none of the condition is true
		fmt.Printf("None of the values is matching\n") 
	}
}