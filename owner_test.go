// Copyright (c) 2019.  Yoichiro Shimizu @budougumi0617

package godecov_test

import (
	"os"
	"testing"

	"github.com/budougumi0617/godecov"
)

func TestClient_GetOwner(t *testing.T) {
	// Confirm degradation
	user := "budougumi0617"
	tok := os.Getenv("CODECOV_TOKEN")
	cli := godecov.NewClient(tok)
	_, err := cli.GetOwner(user)
	if err != nil {
		t.Fatal(err)
	}
}
