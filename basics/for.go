package main

import "fmt"


var array = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)


	for j := range array {
		fmt.Println(j)
	}
}
