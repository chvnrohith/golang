package main

import (
	"fmt"
)

// Arithmetic Operations using Functions

func add(x int, y int) int {
	return x + y
}
func sub(x int, y int) int {
	return x - y
}
func mul(x int, y int) int {
	return x * y
}
func div(x int, y int) int {
	return x / y
}
func mod(x int, y int) int {
	return x % y
}

// Arithmetic Operations using Methods
type Vertex struct {
	b, c int
}

func (v Vertex) add() int {
	return v.b + v.c
}
func (v Vertex) sub() int {
	return v.b - v.c
}
func (v Vertex) mul() int {
	return v.b * v.c
}
func (v Vertex) div() int {
	return v.b / v.c
}
func (v Vertex) mod() int {
	return v.b % v.c
}

// Logical Operations Using Pointers with the help of functions

func Re(b int, c int) string {
	var p *int
	var q *int
	p = &b
	q = &c

	if *p == *q {
		return " Both the values are equal"
	}
	if *p > *q {
		return " first value is greater than second value"
	} else {
		return " Second value is greater than first value"
	}

}
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

	// Arithmetic Operators using Pointers d

	var d = &b
	var e = &c

	fmt.Printf(" Arithmetic Operators Using Pointers ------------ \n")
	fmt.Printf("The addition of two varialbes using pointers are %d \n", *d+*e)
	fmt.Printf("The subtraction of two varialbes using pointers are %d\n", *d-*e)
	fmt.Printf("The Multiplication of two varialbes using pointers are %d\n", *d**e)
	fmt.Printf("The divison of two varialbes using pointers are %d\n", *d / *e)
	fmt.Printf("The divison remainder of two varialbes using pointers are %d\n", *d%*e)

	fmt.Printf("Relational Operators Using Pointers ------------ \n")

	if *d == *e {
		fmt.Printf("%d is equal to  %d\n", *d, *e)
	} else {
		fmt.Printf("%d is not equal to %d\n", *d, *e)
	}
	if *d > *e && *d != *e {
		fmt.Printf("%d is greater than %d\n", *d, *e)
	} else if *d != *e {
		fmt.Printf("%d is greater than %d\n", *d, *e)
	}

	if *d >= *e {
		fmt.Printf("%d is either  greater than or equal to  %d\n", *d, *e)
	}
	if *d <= *e {
		fmt.Printf("%d is either less than or equal to %d\n", *d, *e)
	}

	fmt.Println("Relational Operators using pointers with the help of functions")

	fmt.Println(Re(*d, *e))

	fmt.Println("Operations Using Functions with the help of pointers------------")
	fmt.Printf("The addition of two numbers using function %d \n", add(*d, *e))
	fmt.Printf("The subtraction of two numbers using function %d \n", sub(*d, *e))
	fmt.Printf("The multiplication of two numbers using function %d \n", mul(*d, *e))
	fmt.Printf("The divison of two numbers using function %d \n", div(*d, *e))
	fmt.Printf("The modulus of two numbers using function %d \n", mod(*d, *e))

	fmt.Printf("Arithmetic Operations Using Methods --------------\n")

	v := Vertex{b, c}
	fmt.Println(v.add())
	fmt.Println(v.sub())
	fmt.Println(v.mul())
	fmt.Println(v.div())
	fmt.Println(v.mod())

}
