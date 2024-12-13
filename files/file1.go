package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	var content string
	content = "Hello\nHow are you doing?"
	file, err := os.Create("./files/file1.txt")
	if err != nil {
		log.Fatal(err)
	}

	length, err := file.WriteString(content)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Length: %d bytes\n", length)
	defer file.Close()

	readFile("./files/file1.txt")

	// Ask user for more content
	var moreContent string
	again := true
	for again { // Change while to for loop
		fmt.Print("Do you want to add more content? (y/n): ")
		fmt.Scan(&moreContent)
		if moreContent == "y" {
			fmt.Print("Enter more content: ")
			fmt.Scan(&content) // You may want to append instead of overwriting content
			writeMore("./files/file1.txt", content)
			readFile("./files/file1.txt")
		} else {
			fmt.Println("Goodbye!")
			again = false
		}
	}
}

func readFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		log.Fatal(err)
	}

	bs := make([]byte, stat.Size())
	_, err = file.Read(bs)
	if err != nil {
		log.Fatal(err)
	}

	str := string(bs)
	fmt.Println(str)
}

func writeMore(filename string, content string) {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	length, err := file.WriteString(content)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Length: %d bytes\n", length)
}
