package main

import (
	"fmt"
	"go/parser"
	"go/token"
	"io"
	"os"
	replica "replica/parser"
	"strings"
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

	renderer, err := replica.NewRenderer()
	if err != nil {
		return err
	}

	writer := replica.NewWriter(renderer)

	creator := replica.New()

	mockFile := creator.CreateMockFile(f)

	reader, err := writer.Generate(mockFile)
	if err != nil {
		return err
	}
	
	content, err := io.ReadAll(reader)
	if err != nil {
		return err
	}

	mock, err := os.Create(fmt.Sprintf("%s_mocks.go", strings.Trim(file, ".go")))
	if err != nil {
		return err
	}

	_, err = mock.Write(content)

	return err
}