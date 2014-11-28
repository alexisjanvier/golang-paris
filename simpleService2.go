package main

import (
    "net/http"
    "fmt"
    "errors"
)

// START OMIT
func processRequest(response http.ResponseWriter, request *http.Request) {
    defer func() {
        if panicError := recover(); panicError != nil {
            errorMsg := fmt.Sprintf("%q", panicError)
            http.Error(response, errorMsg, 400)
        }
    }()
    if request.Method != "POST" {
        http.Error(response, "You must send your request in POST.", 405)
        return
    }
    if request.URL.Path != "/golang-paris" {
        errorMsg := fmt.Sprintf("%s is not a valid url", request.URL.Path)
        panic(errors.New(errorMsg))
    }
    fmt.Fprintf(response, "Hello Golang Paris")
}
// END OMIT

func main() {
    http.HandleFunc("/", processRequest)
    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        fmt.Println(err)
    }
}
