//go:generate replica
package parser

import (
	ast "go/ast"
	"go/token"
	"strings"
)

type (
	MockFile struct {
		Mocks   []*Mock
		Imports []string
		Package string
	}

	MockIdents struct {
		Mocks  []*Mock
		Idents map[string]int
	}

	Mock struct {
		Name    string
		Methods []*Method
	}

	Method struct {
		Name     string
		Function *Function
	}

	Function struct {
		Types   []*Field
		Params  []*Field
		Returns []*Field
	}

	Field struct {
		Name string
		Type string
	}

	Import struct {
		Name  string
		Alias string
		Path  string
	}
)

// replica:gen
type Descender interface {
	ParseUsedImports(idents map[string]int, imports map[string]*Import) []string
	ParseImports(specs []*ast.ImportSpec) map[string]*Import
	ParseFile(file *ast.File, generateAll bool) *MockIdents
}

type (
	MockFileCreator struct {
		descender   Descender
		generateAll bool
	}

	RecursiveDescender struct {
		idents map[string]int
	}
)

func InitGenerateAll(generateAll bool) func(*MockFileCreator) {
	return func(mfc *MockFileCreator) {
		mfc.generateAll = generateAll
	}
}

func InitDescender(descender Descender) func(*MockFileCreator) {
	return func(mfc *MockFileCreator) {
		mfc.descender = descender
	}
}

func New(options ...func(*MockFileCreator)) *MockFileCreator {
	creator := &MockFileCreator{
		descender: NewRecursiveDescender(make(map[string]int)),
	}

	for _, option := range options {
		option(creator)
	}

	return creator
}

func (creator *MockFileCreator) CreateMockFile(file *ast.File) *MockFile {
	mockIdents := creator.descender.ParseFile(file, creator.generateAll)
	imports := creator.descender.ParseImports(file.Imports)
	used := creator.descender.ParseUsedImports(mockIdents.Idents, imports)

	return &MockFile{
		Mocks:   mockIdents.Mocks,
		Imports: used,
		Package: file.Name.Name,
	}
}

func NewRecursiveDescender(idents map[string]int) *RecursiveDescender {
	return &RecursiveDescender{
		idents: idents,
	}
}

func (descender *RecursiveDescender) ParseUsedImports(idents map[string]int, imports map[string]*Import) []string {
	var used []string
	for ident := range idents {
		imp, ok := imports[ident]
		if ok {
			used = append(used, imp.Alias+" "+imp.Path)
		}
	}

	return used
}

func (descender *RecursiveDescender) ParseImports(specs []*ast.ImportSpec) map[string]*Import {
	m := make(map[string]*Import)

	for _, spec := range specs {
		imp := descender.ParseImport(spec)
		m[imp.Name] = imp
	}

	return m
}

func (descender *RecursiveDescender) ParseImport(spec *ast.ImportSpec) *Import {
	if spec.Name != nil {
		return &Import{
			Name:  spec.Name.Name,
			Alias: spec.Name.Name,
			Path:  spec.Path.Value,
		}
	}

	trimmed := strings.Trim(spec.Path.Value, "\"")
	parts := strings.Split(trimmed, "/")
	return &Import{
		Name: parts[len(parts)-1],
		Path: spec.Path.Value,
	}
}

func (descender *RecursiveDescender) ParseFile(file *ast.File, generateAll bool) *MockIdents {
	var mocks []*Mock
	for _, decl := range file.Decls {
		switch d := decl.(type) {
		case *ast.GenDecl:
			mocks = append(mocks, descender.ParseGenDecl(d, generateAll)...)
		}
	}

	return &MockIdents{
		Idents: descender.idents,
		Mocks:  mocks,
	}
}

func (descender *RecursiveDescender) ParseGenDecl(decl *ast.GenDecl, generateAll bool) []*Mock {
	var mocks []*Mock
	gen := generateAll || CheckComments(decl.Doc)
	for _, spec := range decl.Specs {
		switch s := spec.(type) {
		case *ast.TypeSpec:
			res, ok := descender.ParseTypeSpec(s, gen)
			if ok {
				mocks = append(mocks, res)
			}
		}
	}

	return mocks
}

func (descender *RecursiveDescender) ParseTypeSpec(spec *ast.TypeSpec, generate bool) (*Mock, bool) {
	switch s := spec.Type.(type) {
	case *ast.InterfaceType:
		if generate || CheckComments(spec.Doc) {
			return &Mock{
				Methods: descender.ParseInterface(s),
				Name:    spec.Name.Name,
			}, true
		}
	}

	return nil, false
}

func (descender *RecursiveDescender) ParseInterface(typ *ast.InterfaceType) []*Method {
	var methods []*Method
	for _, method := range typ.Methods.List {
		var name string
		if len(method.Names) > 0 {
			name = method.Names[0].Name
		}
		switch m := method.Type.(type) {
		case *ast.FuncType:
			methods = append(methods, &Method{
				Name:     name,
				Function: descender.ParseFunction(m),
			})
		}
	}

	return methods
}

func (descender *RecursiveDescender) ParseFunction(fun *ast.FuncType) *Function {
	var types []*Field
	if fun.TypeParams != nil {
		for _, field := range fun.TypeParams.List {
			types = append(types, descender.ParseField(field)...)
		}
	}

	var params []*Field
	if fun.Params != nil {
		for _, param := range fun.Params.List {
			params = append(params, descender.ParseField(param)...)
		}
	}

	var returns []*Field
	if fun.Results != nil {
		for _, result := range fun.Results.List {
			returns = append(returns, descender.ParseField(result)...)
		}
	}

	return &Function{
		Types:   types,
		Params:  params,
		Returns: returns,
	}
}

func (descender *RecursiveDescender) ParseFuncType(fun *ast.FuncType) string {
	var typeParams string
	if fun.TypeParams != nil {
		var types []string
		for _, field := range fun.TypeParams.List {
			for _, typ := range descender.ParseField(field) {
				types = append(types, typ.Name+" "+typ.Type)
			}
		}
		typeParams = "[" + strings.Join(types, ", ") + "]"
	}

	var params []string
	if fun.Params != nil {
		for _, field := range fun.Params.List {
			for _, typ := range descender.ParseField(field) {
				params = append(params, typ.Name+" "+typ.Type)
			}
		}
	}

	var returns []string
	if fun.Results != nil {
		for _, field := range fun.Results.List {
			for _, typ := range descender.ParseField(field) {
				returns = append(returns, typ.Name+" "+typ.Type)
			}
		}
	}

	var fnc string
	if fun.Func != token.NoPos {
		fnc = "func"
	}

	return fnc + typeParams + "(" + strings.Join(params, ", ") + ")" + strings.Join(returns, ", ")
}

func (descender *RecursiveDescender) ParseArrayType(typ *ast.ArrayType) string {
	l := "[]"
	if typ.Len != nil {
		l = descender.ParseExpr(typ.Len)
	}
	return l + descender.ParseExpr(typ.Elt)
}

func (descender *RecursiveDescender) ParseField(f *ast.Field) []*Field {
	typ := descender.ParseExpr(f.Type)

	if len(f.Names) == 0 {
		return []*Field{
			{
				Type: typ,
			},
		}
	}

	fields := make([]*Field, len(f.Names))
	for idx, ident := range f.Names {
		fields[idx] = &Field{
			Type: typ,
			Name: descender.ParseIdent(ident),
		}
	}
	return fields
}

func (descender *RecursiveDescender) ParseIdent(i *ast.Ident) string {
	// Add ident to idents map.
	descender.idents[i.Name]++
	return i.Name
}

func (descender *RecursiveDescender) ParseExpr(expr ast.Expr) string {
	switch e := expr.(type) {
	case *ast.IndexExpr:
		return descender.ParseIndexExpr(e)
	case *ast.IndexListExpr:
		return descender.ParseIndexListExpr(e)
	case *ast.SelectorExpr:
		return descender.ParseSelectorExpr(e)
	case *ast.Ident:
		return descender.ParseIdent(e)
	case *ast.ArrayType:
		return descender.ParseArrayType(e)
	case *ast.FuncType:
		return descender.ParseFuncType(e)
	case *ast.StructType:
		return descender.ParseStructType(e)
	case *ast.Ellipsis:
		return descender.ParseEllipsis(e)
	case *ast.InterfaceType:
		return descender.ParseInterfaceType(e)
	case *ast.StarExpr:
		return descender.ParseStarExpression(e)
	case *ast.MapType:
		return descender.ParseMapType(e)
	case *ast.ChanType:
		return descender.ParseChanType(e)
	}

	return ""
}

func (descender *RecursiveDescender) ParseIndexExpr(expr *ast.IndexExpr) string {
	return descender.ParseExpr(expr.X) + "[" + descender.ParseExpr(expr.Index) + "]"
}

func (descender *RecursiveDescender) ParseIndexListExpr(expr *ast.IndexListExpr) string {

	x := descender.ParseExpr(expr.X)

	indices := make([]string, len(expr.Indices))
	for idx, index := range expr.Indices {
		indices[idx] = descender.ParseExpr(index)
	}

	return x + "[" + strings.Join(indices, ", ") + "]"
}

func (descender *RecursiveDescender) ParseSelectorExpr(expr *ast.SelectorExpr) string {
	return descender.ParseExpr(expr.X) + "." + descender.ParseIdent(expr.Sel)
}

func (descender *RecursiveDescender) ParseStructType(s *ast.StructType) string {
	var fields []string
	for _, field := range s.Fields.List {
		res := descender.ParseField(field)
		for _, field := range res {
			fields = append(fields, field.Name+" "+field.Type)
		}
	}
	return "struct {" + strings.Join(fields, ",") + "}"
}

func (descender *RecursiveDescender) ParseEllipsis(e *ast.Ellipsis) string {
	return "..." + descender.ParseExpr(e.Elt)
}

func (descender *RecursiveDescender) ParseInterfaceType(typ *ast.InterfaceType) string {
	var methods []string
	for _, method := range typ.Methods.List {
		res := descender.ParseField(method)
		for _, field := range res {
			methods = append(methods, field.Name+" "+field.Type)
		}
	}
	return "interface {" + strings.Join(methods, ",") + "}"
}

func (descender *RecursiveDescender) ParseStarExpression(expr *ast.StarExpr) string {
	return "*" + descender.ParseExpr(expr.X)
}

func (descender *RecursiveDescender) ParseMapType(typ *ast.MapType) string {
	return "map[" + descender.ParseExpr(typ.Key) + "]" + descender.ParseExpr(typ.Value)
}

func (descender *RecursiveDescender) ParseChanType(typ *ast.ChanType) string {
	res := descender.ParseExpr(typ.Value)
	if typ.Arrow != token.NoPos {
		if typ.Dir == ast.SEND {
			return "chan <- " + res
		}
		return "<-chan " + res
	}
	return "chan " + res
}
