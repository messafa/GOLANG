package main

import "fmt"


func calc(x int, y int) (res1 int,res2 int) {
	res1 = x + y
	res2 = x - y
	return
}


func main() {
	fmt.Println(calc(5, 6))
}