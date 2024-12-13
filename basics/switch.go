package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano()) // seed the random number generator
	x := rand.Intn(6)  // generate a random number between 0 and 9

	switch x {
	case 0:
		fmt.Println("x is 0")
	case 1:
		fmt.Println("x is 1")

	case 2:	
		fmt.Println("x is 2")

	case 3:
		fmt.Println("x is 3")

	case 4:
		fmt.Println("x is 4")

	case 5:
		fmt.Println("x is 5")

	default:
		fmt.Println("x is greater than 5 ==> ", x)

	}

}