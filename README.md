# El - A Functional Lisp-like Programming Language

El is a functional programming language inspired by Lisp, featuring a simple yet powerful syntax with support for higher-order functions, pattern matching, and a rich set of built-in operations.

## Features

### Core Language Features
- **Functional Programming**: First-class functions, closures, and lambda expressions
- **Pattern Matching**: Powerful `match` expressions for conditional logic
- **Lexical Scoping**: Proper closure support with lexical scoping
- **Tail Recursion**: Optimized tail recursion for efficient recursive algorithms
- **Type System**: Dynamic typing with runtime type introspection
- **S-Expressions**: Lisp-style syntax with parentheses-based expressions

### Syntax Features
- **Infix Arithmetic**: Support for infix operators like `{1 + 2 * 3}`
- **Arrow Functions**: Lambda syntax with `{x y => {x + y}}`
- **List Literals**: Square bracket notation `[1 2 3]` for lists
- **String Literals**: Double-quoted strings with escape sequences
- **Comments**: Hash-prefixed comments `# this is a comment`

### Built-in Operations
- **Arithmetic**: `add`, `sub`, `mul`, `div`, `mod`
- **Comparison**: `eq`, `ne`, `lt`, `le`, `gt`, `ge`
- **Lists**: `list`, `len`, `slice`, `range`
- **I/O**: `print` for output
- **Control Flow**: `let` for variable binding, `match` for pattern matching
- **Functions**: `lambda` for function definition

## Installation

### Prerequisites
- Go 1.24.5 or later

### Building from Source
```bash
git clone <repository-url>
cd el
go mod tidy
go build -o el cmd/basic/main.go
```

## Quick Start

### Hello World
```lisp
(print "Hello, World!")
```

### Basic Arithmetic
```lisp
(let
    x {1 + 2 * 3}
    y {10 - 5}
    _ (print "x =" x)
    _ (print "y =" y)
    {x + y}
)
```

### Function Definition
```lisp
(let
    square (lambda x {x * x})
    _ (print "5 squared =" (square 5))
    nil
)
```

### List Operations
```lisp
(let
    numbers [1 2 3 4 5]
    doubled (map numbers (lambda x {x * 2}))
    _ (print "original:" numbers)
    _ (print "doubled:" doubled)
    nil
)
```

## Language Syntax

### Basic Syntax
- **Numbers**: `1`, `42`, `-10`
- **Strings**: `"hello"`, `"world"`
- **Booleans**: `true`, `false`
- **Nil**: `nil` (empty value)

### Expressions
- **Function Calls**: `(function arg1 arg2 ...)`
- **Variable Binding**: `(let var1 value1 var2 value2 ... result)`
- **Lambda Functions**: `(lambda param1 param2 ... body)`
- **Pattern Matching**: `(match value pattern1 result1 pattern2 result2 ... default)`

### Special Syntax
- **Infix Arithmetic**: `{1 + 2 - 3}`
- **Arrow Functions**: `{x y => {x + y}}`
- **List Literals**: `[1 2 3]`
- **Unwrap Operator**: `*list` (unwraps list elements as arguments)

## Examples

### Fibonacci Sequence
```lisp
(let
    fib (lambda n (match {n <= 1}
        true n
        {fib {n - 1} + fib {n - 2}}
    ))
    _ (print "fib(10) =" (fib 10))
    nil
)
```

### List Processing
```lisp
(let
    numbers [1 2 3 4 5]
    evens (filter numbers (lambda x {x % 2 == 0}))
    sum (fold numbers 0 (lambda acc x {acc + x}))
    _ (print "evens:" evens)
    _ (print "sum:" sum)
    nil
)
```

### Higher-Order Functions
```lisp
(let
    apply-twice (lambda f x (f (f x)))
    increment (lambda x {x + 1})
    result (apply-twice increment 5)
    _ (print "apply-twice increment 5 =" result)
    nil
)
```

## Running Programs

### Interactive Mode
```bash
./el
```

### From File
```bash
./el < program.el
```

## Built-in Functions Reference

### Arithmetic Functions
- `(add a b ...)` - Addition
- `(sub a b ...)` - Subtraction  
- `(mul a b ...)` - Multiplication
- `(div a b ...)` - Division
- `(mod a b)` - Modulo

### Comparison Functions
- `(eq a b)` - Equality
- `(ne a b)` - Not equal
- `(lt a b)` - Less than
- `(le a b)` - Less than or equal
- `(gt a b)` - Greater than
- `(ge a b)` - Greater than or equal

### List Functions
- `(list a b ...)` - Create list
- `(len list)` - Get length
- `(slice list indices)` - Get elements by indices
- `(range start end)` - Create range

### Control Flow
- `(let var1 val1 var2 val2 ... result)` - Variable binding
- `(match value pattern1 result1 ... default)` - Pattern matching
- `(lambda params ... body)` - Function definition

### I/O
- `(print value ...)` - Print values

## Advanced Features

### Closures
Functions capture their lexical environment:
```lisp
(let
    make-counter (lambda start (lambda {start + 1}))
    counter (make-counter 10)
    _ (print (counter))  ; prints 11
    _ (print (counter))  ; prints 12
    nil
)
```

### Recursive Functions
```lisp
(let
    factorial (lambda n (match {n <= 1}
        true 1
        {n * factorial {n - 1}}
    ))
    _ (print "5! =" (factorial 5))
    nil
)
```

### Type Introspection
```lisp
(let
    x 42
    y "hello"
    z [1 2 3]
    _ (print "type of 42:" (type x))
    _ (print "type of 'hello':" (type y))
    _ (print "type of [1 2 3]:" (type z))
    nil
)
```

## Contributing

Contributions are welcome! Please feel free to submit issues, feature requests, or pull requests.

## License

This project is licensed under the terms specified in the LICENSE file.
