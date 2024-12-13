package main

import "fmt"

// vars in Golang
var int1 int
var bool1 bool
var string1 string
var float1 float64
var c complex128 



func main() {
	fmt.Printf("The default value of int is %d\n", int1)         // ==> O is the default value of int
	fmt.Printf("The default value of bool is %t\n", bool1)		 // ==> false is the default value of bool
	fmt.Printf("The default value of string is %s\n", string1)	 // ==> "" is the default value of string
	fmt.Printf("The default value of float is %f\n", float1)	 // ==> 0.000000 is the default value of float
	fmt.Printf("The default value of complex is %v\n", c)		 // ==> (0+0i) is the default value of complex

}