package main

import (
	"crypto/sha256"
	"encoding/hex"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
)

type httpBase struct {
	Name   string
	Urls   []string
	Sha256 string
}

func main() {
	rule := flag.String("rule", "", "Variant of http rule (file, archive, jar)")
	name := flag.String("name", "", "Name of the rule")
	flag.Parse()
	urls := flag.Args()

	if *rule == "" {
		log.Fatal("Rule type was not specified!")
	}
	if *name == "" {
		log.Fatal("Name was not specified!")
	}
	if len(urls) == 0 {
		log.Fatal("No URLs were provided!")
	}

	hasher := sha256.New()

	resp, err := http.Get(urls[0])
	if err != nil {
		log.Fatal(err)
	}

	_, err = io.Copy(hasher, resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	shaString := hex.EncodeToString(hasher.Sum(nil))

	fmt.Printf("http_%s(\n", *rule)
	fmt.Printf("    name = \"%s\",\n", *name)
	fmt.Printf("    sha256 = \"%s\",\n", shaString)
	fmt.Printf("    urls = [\n")
	for _, url := range urls {
		fmt.Printf("        \"%s\",\n", url)
	}
	fmt.Printf("    ],\n")
	fmt.Printf(")\n")
}
