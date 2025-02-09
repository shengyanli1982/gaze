[English](./README.md) | 中文

<div align="center">
	<img src="assets/logo.png" alt="logo" width="500px">
</div>

[![Go Report Card](https://goreportcard.com/badge/github.com/shengyanli1982/gaze)](https://goreportcard.com/report/github.com/shengyanli1982/gaze)
[![Build Status](https://github.com/shengyanli1982/gaze/actions/workflows/test.yaml/badge.svg)](https://github.com/shengyanli1982/gaze/actions)
[![Go Reference](https://pkg.go.dev/badge/github.com/shengyanli1982/gaze.svg)](https://pkg.go.dev/github.com/shengyanli1982/gaze)

# GAZE：一个轻量级用于在变量值变化时触发回调的 Go 软件包

在开发应用程序时，常常需要编写代码来监控变量值的变化并触发相应的操作。通常，这需要写大量的样板代码来监控这些变化，这既繁琐又耗时。而 `GAZE` 就是专门解决这个问题！

### 为什么创建 `GAZE`？

编写重复和复杂的代码不仅浪费时间，还会削弱编程的乐趣。我希望节省自己的时间，也帮助其他人节省时间，这就是我创建 `GAZE` 的原因。这个包简化了监控变量变化的过程，让你能够专注于代码中真正重要的部分。

### 为什么使用 `GAZE`？

你想简化编码过程，减少不必要的复杂性吗？`GAZE` 提供了一个简单明了的解决方案，允许你：

-   **专注于你的代码**：将更多时间花在核心功能上，减少重复任务。
-   **提高生产力**：更快地完成工作，并减少错误。
-   **享受更多自由时间**：早点完成任务，让你可以享受个人时光。

以下是 `GAZE` 的优势：

-   **简单**：`GAZE` 使用方便，易于理解。
-   **轻量**：它没有依赖项，保持项目的轻量级。
-   **可靠**：经过彻底测试，确保稳定性。
-   **支持回调**：在变量值变化时无缝集成回调函数。

### 实际使用场景

`GAZE` 适合以下场景：

-   **配置管理**：在检测到配置变化时自动重新加载。
-   **日志记录**：记录关键变量的变化，方便审计或调试。
-   **自动化**：在满足特定条件时实时执行特定任务。
-   **监控**：监控重要变量的值，并在发生变化时采取行动。
-   **通知系统**：在变量变化时触发通知或报警系统。

# 性能测试

> [!TIP]
>
> 因为 `GAZE` 使用 `reflect` 来比较旧值和新值，这会影响性能。然而，对于大多数使用场景来说，其性能是可以接受的。

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

# 安装

要安装 `GAZE`，只需运行：

```bash
go get github.com/shengyanli1982/gaze
```

# 快速开始

`GAZE` 使用泛型来定义变量类型，因此可以兼容大多数类型的变量。

### 方法

`GAZE` 提供两个主要方法：`Set` 和 `Get` 用于设置和获取变量的值。如果变量的值发生变化，将触发回调函数。

-   **Set**：设置变量的值。
-   **Get**：获取变量的值。

### 回调

`GAZE` 会在变量值变化时触发回调函数。对于异步回调，可以在回调函数中使用 goroutines，或者将值放入 `channel` 或 `queue` 中进行异步处理。

-   **OnChange**：当变量的值在设置时发生变化时，触发回调函数。

### 示例

使用 `GAZE` 非常简单。以下是一个示例：

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
    // 创建一个带有初始值的 ReactiveValue
    ov := gaze.NewReactiveValue(11, newTestCallback[int]())

    // 获取 ReactiveValue 的值
    got := ov.Get()
    fmt.Printf("got: %v\n", got)

    // 设置 ReactiveValue 的新值并再次获取
    ov.Set(3)
    got = ov.Get()
    fmt.Printf("got: %v\n", got)
}
```

**结果**

```bash
$ go run demo.go
got: 11
>> OnChange: 11 -> 3
got: 3
```

# 贡献

我欢迎贡献！如果你有任何想法、建议或问题，请在我的 GitHub 仓库中打开一个 pull request 或 issue。

# 许可证

该项目根据 MIT 许可证授权。详见 LICENSE 文件。
