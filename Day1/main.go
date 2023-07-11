package main

import (
	"fmt"
)

func main() {
	var a int
	var b int
	fmt.Println("Enter the values a and b")
	fmt.Scanf("%d", &a)
	fmt.Scanf("%d", &b)
	fmt.Printf("a value is %d\n", a)
	fmt.Printf("b value is %d\n", b)

	/* Arithmetic Operators
	    +	Adds two operands
		-	Subtracts second operand from the first
		*	Multiplies both operands
		/	Divides the numerator by the denominator
		%	Modulus operator; gives the remainder after an integer division*/
	fmt.Printf("The addition of two varialbes are %d\n", a+b)
	fmt.Printf("The subtraction of two varialbes are %d\n", a-b)
	fmt.Printf("The Multiplication of two varialbes are %d\n", a*b)
	fmt.Printf("The divison of two varialbes are %d\n", a/b)
	fmt.Printf("The divison remainder of two varialbes are %d\n", a%b)
	// Relational Operators
	if a == b {
		fmt.Printf("%d is equal to  %d\n", a, b)
	} else {
		fmt.Printf("%d is not equal to %d\n", a, b)
	}
	if a > b && a != b {
		fmt.Printf("%d is greater than %d\n", a, b)
	} else if a != b {
		fmt.Printf("%d is greater than %d\n", b, a)
	}

	if a >= b {
		fmt.Printf("%d is either  greater than or equal to  %d\n", a, b)
	}
	if a <= b {
		fmt.Printf("%d is either less than or equal to %d\n", a, b)
	}
}
