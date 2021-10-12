package main

import "fmt"

func main() {
	s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	result := ""
	for i := range(s) {
		oddOrEven := i % 2 == 0
		if oddOrEven == true {
			result = "even"
		} else {
			result = "odd"
		}
		fmt.Printf("%v is %v\n", i, result)
	}
}
