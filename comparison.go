package main

import (
	"fmt"
	"generic/operation"
)

func main() {

	// default it search in whole array
	y := []int32{1, 2, 3, 4, 5, 1}
	_, in, er := (operation.Presence(y, int32(5))) // panic or  error message
	fmt.Println(in)
	x := []string{"as", "fs", "wsf", "wd"}

	// starting index is given
	p, in, er := (operation.Presence(x, "fs", 3)) // panic or  error message
	fmt.Println(p, in, er)

	// when start and end index both are present
	p, in, er = (operation.Presence(x, "fs", 1, 3)) // panic or  error message
	fmt.Println(p, in, er)

	// Error when more than 4 aurguments
	_, _, er = (operation.Presence(x, "fs", 1, 3, 5)) // panic or  error message
	fmt.Println(er)

}
