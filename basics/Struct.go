package main

import "fmt"

type Person struct {
    Name  string
    Age   int
}

type Hanter struct {
    id int
    name string
}

func main() {
    // Array of Hanters
    hanters := []Hanter{
        {44, "Hisoka"},
        {405, "Gon"},
        {99, "Killua"},
    }

    p := Person{"Miryoum", 25}

    fmt.Println(p)

    fmt.Println(hanters)
}