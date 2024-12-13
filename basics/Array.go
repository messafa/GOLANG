package main

import "fmt"

var numbers [5]int = [5]int{1, 2, 3, 4, 5}
var names [3]string = [3]string{"Hisoka", "Gon", "Killua"}


func main() {
	for i := 0; i < len(numbers); i++ {
		fmt.Println(numbers[i])
	}
	fmt.Println("**********")
	fmt.Printf("%+v\n", names)
}