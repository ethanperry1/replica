package main

import (
	replica "replica/parser"
	"go/parser"
	"go/token"
	"os"
)

func main() {
	err := run()
	if err != nil {
		panic(err)
	}
}

func run() error {
	fset := token.NewFileSet()
	file := os.Getenv("GOFILE")
	f, err := parser.ParseFile(fset, file, nil, parser.ParseComments)
	if err != nil {
		return err
	}

	replica.ParseFile(f)

	return nil
}