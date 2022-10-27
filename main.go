package main

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("hexdump.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() { // internally, it advances token based on sperator
		original, err := base64.StdEncoding.DecodeString(scanner.Text())
		if err != nil {
			panic(err)
		}
		fmt.Println(original) // token in unicode-char
	}

}
