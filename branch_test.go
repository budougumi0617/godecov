// Copyright (c) 2019. Yoichiro Shimizu @budougumi0617

package godecov_test

import (
	"os"
	"testing"

	"github.com/budougumi0617/godecov"
)

func TestClient_GetBranch(t *testing.T) {
	// Confirm degradation
	user := "budougumi0617"
	repo := "gopl"
	br := "master"
	tok := os.Getenv("CODECOV_API_TOKEN")
	cli := godecov.NewClient(tok)
	_, err := cli.GetBranch(user, repo, br)
	if err != nil {
		t.Fatal(err)
	}
}
