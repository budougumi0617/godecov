package main

import (
	"fmt"
	"log"
	"os"

	"github.com/budougumi0617/godecov"
)

func main() {
	tok := os.Getenv("CODECOV_API_TOKEN")
	cli := godecov.NewClient(tok)
	if len(os.Args) < 3 {
		log.Fatal("Need input args as owner name")
	}
	res, err := cli.GetBranches(os.Args[1], os.Args[2])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("res =\n%v\n", res)

	fmt.Println("main finished")
}
