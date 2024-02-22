package parser

import "io"

func NewRenderer() (*Templater[MockFile], error) {
	executor, err := NewExecutor(Template)
	if err != nil {
		return nil, err
	}

	return NewTemplater[MockFile](executor), nil
}

type Writer struct {
	renderer TemplateRenderer[MockFile]
}

func NewWriter(renderer TemplateRenderer[MockFile]) *Writer {
	writer := &Writer{
		renderer: renderer,
	}
	return writer
}

func (writer *Writer) Generate(values *MockFile) (io.Reader, error) {
	return writer.renderer.Render(*values)
}