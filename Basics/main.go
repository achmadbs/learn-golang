package main

import "fmt"

func constVariable() {
    const initConst = "constant variable"
    fmt.Println(initConst)

    const (
        firstConst = 365
        secondConst = int32(366)
        thirdConst = 3e20 / firstConst
    )
    fmt.Println(firstConst, secondConst, thirdConst)
}

func stringInterpolation() {
    fmt.Printf("I am %v Programmer", "Golang");
    fmt.Printf("I have %v years of experience", 4);
}

func main() {
    var a string = "initial string"
    fmt.Println(a)

    var b,c = 1, 2
    fmt.Println(b + c)

    var d = true
    fmt.Println(d)

    var e int
    fmt.Println(e)

    f := "short hand variable"
    fmt.Println(f)

    constVariable()
    stringInterpolation()
}