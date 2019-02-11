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
	if len(os.Args) < 4 {
		log.Fatal("Need input args as owner name")
	}
	res, err := cli.GetCommit(os.Args[1], os.Args[2], os.Args[3])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("res =\n%#v\n", res)

	fmt.Println("main finished")
}
