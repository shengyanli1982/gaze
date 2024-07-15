package gaze

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCallback[T any] struct {
	t *testing.T
}

func (cb *testCallback[T]) OnSet(value T) {
	cb.t.Logf(">> OnSet: %v", value)
}

func (cb *testCallback[T]) OnGet(value T) {
	cb.t.Logf(">> OnGet: %v", value)
}

func (cb *testCallback[T]) OnChange(oldValue T, newValue T) {
	cb.t.Logf(">> OnChange: %v -> %v", oldValue, newValue)
}

func newTestCallback[T any](t *testing.T) CallbackFuncs[T] {
	return &testCallback[T]{t}
}

func TestObservableValue_WithoutCallback(t *testing.T) {
	t.Run("Boolean", func(t *testing.T) {
		// Create a new ObservableValue with a boolean value
		ov := NewNopObservableValue(true)

		// Get the value from the ObservableValue
		got := ov.Get()
		assert.Equal(t, true, got, "Get() should return the true boolean")

		// Set a new value to the ObservableValue and get it
		ov.Set(false)
		got = ov.Get()
		assert.Equal(t, false, got, "Get() should return the false boolean")
	})

	t.Run("String", func(t *testing.T) {
		// Create a new ObservableValue with a string value
		ov := NewNopObservableValue("hello")

		// Get the value from the ObservableValue
		got := ov.Get()
		assert.Equal(t, "hello", got, "Get() should return the hello string")

		// Set a new value to the ObservableValue and get it
		ov.Set("world")
		got = ov.Get()
		assert.Equal(t, "world", got, "Get() should return the world string")
	})

	t.Run("Int", func(t *testing.T) {
		// Create a new ObservableValue with an int value
		ov := NewNopObservableValue(42)

		// Get the value from the ObservableValue
		got := ov.Get()
		assert.Equal(t, 42, got, "Get() should return the 42 int")

		// Set a new value to the ObservableValue and get it
		ov.Set(100)
		got = ov.Get()
		assert.Equal(t, 100, got, "Get() should return the 100 int")
	})

	t.Run("Float", func(t *testing.T) {
		// Create a new ObservableValue with a float value
		ov := NewNopObservableValue(3.14)

		// Get the value from the ObservableValue
		got := ov.Get()
		assert.Equal(t, 3.14, got, "Get() should return the 3.14 float")

		// Set a new value to the ObservableValue and get it
		ov.Set(2.71)
		got = ov.Get()
		assert.Equal(t, 2.71, got, "Get() should return the 2.71 float")
	})

	t.Run("Complex", func(t *testing.T) {
		// Create a new ObservableValue with a complex value
		ov := NewNopObservableValue(1 + 2i)

		// Get the value from the ObservableValue
		got := ov.Get()
		assert.Equal(t, 1+2i, got, "Get() should return the 1+2i complex")

		// Set a new value to the ObservableValue and get it
		ov.Set(3 + 4i)
		got = ov.Get()
		assert.Equal(t, 3+4i, got, "Get() should return the 3+4i complex")
	})

	t.Run("Struct", func(t *testing.T) {
		// Create a new ObservableValue with a struct value
		ov := NewNopObservableValue(struct{ Name string }{Name: "John"})

		// Get the value from the ObservableValue
		got := ov.Get()
		assert.Equal(t, struct{ Name string }{Name: "John"}, got, "Get() should return the struct")

		// Set a new value to the ObservableValue and get it
		ov.Set(struct{ Name string }{Name: "Doe"})
		got = ov.Get()
		assert.Equal(t, struct{ Name string }{Name: "Doe"}, got, "Get() should return the struct")
	})

	t.Run("Slice", func(t *testing.T) {
		// Create a new ObservableValue with a slice value
		ov := NewNopObservableValue([]int{1, 2, 3})

		// Get the value from the ObservableValue
		got := ov.Get()
		assert.Equal(t, []int{1, 2, 3}, got, "Get() should return the slice")

		// Set a new value to the ObservableValue and get it
		ov.Set([]int{4, 5, 6})
		got = ov.Get()
		assert.Equal(t, []int{4, 5, 6}, got, "Get() should return the slice")
	})

	t.Run("Map", func(t *testing.T) {
		// Create a new ObservableValue with a map value
		ov := NewNopObservableValue(map[string]int{"one": 1, "two": 2})

		// Get the value from the ObservableValue
		got := ov.Get()
		assert.Equal(t, map[string]int{"one": 1, "two": 2}, got, "Get() should return the map")

		// Set a new value to the ObservableValue and get it
		ov.Set(map[string]int{"three": 3, "four": 4})
		got = ov.Get()
		assert.Equal(t, map[string]int{"three": 3, "four": 4}, got, "Get() should return the map")
	})

	t.Run("Pointer", func(t *testing.T) {
		// Create a new ObservableValue with a pointer value
		n1 := 1
		ov := NewNopObservableValue(&n1)

		// Get the value from the ObservableValue
		got := ov.Get()
		assert.NotNil(t, got, "Get() should return the pointer")

		// Set a new value to the ObservableValue and get it
		n2 := 2
		ov.Set(&n2)
		got = ov.Get()
		assert.NotNil(t, got, "Get() should return the pointer")
	})
}

func TestObservableValue_Callback(t *testing.T) {
	t.Run("Boolean", func(t *testing.T) {
		// Create a new ObservableValue with a boolean value
		ov := NewObservableValue(true, newTestCallback[bool](t))

		// Get the value from the ObservableValue
		got := ov.Get()
		assert.Equal(t, true, got, "Get() should return the true boolean")

		// Set a new value to the ObservableValue and get it
		ov.Set(false)
		got = ov.Get()
		assert.Equal(t, false, got, "Get() should return the false boolean")
	})

	t.Run("String", func(t *testing.T) {
		// Create a new ObservableValue with a string value
		ov := NewObservableValue("hello", newTestCallback[string](t))

		// Get the value from the ObservableValue
		got := ov.Get()
		assert.Equal(t, "hello", got, "Get() should return the hello string")

		// Set a new value to the ObservableValue and get it
		ov.Set("world")
		got = ov.Get()
		assert.Equal(t, "world", got, "Get() should return the world string")
	})

	t.Run("Int", func(t *testing.T) {
		// Create a new ObservableValue with an int value
		ov := NewObservableValue(42, newTestCallback[int](t))

		// Get the value from the ObservableValue
		got := ov.Get()
		assert.Equal(t, 42, got, "Get() should return the 42 int")

		// Set a new value to the ObservableValue and get it
		ov.Set(100)
		got = ov.Get()
		assert.Equal(t, 100, got, "Get() should return the 100 int")
	})

	t.Run("Float", func(t *testing.T) {
		// Create a new ObservableValue with a float value
		ov := NewObservableValue(3.14, newTestCallback[float64](t))

		// Get the value from the ObservableValue
		got := ov.Get()
		assert.Equal(t, 3.14, got, "Get() should return the 3.14 float")

		// Set a new value to the ObservableValue and get it
		ov.Set(2.71)
		got = ov.Get()
		assert.Equal(t, 2.71, got, "Get() should return the 2.71 float")
	})

	t.Run("Complex", func(t *testing.T) {
		// Create a new ObservableValue with a complex value
		ov := NewObservableValue(1+2i, newTestCallback[complex128](t))

		// Get the value from the ObservableValue
		got := ov.Get()
		assert.Equal(t, 1+2i, got, "Get() should return the 1+2i complex")

		// Set a new value to the ObservableValue and get it
		ov.Set(3 + 4i)
		got = ov.Get()
		assert.Equal(t, 3+4i, got, "Get() should return the 3+4i complex")
	})

	t.Run("Struct", func(t *testing.T) {
		// Create a new ObservableValue with a struct value
		ov := NewObservableValue(struct{ Name string }{Name: "John"}, newTestCallback[struct{ Name string }](t))

		// Get the value from the ObservableValue
		got := ov.Get()
		assert.Equal(t, struct{ Name string }{Name: "John"}, got, "Get() should return the struct")

		// Set a new value to the ObservableValue and get it
		ov.Set(struct{ Name string }{Name: "Doe"})
		got = ov.Get()
		assert.Equal(t, struct{ Name string }{Name: "Doe"}, got, "Get() should return the struct")
	})

	t.Run("Slice", func(t *testing.T) {
		// Create a new ObservableValue with a slice value
		ov := NewObservableValue([]int{1, 2, 3}, newTestCallback[[]int](t))

		// Get the value from the ObservableValue
		got := ov.Get()
		assert.Equal(t, []int{1, 2, 3}, got, "Get() should return the slice")

		// Set a new value to the ObservableValue and get it
		ov.Set([]int{4, 5, 6})
		got = ov.Get()
		assert.Equal(t, []int{4, 5, 6}, got, "Get() should return the slice")
	})

	t.Run("Map", func(t *testing.T) {
		// Create a new ObservableValue with a map value
		ov := NewObservableValue(map[string]int{"one": 1, "two": 2}, newTestCallback[map[string]int](t))

		// Get the value from the ObservableValue
		got := ov.Get()
		assert.Equal(t, map[string]int{"one": 1, "two": 2}, got, "Get() should return the map")

		// Set a new value to the ObservableValue and get it
		ov.Set(map[string]int{"three": 3, "four": 4})
		got = ov.Get()
		assert.Equal(t, map[string]int{"three": 3, "four": 4}, got, "Get() should return the map")
	})

	t.Run("Pointer", func(t *testing.T) {
		// Create a new ObservableValue with a pointer value
		n1 := 1
		ov := NewObservableValue(&n1, newTestCallback[*int](t))

		// Get the value from the ObservableValue
		got := ov.Get()
		assert.NotNil(t, got, "Get() should return the pointer")

		// Set a new value to the ObservableValue and get it
		n2 := 2
		ov.Set(&n2)
		got = ov.Get()
		assert.NotNil(t, got, "Get() should return the pointer")
	})
}
