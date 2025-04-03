package main 

import (
        "fmt"
        "log"
        "os"
)

func main() {

        file, err := os.Create("./file2.txt")
        if err != nil {
                log.Fatal(err)
        }

        var content string
        fmt.Print("Enter content: ")
        // write string with spaces 
        fmt.Scanln(&content)
        fmt.Println(content)


        length, err := file.WriteString(content)

        if err != nil {
                log.Fatal(err)
        }

        fmt.Printf("Length: %d bytes\n", length)

        defer file.Close()




}