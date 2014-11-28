package main

import (
    "net/http"
    "fmt"
)

func processRequest(response http.ResponseWriter, request *http.Request) {
    fmt.Println(request.URL.Path)
    fmt.Fprintf(response, "Hello gophers'")
}

func main() {
    http.HandleFunc("/", processRequest)
    http.ListenAndServe(":8080", nil)
}
