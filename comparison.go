package main

import (
	"fmt"
	"generic/operation"
)

func main() {

	/* test cases for presence of element  */
	// default it search in whole array
	// y := []int{1, 2, 3, 4, 5, 1}
	// pp, in, er := operation.Presence(y, int64(3)) // panic or  error message
	// fmt.Println(pp, in, er)
	// x := []string{"as", "fs", "wsf", "wd"}

	// starting index is given
	// p, in, er := operation.Presence(x, "fs", 3) // panic or  error message
	// fmt.Println(p, in, er)

	// when start and end index both are present
	// p, in, er = operation.Presence(x, "fs", 1, 5) // panic or  error message
	// fmt.Println(p, in, er)

	// Error when more than 4 aurguments
	// _, _, er = operation.Presence(x, "fs", 1, 3, 5) // panic or  error message
	// fmt.Println(er)

	/* Test cases for deletion of element*/
	y := []int32{1, 2, 3, 4, 5, 1}

	// output : [1 2 3 4 5 1] 6 You can't use more than 3 aurguments in DeleteEle function
	// p, q, r := operation.DeleteEle(y, int32(1), 12, 3)
	// fmt.Println(p, q, r)

	// Delete All  1
	p, q, r := operation.DeleteEle(y, int32(1))
	fmt.Println(p, q, r)

	/* Test caes for pop back */
	// y := []int64{1, 2, 3, 4, 5, 1, 8}
	// a, b, err := operation.PopBack(y)
	// fmt.Println(a, b, err)

	// a, b, err = operation.PopBack(5)
	// fmt.Println(a, b, err)

	// fmt.Printf("%T", b)
}

// wget --mirror -p --convert-links -P ./LOCAL-DIR WEBSITE-URL
