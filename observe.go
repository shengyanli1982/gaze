package gaze

import "reflect"

// CallbackFuncs 是一个泛型接口，它定义了一个 OnChange 方法，该方法在值发生更改时调用。
// CallbackFuncs is a generic interface that defines an OnChange method, which is called when a value changes.
type CallbackFuncs[T any] interface {
	// OnChange 是当值发生更改时调用的方法。
	// OnChange is the method called when a value changes.
	OnChange(oldValue T, newValue T)
}

// nopCallbackImpl 是 CallbackFuncs 接口的一个实现，但所有方法都没有实际操作。
// nopCallbackImpl is an implementation of the CallbackFuncs interface, but all methods do nothing.
type nopCallbackImpl[T any] struct{}

// OnChange 是 nopCallbackImpl 的 OnChange 方法实现，但没有实际操作。
// OnChange is the implementation of the OnChange method for nopCallbackImpl, but it does nothing.
func (cb *nopCallbackImpl[T]) OnChange(oldValue T, newValue T) {}

// NewNopCallbackFuncs 是一个工厂函数，用于创建一个新的 nopCallbackImpl 实例。
// NewNopCallbackFuncs is a factory function for creating a new instance of nopCallbackImpl.
func NewNopCallbackFuncs[T any]() CallbackFuncs[T] {
	// 返回一个新的 nopCallbackImpl 实例。
	// Returns a new instance of nopCallbackImpl.
	return &nopCallbackImpl[T]{}
}

// ReactiveValue 是一个泛型结构体，它包含一个值和一个回调函数。
// ReactiveValue is a generic struct that contains a value and a callback function.
type ReactiveValue[T any] struct {
	// value 是存储在 ReactiveValue 中的值。
	// value is the value stored in the ReactiveValue.
	value T

	// callbacks 是当值发生变化时调用的回调函数。
	// callbacks is the callback function that is called when the value changes.
	callbacks CallbackFuncs[T]
}

// NewReactiveValue 是一个工厂函数，用于创建一个新的 ReactiveValue 实例。
// NewReactiveValue is a factory function for creating a new instance of ReactiveValue.
func NewReactiveValue[T any](initialValue T, callbacks CallbackFuncs[T]) *ReactiveValue[T] {
	// 如果没有提供回调函数，我们将使用一个无操作的回调函数。
	// If no callback function is provided, we will use a no-operation callback function.
	if callbacks == nil {
		callbacks = NewNopCallbackFuncs[T]()
	}

	// 返回一个新的 ReactiveValue 实例，其中包含提供的值和回调函数。
	// Returns a new instance of ReactiveValue that contains the provided value and callback function.
	return &ReactiveValue[T]{
		value:     initialValue,
		callbacks: callbacks,
	}
}

// NewNopReactiveValue 是一个工厂函数，用于创建一个新的 ReactiveValue 实例，但不提供回调函数。
// NewNopReactiveValue is a factory function for creating a new instance of ReactiveValue, but without providing a callback function.
func NewNopReactiveValue[T any](initialValue T) *ReactiveValue[T] {
	// 使用 nil 作为回调函数调用 NewReactiveValue。
	// Calls NewReactiveValue using nil as the callback function.
	return NewReactiveValue[T](initialValue, nil)
}

// Get 是 ReactiveValue 的一个方法，用于获取存储的值。
// Get is a method of ReactiveValue for getting the stored value.
func (rv *ReactiveValue[T]) Get() T { return rv.value }

// Set 是 ReactiveValue 的一个方法，用于设置存储的值。
// Set is a method of ReactiveValue for setting the stored value.
func (rv *ReactiveValue[T]) Set(newValue T) {
	// 我们首先检查新值是否与旧值相等。
	// We first check if the new value is equal to the old value.
	if reflect.DeepEqual(rv.value, newValue) {
		// 如果新值与旧值相等，我们只需将新值设置为旧值。 此处主要是解决 map, slice, struct 等引用类型的问题 (值相等，但是对象不同)。
		// If the new value is equal to the old value, we just set the new value to the old value.
		// This is mainly to solve the problem of reference types such as map, slice, struct (the values are equal, but the objects are different).
		rv.value = newValue
	} else {
		// 如果新值与旧值不等，我们首先保存旧值。
		// If the new value is not equal to the old value, we first save the old value.
		oldValue := rv.value

		// 然后我们设置新值。
		// Then we set the new value.
		rv.value = newValue

		// 最后，我们调用回调函数的 OnChange 方法，传入旧值和新值。
		// Finally, we call the OnChange method of the callback function, passing in the old value and the new value.
		rv.callbacks.OnChange(oldValue, newValue)
	}
}
