package gaze

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCallback[T any] struct {
	t *testing.T
}

func (cb *testCallback[T]) OnChange(oldValue T, newValue T) {
	cb.t.Logf(">> OnChange: %v -> %v", oldValue, newValue)
}

func newTestCallback[T any](t *testing.T) CallbackFuncs[T] {
	return &testCallback[T]{t}
}

func TestReactiveValue_WithoutCallback(t *testing.T) {
	t.Run("Boolean", func(t *testing.T) {
		// Create a new ReactiveValue with a boolean value
		ov := NewNopReactiveValue(true)

		// Get the value from the ReactiveValue
		got := ov.Get()
		assert.Equal(t, true, got, "Get() should return the true boolean")

		// Set a new value to the ReactiveValue and get it
		ov.Set(false)
		got = ov.Get()
		assert.Equal(t, false, got, "Get() should return the false boolean")
	})

	t.Run("String", func(t *testing.T) {
		// Create a new ReactiveValue with a string value
		ov := NewNopReactiveValue("hello")

		// Get the value from the ReactiveValue
		got := ov.Get()
		assert.Equal(t, "hello", got, "Get() should return the hello string")

		// Set a new value to the ReactiveValue and get it
		ov.Set("world")
		got = ov.Get()
		assert.Equal(t, "world", got, "Get() should return the world string")
	})

	t.Run("Int", func(t *testing.T) {
		// Create a new ReactiveValue with an int value
		ov := NewNopReactiveValue(42)

		// Get the value from the ReactiveValue
		got := ov.Get()
		assert.Equal(t, 42, got, "Get() should return the 42 int")

		// Set a new value to the ReactiveValue and get it
		ov.Set(100)
		got = ov.Get()
		assert.Equal(t, 100, got, "Get() should return the 100 int")
	})

	t.Run("Float", func(t *testing.T) {
		// Create a new ReactiveValue with a float value
		ov := NewNopReactiveValue(3.14)

		// Get the value from the ReactiveValue
		got := ov.Get()
		assert.Equal(t, 3.14, got, "Get() should return the 3.14 float")

		// Set a new value to the ReactiveValue and get it
		ov.Set(2.71)
		got = ov.Get()
		assert.Equal(t, 2.71, got, "Get() should return the 2.71 float")
	})

	t.Run("Complex", func(t *testing.T) {
		// Create a new ReactiveValue with a complex value
		ov := NewNopReactiveValue(1 + 2i)

		// Get the value from the ReactiveValue
		got := ov.Get()
		assert.Equal(t, 1+2i, got, "Get() should return the 1+2i complex")

		// Set a new value to the ReactiveValue and get it
		ov.Set(3 + 4i)
		got = ov.Get()
		assert.Equal(t, 3+4i, got, "Get() should return the 3+4i complex")
	})

	t.Run("Struct", func(t *testing.T) {
		// Create a new ReactiveValue with a struct value
		ov := NewNopReactiveValue(struct{ Name string }{Name: "John"})

		// Get the value from the ReactiveValue
		got := ov.Get()
		assert.Equal(t, struct{ Name string }{Name: "John"}, got, "Get() should return the struct")

		// Set a new value to the ReactiveValue and get it
		ov.Set(struct{ Name string }{Name: "Doe"})
		got = ov.Get()
		assert.Equal(t, struct{ Name string }{Name: "Doe"}, got, "Get() should return the struct")
	})

	t.Run("Slice", func(t *testing.T) {
		// Create a new ReactiveValue with a slice value
		ov := NewNopReactiveValue([]int{1, 2, 3})

		// Get the value from the ReactiveValue
		got := ov.Get()
		assert.Equal(t, []int{1, 2, 3}, got, "Get() should return the slice")

		// Set a new value to the ReactiveValue and get it
		ov.Set([]int{4, 5, 6})
		got = ov.Get()
		assert.Equal(t, []int{4, 5, 6}, got, "Get() should return the slice")
	})

	t.Run("Map", func(t *testing.T) {
		// Create a new ReactiveValue with a map value
		ov := NewNopReactiveValue(map[string]int{"one": 1, "two": 2})

		// Get the value from the ReactiveValue
		got := ov.Get()
		assert.Equal(t, map[string]int{"one": 1, "two": 2}, got, "Get() should return the map")

		// Set a new value to the ReactiveValue and get it
		ov.Set(map[string]int{"three": 3, "four": 4})
		got = ov.Get()
		assert.Equal(t, map[string]int{"three": 3, "four": 4}, got, "Get() should return the map")
	})

	t.Run("Pointer", func(t *testing.T) {
		// Create a new ReactiveValue with a pointer value
		n1 := 1
		ov := NewNopReactiveValue(&n1)

		// Get the value from the ReactiveValue
		got1 := ov.Get()
		assert.NotNil(t, got1, "Get() should return the pointer")

		// Set a same value with different ptr to the ReactiveValue and get it
		n2 := 1
		ov.Set(&n2)
		got2 := ov.Get()
		assert.NotNil(t, got2, "Get() should return the pointer")

		// Verify the pointers are different but the values are equal
		assert.Equal(t, got1, got2, "The pointer values should be equal")

		// Set a new value to the ReactiveValue and get it
		n3 := 2
		ov.Set(&n3)
		got1 = ov.Get()
		assert.NotNil(t, got1, "Get() should return the pointer")
	})
}

func TestReactiveValue_Callback(t *testing.T) {
	t.Run("Boolean", func(t *testing.T) {
		// Create a new ReactiveValue with a boolean value
		ov := NewReactiveValue(true, newTestCallback[bool](t))

		// Get the value from the ReactiveValue
		got := ov.Get()
		assert.Equal(t, true, got, "Get() should return the true boolean")

		// Set a new value to the ReactiveValue and get it
		ov.Set(false)
		got = ov.Get()
		assert.Equal(t, false, got, "Get() should return the false boolean")
	})

	t.Run("String", func(t *testing.T) {
		// Create a new ReactiveValue with a string value
		ov := NewReactiveValue("hello", newTestCallback[string](t))

		// Get the value from the ReactiveValue
		got := ov.Get()
		assert.Equal(t, "hello", got, "Get() should return the hello string")

		// Set a new value to the ReactiveValue and get it
		ov.Set("world")
		got = ov.Get()
		assert.Equal(t, "world", got, "Get() should return the world string")
	})

	t.Run("Int", func(t *testing.T) {
		// Create a new ReactiveValue with an int value
		ov := NewReactiveValue(42, newTestCallback[int](t))

		// Get the value from the ReactiveValue
		got := ov.Get()
		assert.Equal(t, 42, got, "Get() should return the 42 int")

		// Set a new value to the ReactiveValue and get it
		ov.Set(100)
		got = ov.Get()
		assert.Equal(t, 100, got, "Get() should return the 100 int")
	})

	t.Run("Float", func(t *testing.T) {
		// Create a new ReactiveValue with a float value
		ov := NewReactiveValue(3.14, newTestCallback[float64](t))

		// Get the value from the ReactiveValue
		got := ov.Get()
		assert.Equal(t, 3.14, got, "Get() should return the 3.14 float")

		// Set a new value to the ReactiveValue and get it
		ov.Set(2.71)
		got = ov.Get()
		assert.Equal(t, 2.71, got, "Get() should return the 2.71 float")
	})

	t.Run("Complex", func(t *testing.T) {
		// Create a new ReactiveValue with a complex value
		ov := NewReactiveValue(1+2i, newTestCallback[complex128](t))

		// Get the value from the ReactiveValue
		got := ov.Get()
		assert.Equal(t, 1+2i, got, "Get() should return the 1+2i complex")

		// Set a new value to the ReactiveValue and get it
		ov.Set(3 + 4i)
		got = ov.Get()
		assert.Equal(t, 3+4i, got, "Get() should return the 3+4i complex")
	})

	t.Run("Struct", func(t *testing.T) {
		// Create a new ReactiveValue with a struct value
		ov := NewReactiveValue(struct{ Name string }{Name: "John"}, newTestCallback[struct{ Name string }](t))

		// Get the value from the ReactiveValue
		got := ov.Get()
		assert.Equal(t, struct{ Name string }{Name: "John"}, got, "Get() should return the struct")

		// Set a new value to the ReactiveValue and get it
		ov.Set(struct{ Name string }{Name: "Doe"})
		got = ov.Get()
		assert.Equal(t, struct{ Name string }{Name: "Doe"}, got, "Get() should return the struct")
	})

	t.Run("Slice", func(t *testing.T) {
		// Create a new ReactiveValue with a slice value
		ov := NewReactiveValue([]int{1, 2, 3}, newTestCallback[[]int](t))

		// Get the value from the ReactiveValue
		got := ov.Get()
		assert.Equal(t, []int{1, 2, 3}, got, "Get() should return the slice")

		// Set a new value to the ReactiveValue and get it
		ov.Set([]int{4, 5, 6})
		got = ov.Get()
		assert.Equal(t, []int{4, 5, 6}, got, "Get() should return the slice")
	})

	t.Run("Map", func(t *testing.T) {
		// Create a new ReactiveValue with a map value
		ov := NewReactiveValue(map[string]int{"one": 1, "two": 2}, newTestCallback[map[string]int](t))

		// Get the value from the ReactiveValue
		got := ov.Get()
		assert.Equal(t, map[string]int{"one": 1, "two": 2}, got, "Get() should return the map")

		// Set a new value to the ReactiveValue and get it
		ov.Set(map[string]int{"three": 3, "four": 4})
		got = ov.Get()
		assert.Equal(t, map[string]int{"three": 3, "four": 4}, got, "Get() should return the map")
	})

	t.Run("Pointer", func(t *testing.T) {
		// Create a new ReactiveValue with a pointer value
		n1 := 1
		ov := NewReactiveValue(&n1, newTestCallback[*int](t))

		// Get the value from the ReactiveValue
		got1 := ov.Get()
		assert.NotNil(t, got1, "Get() should return the pointer")

		// Set a same value with different ptr to the ReactiveValue and get it
		n2 := 1
		ov.Set(&n2)
		got2 := ov.Get()
		assert.NotNil(t, got2, "Get() should return the pointer")

		// Verify the pointers are different but the values are equal
		assert.Equal(t, got1, got2, "The pointer values should be equal")

		// Set a new value to the ReactiveValue and get it
		n3 := 2
		ov.Set(&n3)
		got1 = ov.Get()
		assert.NotNil(t, got1, "Get() should return the pointer")
	})
}
