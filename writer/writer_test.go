package writer

import (
	"io"
	"testing"

	"github.com/stretchr/testify/require"
)

const (
	exampleImport = "example/example"
)

func TestMain(m *testing.M) {
	m.Run()
}

func TestWriter(t *testing.T) {
	renderer, err := NewRenderer()
	require.NoError(t, err)

	writer := New(renderer)
	res, err := writer.Generate(&Values{
		Imports: []string{
			exampleImport,
		},
	})
	require.NoError(t, err)
	content, err := io.ReadAll(res)
	require.NoError(t, err)

	require.Equal(t, "", string(content))
}