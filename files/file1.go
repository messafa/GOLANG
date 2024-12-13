package main

import (
	"fmt"
	"log"
	"os"
)


func main() {
	var content string
	content = "Hello Hisoka\nHow are you doing?"
	file,err := os.Create("./files/file1.txt")
	if err != nil {
		log.Fatal(err)
	}

	lenghth, err := file.WriteString(content)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Lenght: %d bytes\n", lenghth)
	defer file.Close()
}