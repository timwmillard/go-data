package data

import "reflect"

func Zero[T any]() T {
	var result T
	return result
}

func FromPtr[T any](s *T) T {
	if s == nil {
		return Zero[T]()
	}
	return *s
}

func ToPtr[T any](s T) *T {
	if IsZero(s) {
		return nil
	}
	return &s
}

func IsZero[T any](v T) bool {
	return reflect.ValueOf(&v).Elem().IsZero()
}
