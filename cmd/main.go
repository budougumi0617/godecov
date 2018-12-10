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
	if len(os.Args) < 2 {
		log.Fatal("Need input args as owner name")
	}
	res, err := cli.GetOwner(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("res =\n%#v\n", res.Repos[0].Cache.Commit.Totals)
	fmt.Printf("res =\n%#v\n", res.Repos[0].Cache.Commit)
	fmt.Println("main finished")
}
