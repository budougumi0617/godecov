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
	if len(os.Args) < 3 {
		log.Fatal("Need input args as owner name")
	}
	res, err := cli.GetSingleRepository(os.Args[1], os.Args[2])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("AvatarURL =\n\"%+v\"\n", res.Owner.AvatarURL)
	fmt.Printf("res =\n%#v\n", res)

	reso, err := cli.GetOwner(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("res =\n%#v\n", reso)
	fmt.Println("main finished")
}
