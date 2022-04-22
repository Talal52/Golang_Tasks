package main

import "fmt"

func main() {
	h := 0
	count := 0
	for i := 0; i < 200; i++ {
		if i == 0 || i == 1 {
			continue
		}
		h = 0
		for j := 2; j <= i/2; j++ {
			a := i % j
			if a == 0 {
				h = 1
				break
			}
		}
		if h == 0 {
			count++
			fmt.Print(i, " ")
		}
	}
	fmt.Println("")
	fmt.Println("Total Prime Numbers : ", count)
}
