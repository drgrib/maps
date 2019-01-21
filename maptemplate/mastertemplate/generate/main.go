package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

const (
	masterFilePath = "mapstringint/map_string_int.go"
)

func main() {
	s := getMasterFileString()
	fmt.Println(s)
}

func getMasterFileString() string {
	b, err := ioutil.ReadFile(masterFilePath)
	if err != nil {
		log.Fatal(err)
	}
	return string(b)
}
