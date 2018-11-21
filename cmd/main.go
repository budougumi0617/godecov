package main

import (
	"fmt"
	"log"
	"os"

	"github.com/budougumi0617/godecov"
)

func main() {
	tok := os.Getenv("CODECOV_TOKEN")
	cli := godecov.NewClient(tok)
	res, err := cli.GetOwner("budougumi0617")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("res =\n%#v\n", res.Repos[0].Cache.Commit.Totals)
	fmt.Printf("res =\n%#v\n", res.Repos[0].Cache.Commit)
	fmt.Println("main finished")
}
