package Optional

type Optional[T any] interface {
	IsJust() bool
	IsNothing() bool
	FromJust() T
}

type OptionalImpl[T any] struct {
	b   bool
	val T
}

// Constructors
func NewOptional[T any]() OptionalImpl[T] {
	return OptionalImpl[T]{
		b: false,
	}
}

func NewOptionalV[T any](val T) OptionalImpl[T] {
	return OptionalImpl[T]{
		b:   true,
		val: val,
	}
}

// Functions
func (o OptionalImpl[T]) IsJust() bool {
	return o.b
}

func (o OptionalImpl[T]) IsNothing() bool {
	return !o.b
}

func (o OptionalImpl[T]) FromJust() T {
	return o.val
}

// Shorts
func Nothing[T any]() OptionalImpl[T] {
	return NewOptional[T]()
}

func Just[T any](v T) OptionalImpl[T] {
	return NewOptionalV(v)
}
