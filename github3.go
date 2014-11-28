package main

import (
    "github.com/google/go-github/github"
    "fmt"
    "time"
)

func main() {
    now := time.Now()
    oneWeeksAgo := now.Add(-1 * 7 * 24 * time.Hour)
    client := github.NewClient(nil)
    opt := &github.CommitsListOptions{SHA: "master", Since: oneWeeksAgo}
    commits, _, err := client.Repositories.ListCommits("marmelab", "ng-admin", opt)
    if err != nil {
        fmt.Println(err)
        return
    }
    for _, commit := range commits {
        fmt.Printf("- %s\n", *commit.SHA)
    }
}
