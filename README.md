# Replica 💎

Replica is a golang code generation tool for building mock implementations of interfaces.

## Usage

```sh
# Must add directive to file of choice: //go:generate replica
# And add the appropriate comment to the interface of choice: //replica:gen
go install github.com/ethanperry1/replica@latest
go generate ./...
```

## Example

### Example File w/ Go Generate Statement

```go
//go:generate replica
package main

import (
	"os"
)

// replica:gen
type (
	Example[T any] interface {
		Method(T) (any, os.File)
	}
)
```

### Automatically Generated Mock

```go
// This is an automatically generated file! Do not modify.
package main
 
import( 
	 "os" 
)

type (
	// MockExample is an automatically generated function mocking the Example interface.
	MockExample[T any,] struct { 
		OnMethod func( 
			t T,
		) ( 
			any,
			os.File,
		) 
	} 
)

// Method is an automatically generated function used for mocking.
func (mock *MockExample[T,]) Method(
	t T,
) (
	any, 
	os.File, 
) {  
	return mock.OnMethod(
		t,
	)
}
```