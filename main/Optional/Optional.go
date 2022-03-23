package main

type Optional[T any] struct {
	b   bool
	val T
}

// Constructors
func NewOptional[T any]() Optional[T] {
	return Optional[T]{
		b: false,
	}
}

func NewOptionalV[T any](val T) Optional[T] {
	return Optional[T]{
		b:   false,
		val: val,
	}
}

// Functions
func (o Optional[T]) IsJust() bool {
	return o.b
}

func (o Optional[T]) IsNothing() bool {
	return !o.b
}

func (o Optional[T]) FromJust() T {
	return o.val
}

// Shorts
func Nothing[T any]() Optional[T] {
	return NewOptional[T]()
}

func Just[T any](v T) Optional[T] {
	return NewOptionalV(v)
}
