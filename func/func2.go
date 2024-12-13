package main

import "fmt"


func calc(x int, y int) (int, int) {
	return x + y, x - y
	
}


func main() {
	fmt.Println(calc(5, 6))
}