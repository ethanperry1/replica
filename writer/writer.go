package writer

import (
	"io"
	"replica/templater"
)

const (
	Template = `
{{ if .Imports }} 
import(
	{{ range .Imports }}
	"{{ . }}"
	{{ end }}
)
{{ end }}

{{ if .Mocks }}
type (
	{{ range .Mocks }}
	Mock{{ .Name }} struct {
		{{ range .Functions }}

		{{ end }}
	}
	{{ end }}
)
{{ end }}
	`
)

type Values struct {
	Imports []string
	Mocks   []Mock
}

type Mock struct {
	Name string
	Functions Function
}

type Function struct {

}

func NewRenderer() (*templater.Templater[Values], error) {
	executor, err := templater.NewExecutor(Template)
	if err != nil {
		return nil, err
	}

	return templater.New[Values](executor), nil
}

type Writer struct {
	renderer templater.TemplateRenderer[Values]
}

func New(renderer templater.TemplateRenderer[Values]) *Writer {
	writer := &Writer{
		renderer: renderer,
	}
	return writer
}

func (writer *Writer) Generate(values *Values) (io.Reader, error) {
	return writer.renderer.Render(*values)
}
