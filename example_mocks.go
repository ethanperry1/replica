// This is an automatically generated file! Do not modify.
package main

import (
	"os"
)

type (
	MockExampleOne struct {
		OnMethodOne func()

		OnMethodTwo func(
			f func(),
		) ExampleTwo[ExampleOne]
	}
	MockExampleTwo[T any] struct {
		OnMethodThree func(
			t T,
		) any
	}
	MockExampleThree[T ExampleTwo[S], S any] struct {
		OnMethodFour func(
			t T,
		) func() map[os.File]struct {
			a S
			b interface {
				M() string
			}
		}
	}
	MockExampleFour struct {
		OnMethodFive func(
			a A,
			a2 A,
			a3 A,
			b B,
			b2 B,
		) (
			A,
			B,
		)
	}
)

func (mock *MockExampleOne) MethodOne() {
	mock.OnMethodOne()
}

func (mock *MockExampleOne) MethodTwo(
	f func(),
) ExampleTwo[ExampleOne] {
	return mock.OnMethodTwo(
		f,
	)
}

func (mock *MockExampleTwo[T]) MethodThree(
	t T,
) any {
	return mock.OnMethodThree(
		t,
	)
}

func (mock *MockExampleThree[T, S]) MethodFour(
	t T,
) func() map[os.File]struct {
	a S
	b interface {
		M() string
	}
} {
	return mock.OnMethodFour(
		t,
	)
}

func (mock *MockExampleFour) MethodFive(
	a A,
	a2 A,
	a3 A,
	b B,
	b2 B,
) (
	A,
	B,
) {
	return mock.OnMethodFive(
		a,
		a2,
		a3,
		b,
		b2,
	)
}
