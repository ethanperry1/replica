package parser

import "testing"

func TestMockFileCreator(t *testing.T) {
	mock := &MockDescender{}

	New(InitDescender(mock))
}