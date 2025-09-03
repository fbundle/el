# EL Programming Language

EL is a functional programming language with S-expression syntax, designed for simplicity and expressiveness. It features lambda functions with closures, pattern matching, list operations, and tail-call optimization.

## Table of Contents

- [Installation](#installation)
- [Quick Start](#quick-start)
- [Language Features](#language-features)
- [Syntax Reference](#syntax-reference)
- [Built-in Functions](#built-in-functions)
- [Examples](#examples)
- [Command Line Interface](#command-line-interface)
- [Language Design](#language-design)
- [Contributing](#contributing)

## Installation

### Prerequisites

- Go 1.24.5 or later
- Git

### Building from Source

```bash
# Clone the repository
git clone <repository-url>
cd el

# Build the interpreter
go build -o el cmd/el/main.go

# Or run directly
go run cmd/el/main.go
```

## Quick Start

### Hello World

Create a file `hello.el`:

```el
(print "Hello, World!")
```

Run it:

```bash
go run cmd/el/main.go 1_hello.el
```

### Interactive REPL

Start the interactive REPL:

```bash
go run cmd/el/main.go -repl
```

## Language Features

### Core Features

- **S-expression Syntax**: Everything is an expression
- **Lambda Functions**: First-class functions with closures
- **Pattern Matching**: Powerful conditional expressions
- **Lists**: Immutable list data structures
- **Tail-call Optimization**: Efficient recursion
- **Type System**: Dynamic typing with runtime type checking

### Advanced Features

- **Higher-order Functions**: Functions that take or return functions
- **Function Composition**: Combine functions elegantly
- **Currying**: Partial function application
- **List Comprehensions**: Functional list processing
- **Closures**: Lexical scoping with captured variables

## Syntax Reference

### Basic Syntax

```el
# Comments start with #

# Literals
42                    # Numbers
"hello world"         # Strings
true false            # Booleans
nil                   # Null value
[1 2 3 4]            # Lists
```

### Variables and Bindings

```el
# Let binding
(let
    x 10
    y 20
    z {x + y}
    (print z)
)

# Function parameters
(lambda x y {x + y})
```

### Function Definitions

```el
# Named function
square (lambda x {x * x})

# Arrow function syntax
cube {x => {x * x * x}}

# Multi-parameter function
add (lambda x y {x + y})
```

### Pattern Matching

```el
# Simple pattern matching
(match x
    1 "one"
    2 "two"
    "other"
)

# Conditional pattern matching
(match {x > 0}
    true "positive"
    "non-positive"
)
```

### List Operations

```el
# List construction
[1 2 3 4 5]

# List access
(head [1 2 3])        # 1
(rest [1 2 3])        # [2 3]
(last [1 2 3])        # 3
(get [1 2 3] 1)       # 2

# List processing
(map [1 2 3] (lambda x {x * 2}))     # [2 4 6]
(filter [1 2 3] (lambda x {x % 2 == 0}))  # [2]
```

### Arithmetic and Comparison

```el
# Arithmetic operators
{x + y}               # Addition
{x - y}               # Subtraction
{x * y}               # Multiplication
{x / y}               # Division
{x % y}               # Modulo

# Comparison operators
{x == y}              # Equality
{x != y}              # Inequality
{x < y}               # Less than
{x <= y}              # Less than or equal
{x > y}               # Greater than
{x >= y}              # Greater than or equal

# Boolean operators
(and x y)             # Logical AND
(or x y)              # Logical OR
(not x)               # Logical NOT
```

## Built-in Functions

### Core Functions

- `let`: Variable binding
- `lambda`: Function definition
- `match`: Pattern matching
- `type`: Type introspection
- `print`: Output function

### List Functions

- `list`: Create lists
- `head`: Get first element
- `rest`: Get all but first element
- `last`: Get last element
- `init`: Get all but last element
- `get`: Get element by index
- `len`: Get length
- `slice`: Extract sublist
- `range`: Generate range
- `append`: Concatenate lists
- `cons`: Add element to front

### Arithmetic Functions

- `add`, `sub`, `mul`, `div`, `mod`: Arithmetic operations
- `eq`, `ne`, `lt`, `le`, `gt`, `ge`: Comparison operations

### Higher-order Functions

- `map`: Apply function to each element
- `filter`: Filter elements by predicate
- `foldl`, `foldr`: Reduce operations
- `compose`: Function composition
- `curry`, `uncurry`: Currying operations

### Utility Functions

- `reverse`: Reverse list
- `take`: Take first n elements
- `drop`: Drop first n elements
- `zip`: Zip two lists
- `unzip`: Unzip list of pairs
- `sum`, `product`: Aggregate operations
- `max_list`, `min_list`: Min/max operations

## Examples

### Basic Examples

```bash
# Hello World
go run cmd/el/main.go examples/1_hello.el

# Basic arithmetic
go run cmd/el/main.go examples/2_basic_arithmetic.el

# Lists and operations
go run cmd/el/main.go examples/3_lists_and_operations.el
```

### Advanced Examples

```bash
# Functions and lambdas
go run cmd/el/main.go examples/4_functions_and_lambdas.el

# Pattern matching
go run cmd/el/main.go examples/pattern_matching.el

# Advanced algorithms
go run cmd/el/main.go examples/advanced_algorithms.el

# Functional programming
go run cmd/el/main.go examples/functional_programming.el

# Type system
go run cmd/el/main.go examples/type_system.el

# Comprehensive demo
go run cmd/el/main.go examples/comprehensive_demo.el
```

### Example Programs

#### Fibonacci Sequence

```el
fib (lambda n (match {n <= 1}
    true n
    (let
        p (fib {n - 1})
        q (fib {n - 2})
        {p + q}
    )
))
(print "Fibonacci of 10:" (fib 10))
```

#### List Processing

```el
numbers [1 2 3 4 5 6 7 8 9 10]
squares (map numbers (lambda x {x * x}))
evens (filter numbers (lambda x {x % 2 == 0}))
(print "Squares:" squares)
(print "Evens:" evens)
```

#### Higher-order Functions

```el
# Function composition
compose (lambda f g (lambda x (f (g x))))
add_one (lambda x {x + 1})
square (lambda x {x * x})
add_one_then_square (compose square add_one)
(print "Add one then square 3:" (add_one_then_square 3))
```

## Command Line Interface

### Usage

```bash
go run cmd/el/main.go [OPTIONS] [FILE]
```

### Options

- `-v`: Verbose output - show parsed expressions
- `-timeout DUR`: Execution timeout (default: 30s)
- `-max-stack N`: Maximum stack depth (default: 10000)
- `-repl`: Start interactive REPL
- `-debug`: Enable debug mode with detailed error information
- `-help`: Show help message
- `-version`: Show version information

### Examples

```bash
# Run a file
go run cmd/el/main.go examples/1_hello.el

# Run from stdin
echo '(print "hello world")' | go run cmd/el/main.go

# Start interactive REPL
go run cmd/el/main.go -repl

# Verbose execution
go run cmd/el/main.go -v examples/demo1.el

# With timeout
go run cmd/el/main.go -timeout 10s examples/long_running.el
```

### Interactive REPL

The REPL provides an interactive environment for experimenting with EL:

```bash
go run cmd/el/main.go -repl
```

REPL Commands:
- `help`: Show help message
- `quit`, `exit`: Exit the REPL
- `clear`: Clear the screen

REPL Examples:
```el
el> (print "hello world")
=> hello world

el> (let x 42 (print x))
=> 42

el> (lambda x {x + 1})
=> {closure{...}; x => {x + 1}}

el> [1 2 3 4 5]
=> [1 2 3 4 5]

el> (map [1 2 3] (lambda x {x * 2}))
=> [2 4 6]
```

## Language Design

### Philosophy

EL is designed with the following principles:

1. **Simplicity**: Minimal syntax with maximum expressiveness
2. **Functional**: Functions are first-class values
3. **Immutable**: Data structures are immutable by default
4. **Composable**: Small functions compose into larger programs
5. **Expressive**: Powerful abstractions for common patterns

### Implementation Details

- **Parser**: Recursive descent parser with S-expression syntax
- **Runtime**: Tree-walking interpreter with tail-call optimization
- **Memory**: Immutable data structures with reference sharing
- **Types**: Dynamic typing with runtime type checking
- **Closures**: Lexical scoping with captured environments

### Performance Characteristics

- **Tail-call Optimization**: Recursive functions don't grow the stack
- **Immutable Data**: Safe sharing of data structures
- **Lazy Evaluation**: Potential for future optimization
- **Memory Efficient**: Reference sharing for immutable data

### Language Comparison

EL draws inspiration from:

- **Lisp**: S-expression syntax and homoiconicity
- **Haskell**: Functional programming and pattern matching
- **Scheme**: Lexical scoping and closures
- **ML**: Pattern matching and type inference concepts

## Advanced Topics

### Tail-call Optimization

EL implements tail-call optimization, allowing efficient recursive programming:

```el
# This won't cause stack overflow
countdown (lambda n (match {n <= 0}
    true 0
    (countdown {n - 1})
))
(countdown 10000)  # Efficient execution
```

### Closures and Lexical Scoping

Functions capture their lexical environment:

```el
make_counter (lambda start (lambda {start + 1}))
counter (make_counter 0)
(print (counter))  # 1
(print (counter))  # 2
```

### Higher-order Functions

Functions can take and return other functions:

```el
# Function that returns a function
make_multiplier (lambda factor (lambda x {x * factor}))
double (make_multiplier 2)
triple (make_multiplier 3)
(print (double 5))  # 10
(print (triple 5))  # 15
```

### Pattern Matching

Powerful conditional expressions:

```el
# Multiple conditions
grade (lambda score (match score
    100 "Perfect!"
    90 "Excellent"
    80 "Good"
    70 "Average"
    60 "Below Average"
    "Fail"
))
```

### Type Introspection

Runtime type checking and introspection:

```el
# Type checking
is_number (lambda x (match (type x)
    "number" true
    false
))

# Type-safe operations
safe_add (lambda a b (match (type a)
    "number" (match (type b)
        "number" {a + b}
        "Cannot add number and non-number"
    )
    "Cannot add non-number"
))
```

## Error Handling

### Common Errors

1. **Parse Errors**: Invalid syntax
2. **Runtime Errors**: Type mismatches, division by zero
3. **Stack Overflow**: Deep recursion without tail calls
4. **Timeout**: Long-running computations

### Debug Mode

Use `-debug` flag for detailed error information:

```bash
go run cmd/el/main.go -debug examples/broken.el
```

### Error Recovery

The interpreter provides helpful error messages:

```
Error: parse error: unexpected token ']' at position 15
Error: execution error: division by zero
Error: execution error: index out of bounds
```

## Contributing

### Development Setup

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests
5. Submit a pull request

### Code Style

- Follow Go conventions
- Add comments for public functions
- Write comprehensive tests
- Update documentation

### Testing

```bash
# Run all examples
for file in examples/*.el; do
    echo "Testing $file"
    go run cmd/el/main.go "$file"
done
```

### Reporting Issues

When reporting issues, please include:

1. EL code that reproduces the issue
2. Expected behavior
3. Actual behavior
4. Error messages
5. Go version and platform

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Acknowledgments

- Inspired by Lisp, Scheme, and functional programming languages
- Built with Go and modern language design principles
- Thanks to the functional programming community for inspiration

---

For more information, examples, and updates, visit the project repository.
