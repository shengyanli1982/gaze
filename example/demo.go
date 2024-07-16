package main

import (
	"fmt"

	"github.com/shengyanli1982/gaze"
)

type demoCallback[T any] struct{}

func (cb *demoCallback[T]) OnChange(oldValue T, newValue T) {
	fmt.Printf(">> OnChange: %v -> %v\n", oldValue, newValue)
}

func newTestCallback[T any]() gaze.CallbackFuncs[T] {
	return &demoCallback[T]{}
}

func main() {
	// Create a new ReactiveValue with a int value
	ov := gaze.NewReactiveValue(11, newTestCallback[int]())

	// Get the value from the ReactiveValue
	got := ov.Get()
	fmt.Printf("got: %v\n", got)

	// Set a new value to the ReactiveValue and get it
	ov.Set(3)
	got = ov.Get()
	fmt.Printf("got: %v\n", got)
}
