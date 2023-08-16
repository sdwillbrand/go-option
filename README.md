# go-opt: Golang Option API

## Table of Contents

1. **Introduction**
   - 1.1 Purpose
   - 1.2 Features
   - 1.3 Project Structure

2. **Getting Started**
   - 2.1 Installation
   - 2.2 Usage Example

3. **API Reference**
   - 3.1 Option Struct
   - 3.2 `Some(value T) Option`
   - 3.3 `None() Option`
   - 3.4 `Unwrap() T`
   - 3.5 `UnwrapOr(def T) T`
   - 3.6 `UnwrapOrDefault(f func() T) T`
   - 3.7 `IsSome() bool`
   - 3.8 `IsNone() bool`
   - 3.9 `IsSomeAnd(predicate func(T) bool) bool`
   - 3.10 `Replace(value T) Option`
   - 3.11 `Expect(msg string) T`
   - 3.12 `Clone() Option[T]`

4. **Contributing**
   - 4.1 Bug Reports and Feature Requests
   - 4.2 Pull Requests

5. **License**

---

## 1. Introduction

Welcome to the go-opt documentation! go-opt is a Golang option API inspired by the Rust Option API. It provides the `Option` struct, a wrapper around a value that allows you to perform various operations to handle optional values gracefully.

### 1.1 Purpose

go-opt aims to bring the powerful features of the Rust Option API to Golang developers. By providing methods to handle optional values, go-opt simplifies the way you work with potentially absent data, enhancing code readability and maintainability.

### 1.2 Features

- **Option Struct:** The `Option` struct wraps a value of any type and offers a range of methods to handle it.
- **Some and None:** Create `Option` instances using the `Some` and `None` constructors.
- **Value Extraction:** Extract values from `Option` instances using methods like `Unwrap`, `UnwrapOr`, and `UnwrapOrDefault`.
- **Predicate Checks:** Utilize methods like `IsSome`, `IsNone`, and `IsSomeAnd` to check the presence and validity of values.
- **Error Handling:** The `Expect` function allows for controlled value unwrapping with the ability to provide custom error messages.
- **Method Chaining:** Chain methods on `Option` instances to perform multiple operations sequentially.
- **Cloning:** Create a new `Option` instance with the same value using the `Clone` function.

### 1.3 Project Structure

```
go-opt/
|-- option.go
|-- LICENSE
|-- README.md
|-- example/
|   |-- main.go
|-- tests/
|   |-- go-opt_test.go
|-- .gitignore
```

- `option.go`: Defines the `Option` struct and its methods.
- `example`: Contains an example usage of go-opt.
- `tests`: Unit tests for the go-opt library.
- `LICENSE`: The project's license information.
- `README.md`: The main documentation file.

---

## 2. Getting Started

### 2.1 Installation

To use go-opt in your project, use the `go get` command:

```sh
go get github.com/countersoda/go-opt
```

### 2.2 Usage Example

Here's a basic example of how to use go-opt in your application:

```go
package main

import (
	"fmt"
	"github.com/countersoda/go-opt"
)

func main() {
	// Create an Option with a value
	someOpt := opt.Some(42)

	// Unwrap the value and print it
	fmt.Println(someOpt.Unwrap()) // Output: 42

	// Create an Option with no value
	noneOpt := opt.None[any]()

	// Check if it's None
	if noneOpt.IsNone() {
		fmt.Println("Option is None")
	}
}
```

---

## 3. API Reference

### 3.1 Option Struct

The `Option` struct is the core type provided by go-opt. It wraps a value of any type and provides methods for value extraction and manipulation. The `Option` type can be constructed using the `Some` and `None` constructor functions.

### 3.2 `Some(value T) Option[T]`

The `Some` function constructs an `Option[T]` instance with the provided value.

**Parameters:**

- `value T`: The value to be wrapped.

**Returns:**

- `Option`: An `Option` instance wrapping the provided value.

### 3.3 `None() Option[T]`

The `None` function constructs an empty `Option` instance.

**Returns:**

- `Option[T]`: An empty `Option[T]` instance.

### 3.4 `Unwrap() T`

The `Unwrap` method extracts the value from the `Option`.

**Returns:**

- `T`: The unwrapped value.

### 3.5 `UnwrapOr(def T) T`

The `UnwrapOr` method extracts the value from the `Option` or returns a default value if the `Option` is empty.

**Parameters:**

- `def T`: The default value to return if the `Option` is empty.

**Returns:**

- `T`: The unwrapped value or the default value.

### 3.6 `UnwrapOrDefault() T`

The `UnwrapOrDefault` method extracts the value from the `Option` or returns the default value of `T`

**Returns:**

- `T`: The unwrapped value or the default value of `T`

### 3.7 `IsSome() bool`

The `IsSome` method checks if the `Option` contains a value.

**Returns:**

- `bool`: `true` if the `Option` contains a value, `false` otherwise.

### 3.8 `IsNone() bool`

The `IsNone` method checks if the `Option` is empty.

**Returns:**

- `bool`: `true` if the `Option` is empty, `false` otherwise.

### 3.9 `IsSomeAnd(predicate func(T) bool) bool`

The `IsSomeAnd` method checks if the `Option` contains a value and the provided predicate function returns `true` for that value.

**Parameters:**

- `predicate func(T) bool`: The predicate function to check the value.

**Returns:**

- `bool`: `true` if the `Option` contains a value and the predicate is `true`, `false` otherwise.

### 3.10 `Replace(value T) Option`

The `Replace` method replaces the value inside the `Option` with the provided value.

**Parameters:**

- `value T`: The new value.

**Returns:**

- `Option`: The modified `Option` instance.

### 3.11 `Expect(msg string) T`

The `Expect` method unwraps the value from the `Option` and panics

 with a custom error message if the `Option` is empty.

**Parameters:**

- `msg string`: The error message to panic with.

**Returns:**

- `T`: The unwrapped value.

### 3.12 `Clone() Option[T]`

The `Clone` method creates a new `Option` instance with the same value as the original `Option`.

**Returns:**

- `Option[T]`: A new `Option` instance with the same value.

---

## 4. Contributing

Contributions to go-opt are welcome! If you encounter any bugs or have ideas for improvements, please follow the guidelines below.

### 4.1 Bug Reports and Feature Requests

If you find a bug or have a feature request, please open an issue on the project's GitHub repository. Provide as much detail as possible to help us understand and address the problem.

### 4.2 Pull Requests

If you want to contribute code to go-opt, follow these steps:

1. Fork the repository.
2. Create a new branch for your feature or bug fix.
3. Make your changes and ensure that the code passes all tests.
4. Open a pull request, explaining the changes you've made.

We will review your pull request and provide feedback as needed.

---

## 5. License

go-opt is licensed under the [MIT License](LICENSE).

---

Feel free to explore, use, and contribute to the go-opt project. Happy coding with enhanced optional value handling! 🚀