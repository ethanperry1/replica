package parser

import (
	"io"
	"testing"

	"github.com/stretchr/testify/require"
)

const (
	exampleImport = "example/example"
)

func TestWriter(t *testing.T) {
	renderer, err := NewRenderer()
	require.NoError(t, err)

	writer := NewWriter(renderer)
	res, err := writer.Generate(&MockFile{
		Imports: []string{
			exampleImport,
		}, Mocks: []*Mock{
			{
				Name: "Example",
				Methods: []*Method{
					{
						Name: "Example",
						Function: &Function{
							Params: []*Field{
								{
									Name: "example",
									Type: "int",
								},
							},
							Returns: []*Field{
								{
									Type: "error",
								},
							},
						},
					},
				},
			},
		},
	})
	require.NoError(t, err)
	content, err := io.ReadAll(res)
	require.NoError(t, err)

	require.Equal(t, "", string(content))
}
