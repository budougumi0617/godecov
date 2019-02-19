// Copyright (c) 2019. Yoichiro Shimizu @budougumi0617

package godecov_test

import (
	"os"
	"testing"

	"github.com/budougumi0617/godecov"
)

func TestClient_GetBranches(t *testing.T) {
	// Confirm degradation
	user := "budougumi0617"
	repo := "godecov"
	tok := os.Getenv("CODECOV_API_TOKEN")
	cli := godecov.NewClient(tok)
	_, err := cli.GetBranches(user, repo)
	if err != nil {
		t.Fatal(err)
	}
}
