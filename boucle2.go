package main

import "fmt"

func main() {
    var greeting string
    if greeting = sayHello(); len(greeting) > 0 {
        fmt.Println(greeting)
    }
    fmt.Println(greeting)
}

func sayHello() string {
    var say string
    say = "Hello !"

    return say
}
