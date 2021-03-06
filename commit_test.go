// Copyright (c) 2019. Yoichiro Shimizu @budougumi0617

package godecov_test

import (
	"os"
	"testing"

	"github.com/budougumi0617/godecov"
)

func TestClient_GetCommit(t *testing.T) {
	// Confirm degradation
	user := "budougumi0617"
	repo := "gopl"
	sha := "d4aa993"
	tok := os.Getenv("CODECOV_API_TOKEN")
	cli := godecov.NewClient(tok)
	_, err := cli.GetCommit(user, repo, sha)
	if err != nil {
		t.Fatal(err)
	}
}

func TestClient_GetCommitEmptyUser(t *testing.T) {
	// Confirm degradation
	user := ""
	repo := "gopl"
	sha := "d4aa993"
	tok := os.Getenv("CODECOV_API_TOKEN")
	cli := godecov.NewClient(tok)
	_, err := cli.GetCommit(user, repo, sha)
	if err == nil {
		t.Fatal("expect error")
	}
}

func TestClient_GetCommitEmptyRepo(t *testing.T) {
	// Confirm degradation
	user := "budougumi0617"
	repo := ""
	sha := "d4aa993"
	tok := os.Getenv("CODECOV_API_TOKEN")
	cli := godecov.NewClient(tok)
	_, err := cli.GetCommit(user, repo, sha)
	if err == nil {
		t.Fatal("expect error")
	}
}

func TestClient_GetCommitEmptySha(t *testing.T) {
	// Confirm degradation
	user := "budougumi0617"
	repo := "gopl"
	sha := ""
	tok := os.Getenv("CODECOV_API_TOKEN")
	cli := godecov.NewClient(tok)
	_, err := cli.GetCommit(user, repo, sha)
	if err == nil {
		t.Fatal("expect error")
	}
}
