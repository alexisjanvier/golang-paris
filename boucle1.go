package main

import "fmt"

func main() {
    if greeting := sayHello(); len(greeting) > 0 {
        fmt.Println(greeting)
    }
    fmt.Println(greeting)
}

func sayHello() string {
    var say string
    say = "Hello !"

    return say
}
