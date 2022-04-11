package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	yyjson "github.com/zchee/yyjson-go"
)

func init() {
	log.SetFlags(log.Lshortfile)
}

func main() {
	file := filepath.Join("testdata", "example.json")

	buf, err := os.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	var v []interface{}
	if err := yyjson.Unmarshal(buf, v); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("v: %#v\n", v)

	if err := yyjson.ReadFile(file); err != nil {
		log.Fatal(err)
	}
}
