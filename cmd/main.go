package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/budougumi0617/godecov"
)

func main() {
	tok := os.Getenv("CODECOV_API_TOKEN")
	cli := godecov.NewClient(tok)
	if len(os.Args) < 4 {
		log.Fatal("Need input args as owner name")
	}
	no, _ := strconv.Atoi(os.Args[3])
	res, err := cli.GetPull(os.Args[1], os.Args[2], no)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("AvatarURL =\n\"%+v\"\n", res)
	fmt.Printf("res =\n%#v\n", res)

	fmt.Println("main finished")
}
