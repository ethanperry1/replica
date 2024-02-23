//go:generate replica
package main

import (
	"os"
)

// replica:gen
type (
	ExampleOne interface {
		MethodOne()
		MethodTwo(func()) ExampleTwo[ExampleOne]
	}
	ExampleTwo[T any] interface {
		MethodThree(T) any
	}
)

// replica:gen
type ExampleThree[T ExampleTwo[S], S any] interface {
	MethodFour(T) func() map[os.File]struct{
		a S
		b interface{
			M() string
		}
	}
}

// replica:gen
type (
	A struct{}
	B struct{}
	ExampleFour interface {
		MethodFive(A, A, A, B, B) (A, B)
	}
)