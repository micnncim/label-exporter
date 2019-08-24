package exporter

import (
	"context"
	"errors"
	"os"

	"github.com/google/go-github/v28/github"
	"golang.org/x/oauth2"
)

type githubClient struct {
	client *github.Client
}

func NewClient() (*githubClient, error) {
	token := os.Getenv("GITHUB_TOKEN")
	if token == "" {
		return nil, errors.New("missing GITHUB_TOKEN")
	}
	cli := newClient(token)
	return &githubClient{
		client: cli,
	}, nil
}

func newClient(token string) *github.Client {
	ts := oauth2.StaticTokenSource(&oauth2.Token{
		AccessToken: token,
	})
	tc := oauth2.NewClient(context.Background(), ts)
	return github.NewClient(tc)
}
