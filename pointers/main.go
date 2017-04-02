package main

import "fmt"

func main() {
	var (
		a *[]string
		b []string
	)

	/*
	pods = new([]int)
	fmt.Printf("value in pods: %#v\n", pods == nil)
	fmt.Printf("value in pods: %v\n", *pods)
	fmt.Printf("value in pods: %v\n", )
	fmt.Printf("value in pods: %#v\n", pods == nil)

	a = &[]string{}
	fmt.Printf("a %#v \n b %#v \n", a, b)

	*/

	c := []string{"5"}
	b = []string{"1", "2"}
	a = &b
	*a = append(*a, "3")
	*a = append(*a, "4")
	*a = c

	fmt.Printf("a %#v \n b %#v \n", a, b)
}
