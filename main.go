package main

import (
    "fmt"
    // "first/test"
    // "first/vgarcia"
     "os"
     "bufio"
)

func main() {
    fmt.Println("test")
    for _, arg := range os.Args {
        fmt.Printf(" %s\n", arg)
    }
    x := "test"
    var yes bool
    var flaot float64
    fmt.Printf("%T\n", x)
    fmt.Println(yes)
    fmt.Println(flaot)
        num := -4
    if num > 0 {
        fmt.Println("positif")
    } else if num < 0 {
        fmt.Println("négatif")
    } else {
        fmt.Println("zéro")
    }
    defer fmt.Println("defer")
    var str string
    var truc bool
    test, testv2 := fmt.Scanln(&str, &truc)
    println(str)
    println(test, testv2)

    reader := bufio.NewReader(os.Stdin)
    fmt.Print("txt : ")
    line, _ := reader.ReadString('\n')
    fmt.Println("result :", line)
}