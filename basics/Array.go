package main

import "fmt"

var numbers [5]int = [5]int{1, 2, 3, 4, 5}
var names [3]string = [3]string{"Hisoka", "Gon", "Killua"}


func main() {
	fmt.Printf("The numbers are %v\n", numbers)
	fmt.Printf("The names are %v\n", names)
}