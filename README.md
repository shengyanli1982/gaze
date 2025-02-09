English | [中文](./README_CN.md)

<div align="center">
	<img src="assets/logo.png" alt="logo" width="500px">
</div>

[![Go Report Card](https://goreportcard.com/badge/github.com/shengyanli1982/gaze)](https://goreportcard.com/report/github.com/shengyanli1982/gaze)
[![Build Status](https://github.com/shengyanli1982/gaze/actions/workflows/test.yaml/badge.svg)](https://github.com/shengyanli1982/gaze/actions)
[![Go Reference](https://pkg.go.dev/badge/github.com/shengyanli1982/gaze.svg)](https://pkg.go.dev/github.com/shengyanli1982/gaze)

# GAZE: A Lightweight Go Package for Triggering Callbacks on Variable Changes

When developing applications, it's common to need code that triggers when a variable's value changes. Typically, this requires writing a significant amount of boilerplate code to monitor these changes, which can be tedious and time-consuming. `GAZE` is here to help!

### Why Create `GAZE`?

Writing repetitive and complex code not only wastes time but also drains the joy out of programming. I wanted to save my time and help others do the same, which is why I created `GAZE`. This package simplifies the process of monitoring variable changes, allowing you to focus on what truly matters in your code.

### Why Use `GAZE`?

Do you want to streamline your coding process and reduce unnecessary complexity? `GAZE` offers a straightforward solution that allows you to:

-   **Focus on Your Code**: Spend more time on core functionalities and less on repetitive tasks.
-   **Enhance Productivity**: Complete your work faster and with fewer errors.
-   **Enjoy More Free Time**: Finish your tasks early so you can enjoy your personal time.

Here’s why `GAZE` stands out:

-   **Simple**: `GAZE` is easy to use and understand.
-   **Lightweight**: It has no dependencies, keeping your project lightweight.
-   **Reliable**: Thoroughly tested for dependability.
-   **Callback Support**: Seamlessly integrates callback functions when variable values change.

### Practical Use Cases

`GAZE` can be particularly useful in the following scenarios:

-   **Configuration Management**: Automatically reload configurations when changes are detected.
-   **Logging**: Log changes to critical variables for auditing or debugging purposes.
-   **Automation**: Execute specific tasks in real-time when certain conditions are met.
-   **Monitoring**: Monitor the values of important variables and take action when changes occur.
-   **Notification Systems**: Trigger notifications or alert systems when a variable changes.

# Benchmark

> [!TIP]
>
> Because `GAZE` uses `reflect` to compare old and new values, this impacts performance. However, for most use cases, its performance is acceptable.

```bash
$ go test -benchmem -run=^$ -bench .
goos: darwin
goarch: amd64
pkg: github.com/shengyanli1982/gaze
cpu: Intel(R) Xeon(R) CPU E5-2643 v2 @ 3.50GHz
BenchmarkReactiveValue_IntGet-12       	1000000000	         0.2745 ns/op	       0 B/op	       0 allocs/op
BenchmarkReactiveValue_IntSet-12       	16565497	        70.25 ns/op	      15 B/op	       1 allocs/op
BenchmarkReactiveValue_IntSetGet-12    	16543251	        72.96 ns/op	      15 B/op	       1 allocs/op
BenchmarkStd_IntSet-12                 	1000000000	         0.2729 ns/op	       0 B/op	       0 allocs/op
BenchmarkStd_IntGet-12                 	1000000000	         0.2771 ns/op	       0 B/op	       0 allocs/op
BenchmarkStd_IntSetGet-12              	1000000000	         0.2740 ns/op	       0 B/op	       0 allocs/op
```

# Installation

To install `GAZE`, simply run:

```bash
go get github.com/shengyanli1982/gaze
```

# Quick Start

`GAZE` uses generics to define variable types, making it compatible with most types of variables.

### Methods

`GAZE` provides two primary methods: `Set` and `Get` to set and get the value of the variable. If the value of the variable changes, the callback function will be triggered.

-   **Set**: Set the value of the variable.
-   **Get**: Get the value of the variable.

### Callbacks

`GAZE` triggers callback functions when a variable's value changes. For asynchronous callbacks, you can use goroutines within the callback function or place the value in a `channel` or `queue` for asynchronous processing.

-   **OnChange**: The callback function is triggered when the variable's value changes upon being set.

### Example

Using `GAZE` is very simple. Here is an example:

```go
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
```

**Result**

```bash
$ go run demo.go
got: 11
>> OnChange: 11 -> 3
got: 3
```

# Contributing

We welcome contributions! If you have any ideas, suggestions, or issues, please open a pull request or an issue on our GitHub repository.

# License

This project is licensed under the MIT License. See the LICENSE file for more details.
