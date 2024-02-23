package parser

const (
	Template = 
`// This is an automatically generated file! Do not modify.
package {{ .Package }}
{{ if .Imports }} 
import( {{ range .Imports }}
	{{ . }} {{ end }}
)
{{ end }}{{ if .Mocks }}
type ({{ range .Mocks }}
	// Mock{{ .Name }} is an automatically generated function mocking the {{ .Name }} interface.
	Mock{{ .Name }}{{ if .Types }}[{{ range .Types }}{{.Name}} {{.Type}},{{ end }}]{{ end }} struct { {{ range .Methods }}
		On{{ .Name }} func( {{ range .Function.Params }}
			{{ .Name }} {{ .Type }},{{ end }}
		) ( {{ range .Function.Returns }}
			{{ .Type }},{{ end }}
		) {{ end }}
	} {{ end }}
)
{{ end }}{{ range $mock := .Mocks }}{{ range .Methods }}
// {{ .Name }} is an automatically generated function used for mocking.
func (mock *Mock{{ $mock.Name }}{{ if $mock.Types }}[{{ range $mock.Types }}{{.Name}},{{ end }}]{{ end }}) {{ .Name }}({{ range .Function.Params }}
	{{ .Name }} {{ .Type }},{{ end }}
){{ if .Function.Returns }} ({{ range .Function.Returns }}
	{{ .Type }}, {{end}}
){{ end }} { {{ if .Function.Returns }} 
	return {{ end }}mock.On{{.Name}}({{ range .Function.Params }}
		{{.Name}},{{ end }}
	)
}
{{ end }}{{ end }}
`
)