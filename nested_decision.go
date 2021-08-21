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
}