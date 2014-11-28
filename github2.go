package main

import (
    "github.com/google/go-github/github"
    "fmt"
)

func main() {
    client := github.NewClient(nil)
    opt := &github.PullRequestListOptions{State: "closed"}
    prs, _, getPrError := client.PullRequests.List("marmelab", "ng-admin", opt)
    if getPrError != nil {
        fmt.Println(getPrError)
        return
    }
    for _, pr := range prs {
        if pr.MergedAt == nil {
            continue
        }
        fmt.Printf("- %s\n", *pr.Title)
    }
}
