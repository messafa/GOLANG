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

	readFile("./files/file1.txt")
	writeMore("./files/file1.txt", "\nI'm doing fine, thank you!")
	readFile("./files/file1.txt")
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

	lenght, err := file.WriteString(content)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Lenght: %d bytes\n", lenght)
}


