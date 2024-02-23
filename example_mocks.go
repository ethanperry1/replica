// This is an automatically generated file! Do not modify.
package main
 
import( 
	 "os" 
)

type (
	// MockExampleOne is an automatically generated function mocking the ExampleOne interface.
	MockExampleOne struct { 
		OnMethodOne func( 
		) ( 
		) 
		OnMethodTwo func( 
			f func(),
		) ( 
			ExampleTwo[ExampleOne],
		) 
	} 
	// MockExampleTwo is an automatically generated function mocking the ExampleTwo interface.
	MockExampleTwo[T any,] struct { 
		OnMethodThree func( 
			t T,
		) ( 
			any,
		) 
	} 
	// MockExampleThree is an automatically generated function mocking the ExampleThree interface.
	MockExampleThree[T ExampleTwo[S],S any,] struct { 
		OnMethodFour func( 
			t T,
		) ( 
			func() map[os.File]struct {
a S
b interface {
M () string}},
		) 
	} 
	// MockExampleFour is an automatically generated function mocking the ExampleFour interface.
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

// MethodOne is an automatically generated function used for mocking.
func (mock *MockExampleOne) MethodOne(
) { mock.OnMethodOne(
	)
}

// MethodTwo is an automatically generated function used for mocking.
func (mock *MockExampleOne) MethodTwo(
	f func(),
) (
	ExampleTwo[ExampleOne], 
) {  
	return mock.OnMethodTwo(
		f,
	)
}

// MethodThree is an automatically generated function used for mocking.
func (mock *MockExampleTwo[T,]) MethodThree(
	t T,
) (
	any, 
) {  
	return mock.OnMethodThree(
		t,
	)
}

// MethodFour is an automatically generated function used for mocking.
func (mock *MockExampleThree[T,S,]) MethodFour(
	t T,
) (
	func() map[os.File]struct {
a S
b interface {
M () string}}, 
) {  
	return mock.OnMethodFour(
		t,
	)
}

// MethodFive is an automatically generated function used for mocking.
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

