package main

import (
	"fmt"
	"os"

	"github.com/budougumi0617/godecov"
)

func main() {
	tok := os.Getenv("CODECOV_TOKEN")
	cli := godecov.NewClient(tok)
	cli.TestMethod()
	fmt.Println("main finished")
}
