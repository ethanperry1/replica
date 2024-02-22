package example

type Example[T Other[S], S any] struct{}

type Other[T any] struct{}