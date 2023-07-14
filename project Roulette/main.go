package main

import (
	"fmt"
	"math/rand"
)

var i int
var k string
var l string
var m int
var j int
var o int

func one() int {
	fmt.Println("Enter the amount you want to load")
	fmt.Scanf("%d\n", &i)
	return i
}

func main() {
	one()
	// This for loop runs for many times until the amount you loaded is 0
	for i != 0 {
		fmt.Printf("******************************************************\n")
		fmt.Printf("Choose from the following options:\n 1. bet on red or black \n 2. bet on high or low (low is 1 -- 18, high is 19 -- 36) \n 3. bet on a specific number \n")
		fmt.Scanf("%d\n", &j)
		fmt.Println("Enter the amount you want to bet/enter '0' if no bet")
		fmt.Scanf("%d\n", &o)
		if o <= i {
			fmt.Println("The bet amount is :", o)
			u := rand.Intn(37)
			if j == 1 {
				fmt.Println("Enter the colour black and red")
				fmt.Scanf("%s\n", &k)
				//checks for even and odd
				if k == "red" || k == "black" {
					if u%2 == 0 && k == "red" {
						fmt.Println("The outcome colour is red! Horray You won :", o)
						i = i + o
						fmt.Println("The updated income is: ", i)
					} else if u%2 != 0 && k == "black" {
						fmt.Println("The outcome colour is black! Horray You won ", o)
						i = i + o
						fmt.Println("The updated income is: ", i)
					} else {
						fmt.Println("The outcome colour is the opposite! You Lost ", o)
						i = i - o
						fmt.Println("The updated income is: ", i)
					}
				}
			} else if j == 2 {
				fmt.Println("Enter the  bet on high or low (low is 1 -- 18, high is 19 -- 36)")
				fmt.Scanf("%s\n", &l)
				//check if high or low
				fmt.Println(" The number is ", u)
				if l == "high" || l == "low" {
					if (0 < u) && (u < 19) && l == "low" {
						fmt.Println("The number is in low range! Horray You won :", o)
						i = i + o
						fmt.Println("The updated income is: ", i)
					} else if (20 < u) && (u < 37) && l == "high" {
						fmt.Println("The number is in high range! Horray You won :", o)
						i = i + o
						fmt.Println("The updated income is: ", i)
					} else {
						fmt.Println("Number not within Range!! you lost:", o)
						i = i - o
						fmt.Println("The updated income is: ", i)
					}

				}

			} else if j == 3 {
				fmt.Println("Enter the bet on a specific number (1-36)")
				fmt.Scanf("%d\n", &m)
				//check if specific number is choosen
				fmt.Println(" The number is ", u)
				if u == m {
					fmt.Println("You have Won by matching numbers! and the amount won is", 36*o)
					i = i + 35*o
					fmt.Println("The updated income is: ", i)
				} else {
					fmt.Println("The number doesn't match! You Lost ", o)
					i = i - o
					fmt.Println("The updated income is: ", i)
				}

			}
		} else if o > i {
			fmt.Println("Your Bet Amount exceeds your Income!")
		} else if o == 0 {
			fmt.Println("You didn't participate in the bet")
		}
	}
}
