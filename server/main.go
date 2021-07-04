package main

import (
	"fmt"
	"io/ioutil"

	"github.com/dbanck/jest-diff-parser/internal/transformer"
)

func main() {
	content, err := ioutil.ReadFile("examples/exampleV1.diff")

	if err != nil {
		panic(err)
	}

	fmt.Print(transformer.Transform(string(content)))

}
