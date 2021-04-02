package main

import (
	"context"
	"fmt"
	"log"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

func main() {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: "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"},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	// list all repositories for the authenticated user
	repos, _, err := client.Repositories.

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(repos)
}
