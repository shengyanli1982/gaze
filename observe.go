package gaze

import "reflect"

type CallbackFuncs[T any] interface {
	OnSet(value T)
	OnGet(value T)
	OnChange(oldValue T, newValue T)
}

type nopCallbackImpl[T any] struct{}

func (cb *nopCallbackImpl[T]) OnSet(value T) {}

func (cb *nopCallbackImpl[T]) OnGet(value T) {}

func (cb *nopCallbackImpl[T]) OnChange(oldValue T, newValue T) {}

func NewNopCallback[T any]() CallbackFuncs[T] {
	return &nopCallbackImpl[T]{}
}

type ObservableValue[T any] struct {
	value T
	cb    CallbackFuncs[T]
}

func NewObservableValue[T any](value T, cb CallbackFuncs[T]) *ObservableValue[T] {
	if cb == nil {
		cb = NewNopCallback[T]()
	}

	return &ObservableValue[T]{
		value: value,
		cb:    cb,
	}
}

func (o *ObservableValue[T]) Get() T {
	o.cb.OnGet(o.value)
	return o.value
}

func (o *ObservableValue[T]) Set(value T) {
	if reflect.DeepEqual(o.value, value) {
		o.cb.OnSet(value)
	} else {
		oldValue := o.value
		o.value = value
		o.cb.OnChange(oldValue, value)
	}
}

func NewNopObservableValue[T any](value T) *ObservableValue[T] {
	return NewObservableValue[T](value, nil)
}
