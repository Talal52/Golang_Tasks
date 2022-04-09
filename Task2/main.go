package main

import "fmt"

func main() {
	a := []int{1, 2, 2, 4, 5}
	b := []int{}
	for i := len(a); i > 0; i-- {
		b = append(b, i)
	}
	fmt.Print(b)

}
