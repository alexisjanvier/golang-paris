package main

import (
    "github.com/google/go-github/github"
    "fmt"
)

func main() {
    client := github.NewClient(nil)
    orgs, _, err := client.Organizations.List("alexisjanvier", nil)
    if err != nil {
        fmt.Println(err)
        return
    }
    for _, org := range orgs {
        fmt.Printf("- %s\n", *org.Login)
    }
}
