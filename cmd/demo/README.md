# EL Programming Language Demos

This directory contains comprehensive demonstrations of the EL programming language features and capabilities.

## Demo Files

### 01_basic_syntax.go
Demonstrates fundamental language constructs:
- Basic literals (numbers, strings, lists, nil)
- Variable binding with `let`
- Nested let expressions
- Function definition and calls
- Infix expressions
- Comments

### 02_arithmetic.go
Shows arithmetic operations and mathematical functions:
- Basic arithmetic (add, sub, mul, div, mod)
- Infix syntax for arithmetic
- Complex expressions with precedence
- Negative numbers
- Comparison operations
- Mathematical functions (abs, max, min, power)

### 03_lists.go
Comprehensive list operations and manipulation:
- List creation and literals
- List length and access operations
- List slicing and range generation
- List concatenation with unwrapping
- List mapping and filtering
- List reduction and searching
- List reversal and sorting

### 04_functions.go
Function definition and closure features:
- Basic and multi-parameter functions
- Variable-arity functions
- Higher-order functions
- Closures and lexical scoping
- Functions returning functions
- Partial application and currying
- Function composition
- Anonymous functions
- Functions as data

### 05_recursion.go
Recursion and mutual recursion examples:
- Basic recursion (factorial, fibonacci)
- Mutual recursion (even/odd)
- Complex recursive algorithms
- Tree recursion
- Recursive list operations
- Binary search
- Tower of Hanoi

### 06_matching.go
Pattern matching capabilities:
- Basic value matching
- Multiple patterns
- String and boolean matching
- List length matching
- Complex conditional matching
- Type-based dispatch
- Nested matching
- Range matching
- Error handling with matching

### 07_unwrapping.go
Argument unwrapping features:
- Basic unwrapping syntax
- Unwrapping in function calls
- Mixed arguments with unwrapping
- Nested unwrapping
- Unwrapping in list construction
- Unwrapping with custom functions
- Higher-order function unwrapping
- Data transformation with unwrapping

### 08_types.go
Type system and introspection:
- Basic type introspection
- Type checking functions
- Type-safe operations
- Type conversion
- Type-based dispatch
- Type validation
- Type-safe list operations
- Type introspection for debugging
- Type-based function selection

### 09_advanced.go
Advanced features and complex examples:
- Advanced data structures (trees, graphs)
- Object-oriented programming simulation
- Functional programming patterns
- Currying and partial application
- Memoization
- Lazy evaluation simulation
- Stream processing
- Event-driven programming simulation
- State machines

### 10_performance.go
Performance testing and benchmarking:
- Algorithm complexity demonstration
- Iterative vs recursive approaches
- List operations performance
- Sorting algorithms
- Memory usage simulation
- Caching and memoization performance
- String operations performance
- Pattern matching performance

### 11_error_handling.go
Error handling and edge cases:
- Safe operations (division, list access)
- Type validation
- Error propagation
- Try-catch simulation
- Input validation
- Arithmetic edge cases
- List operation edge cases
- Function edge cases
- Matching edge cases
- Unwrapping edge cases
- Recursion edge cases
- Error recovery
- Comprehensive error handling

## Running the Demos

To run any demo:

```bash
go run cmd/demo/01_basic_syntax.go
go run cmd/demo/02_arithmetic.go
# ... and so on
```

Or run all demos:

```bash
for demo in cmd/demo/*.go; do
  echo "Running $demo"
  go run "$demo"
  echo "---"
done
```

## Language Features Demonstrated

### Core Syntax
- S-expressions with parentheses
- Infix blocks with curly braces `{a + b * c}`
- List literals with square brackets `[1 2 3]`
- Comments with `#`

### Data Types
- Integers: `42`, `-17`
- Strings: `"hello world"`
- Lists: `[1 2 3]`, `["a" "b" "c"]`
- Functions: `(lambda x {x + 1})`
- Nil: `nil`
- Types: `(type 42)` → `"int"`

### Control Flow
- `let` for variable binding
- `lambda` for function definition
- `match` for pattern matching
- Recursion and mutual recursion

### Advanced Features
- Lexical closures
- Argument unwrapping with `*`
- Type introspection
- Higher-order functions
- Functional programming patterns
- Error handling
- Performance optimization techniques

## Learning Path

1. Start with `01_basic_syntax.go` to understand fundamental concepts
2. Move through `02_arithmetic.go` and `03_lists.go` for basic operations
3. Explore `04_functions.go` and `05_recursion.go` for functional programming
4. Study `06_matching.go` and `07_unwrapping.go` for advanced syntax
5. Examine `08_types.go` for type system understanding
6. Dive into `09_advanced.go` for complex examples
7. Review `10_performance.go` for optimization techniques
8. Study `11_error_handling.go` for robust programming practices

Each demo builds upon previous concepts and demonstrates increasingly sophisticated programming techniques in the EL language.
