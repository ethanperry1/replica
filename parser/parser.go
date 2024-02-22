package parser

import (
	"go/ast"
	os "os"
	"replica/example"
	"strings"
)

type MockFile interface {
	Imports() []string
	Mocks() []Mock
}

// replica
type Mock interface {
	Example() (int, example.Example[example.Other[os.File], os.File], example.Other[os.File], struct{}, func() interface{})
	Name() string
	Methods() []Method
}

type Method interface {
	Name() string
	Params() []Field
	Returns() []Field
}

type Field interface {
	Name() string
	Type() string
}

type Import struct {
	name string
	path string
}

type Entry struct {
	path       string
	references int
}

func ParseImports(specs []*ast.ImportSpec) map[string]*Entry {
	m := make(map[string]*Entry)

	for _, spec := range specs {
		imp := ParseImport(spec)
		m[imp.name] = &Entry{
			path: imp.path,
		}
	}

	return m
}

func ParseImport(spec *ast.ImportSpec) *Import {
	if spec.Name != nil {
		return &Import{
			name: spec.Name.Name,
			path: spec.Path.Value,
		}
	}

	parts := strings.Split(spec.Path.Value, "/")
	return &Import{
		name: parts[len(parts)-1],
		path: spec.Path.Value,
	}
}

func ParseFile(file *ast.File) ([]Mock, error) {
	var mocks []Mock
	for _, decl := range file.Decls {
		switch d := decl.(type) {
		case *ast.GenDecl:
			res, err := ParseGenDecl(d)
			if err != nil {
				return nil, err
			}
			mocks = append(mocks, res...)
		}
	}

	return mocks, nil
}

func ParseGenDecl(decl *ast.GenDecl) ([]Mock, error) {
	var mocks []Mock
	for _, spec := range decl.Specs {
		switch s := spec.(type) {
		case *ast.TypeSpec:
			res, ok := ParseTypeSpec(s)
			if ok {
				mocks = append(mocks, res)
			}
		}
	}

	return mocks, nil
}

func ParseTypeSpec(spec *ast.TypeSpec) (Mock, bool) {
	switch s := spec.Type.(type) {
	case *ast.InterfaceType:
		return ParseInterfaceType(s), true
	}

	return nil, false
}

func ParseInterfaceType(typ *ast.InterfaceType) Mock {
	for _, method := range typ.Methods.List {
		switch m := method.Type.(type) {
		case *ast.FuncType:
			ParseFuncType(m)
		}
	}

	return nil
}

func ParseFuncType(fun *ast.FuncType) Method {
	for _, param := range fun.Params.List {
		ParseParam(param)
	}
	for _, result := range fun.Results.List {
		ParseParam(result)
	}

	return nil
}

func ParseParam(*ast.Field) Field {
	return nil
}

func ParseResult(*ast.Field) Field {
	return nil
}

func ParseIdent(*ast.Ident) {

}

func ParseIdentExpr(*ast.IndexExpr) {

}

func ParseIdentListExpr(*ast.IndexListExpr) {

}

func ParseSelectorExpr(*ast.SelectorExpr) {

}

func ParseStructType(*ast.StructType) {

}
