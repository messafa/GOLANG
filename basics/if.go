package main 

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// x random number between 1 and 10
	rand.Seed(time.Now().UnixNano()) // seed the random number generator
	x := rand.Intn(10)  // generate a random number between 0 and 9
	if x > 5 {
		fmt.Println("x is greater than 5 ==> ", x)
	} else {
		fmt.Println("x is less than 5 ==> ", x)
	}


}