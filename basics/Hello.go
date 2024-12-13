package main // package main is special package that tells the Go compiler that the package should compile as an executable program instead of a shared library



import "fmt" // fmt is package that contains functions for formatted I/O operations



func main() {
    fmt.Println("hello Hisoka")
	fmt.Print("hello again Hisoka\n")
	fmt.Printf("%+v\n", "hello Hisoka")
}