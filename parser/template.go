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
	Mock{{ .Name }} struct { {{ range .Methods }}
		On{{ .Name }} func( {{ range .Function.Params }}
			{{ .Name }} {{ .Type }},{{ end }}
		) ( {{ range .Function.Returns }}
			{{ .Type }},{{ end }}
		) {{ end }}
	} {{ end }}
)
{{ end }}{{ range $mock := .Mocks }}{{ range .Methods }}
func (mock *Mock{{ $mock.Name }}) {{ .Name }}({{ range .Function.Params }}
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