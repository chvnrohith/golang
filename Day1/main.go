package main

import (
	"fmt"
)

func main() {
	var a [2]int
	fmt.Println("Enter two values")
	fmt.Scanf("%d %d", &a[0], &a[1])
	fmt.Println(a)
	b := a[0]
	c := a[1]

	/* Arithmetic Operators
	    +	Adds two operands
		-	Subtracts second operand from the first
		*	Multiplies both operands
		/	Divides the numerator by the denominator
		%	Modulus operator; gives the remainder after an integer division */
	fmt.Printf("The addition of two varialbes are %d\n", b+c)
	fmt.Printf("The subtraction of two varialbes are %d\n", b-c)
	fmt.Printf("The Multiplication of two varialbes are %d\n", b*c)
	fmt.Printf("The divison of two varialbes are %d\n", b/c)
	fmt.Printf("The divison remainder of two varialbes are %d\n", b%c)
	// Relational Operators
	if b == c {
		fmt.Printf("%d is equal to  %d\n", b, c)
	} else {
		fmt.Printf("%d is not equal to %d\n", b, c)
	}
	if b > c && b != c {
		fmt.Printf("%d is greater than %d\n", b, c)
	} else if b != c {
		fmt.Printf("%d is greater than %d\n", c, b)
	}

	if b >= c {
		fmt.Printf("%d is either  greater than or equal to  %d\n", b, c)
	}
	if b <= c {
		fmt.Printf("%d is either less than or equal to %d\n", b, c)
	}
}
