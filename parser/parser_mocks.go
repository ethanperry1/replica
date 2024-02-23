// This is an automatically generated file! Do not modify.
package parser
 
import( 
	ast "go/ast" 
)

type (
	// MockDescender is an automatically generated function mocking the Descender interface.
	MockDescender struct { 
		OnParseUsedImports func( 
			idents map[string]int,
			imports map[string]*Import,
		) ( 
			[]string,
		) 
		OnParseImports func( 
			specs []*ast.ImportSpec,
		) ( 
			map[string]*Import,
		) 
		OnParseFile func( 
			file *ast.File,
			generateAll bool,
		) ( 
			*MockIdents,
		) 
	} 
)

// ParseUsedImports is an automatically generated function used for mocking.
func (mock *MockDescender) ParseUsedImports(
	idents map[string]int,
	imports map[string]*Import,
) (
	[]string, 
) {  
	return mock.OnParseUsedImports(
		idents,
		imports,
	)
}

// ParseImports is an automatically generated function used for mocking.
func (mock *MockDescender) ParseImports(
	specs []*ast.ImportSpec,
) (
	map[string]*Import, 
) {  
	return mock.OnParseImports(
		specs,
	)
}

// ParseFile is an automatically generated function used for mocking.
func (mock *MockDescender) ParseFile(
	file *ast.File,
	generateAll bool,
) (
	*MockIdents, 
) {  
	return mock.OnParseFile(
		file,
		generateAll,
	)
}

