package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	alg := flag.String("algorithm", "sha256", "the name of hash algorithm")

	flag.Parse()

	if *alg == "sha256" {
		flag.Usage()
	}

	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	switch *alg {
	case "sha256":
		fmt.Fprintln(os.Stdout, sha256.Sum256(data))
	case "sha384":
		fmt.Fprintln(os.Stdout, sha512.Sum384(data))
	case "sha512":
		fmt.Fprintln(os.Stdout, sha512.Sum512(data))

	default:
		fmt.Fprintln(os.Stderr, "please supply sha256, sha384 or sha512")
	}
}
