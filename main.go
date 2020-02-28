package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/nrdhm/huffman_compression/core"
)

type symbolCode struct {
	symbol string
	code   string
}

func do() {
	file, err := os.Open("some.txt")
	if err != nil {
		panic(err)
	}
	data, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}
	huf, err := core.Compress(string(data))
	if err != nil {
		panic(err)
	}
	fmt.Println(huf)
}

func undo() {
	file, err := os.Open("some.txt.huf")
	if err != nil {
		panic(err)
	}
	data, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}
	text, err := core.Decompress(string(data))
	if err != nil {
		panic(err)
	}
	fmt.Println(text)
}

func main() {
	// do()
	undo()
}
