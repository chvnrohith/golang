package main

import (
	"fmt"
	"strings"
)

func main() {
	mystr := "Hi i am an engineer i work as a frontend developer" //create strings
	fmt.Println("The string before removal of whitespaces is:", mystr)
	mystr = strings.Replace(mystr, " ", "", -1) //using this built-in function remove whitespaces
	fmt.Println("The string after removing whitespaces is:")
	fmt.Println(mystr) //remove string without whitespaces
}

/* package main

func main() {
	var s string
	fmt.println("Enter the string")
	fmt.scanf("%s", &s)
	fmt.println("The entered string is %v", s)
	var y string
	y =s.lower()
	s=s.lower()
	i:=0
	y:=len(y)-1
	if i<=j {
		if s[i]!=y[j] {
			return False
		}
		i+=1
		j-+1
	return True
	}






}