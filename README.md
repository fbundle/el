# EL Programming Language

`el` is a tiny, Lisp-like language implemented in Go with a minimal AST, a simple lexer/parser, and a small but expressive runtime. It supports first-class functions, lexical closures, list operations, lightweight matching, infix sugar, argument unwrapping, and an extension mechanism for builtins.

## Quick Start

### Run the Basic Example
```bash
go run ./cmd/basic
```

### Run All Demos
```bash
# Run individual demos
go run ./cmd/demo/01_basic_syntax.go
go run ./cmd/demo/02_arithmetic.go
go run ./cmd/demo/03_lists.go
# ... see cmd/demo/README.md for complete list

# Or run all demos at once
for demo in cmd/demo/*.go; do
  echo "Running $demo"
  go run "$demo"
  echo "---"
done
```

### Explore the Demos
The `cmd/demo/` directory contains 11 comprehensive demonstration files covering all language features:
- **01_basic_syntax.go** - Fundamental language constructs
- **02_arithmetic.go** - Mathematical operations and functions  
- **03_lists.go** - List manipulation and operations
- **04_functions.go** - Function definition and closures
- **05_recursion.go** - Recursion and mutual recursion
- **06_matching.go** - Pattern matching capabilities
- **07_unwrapping.go** - Argument unwrapping features
- **08_types.go** - Type system and introspection
- **09_advanced.go** - Advanced features and complex examples
- **10_performance.go** - Performance testing and benchmarking
- **11_error_handling.go** - Error handling and edge cases

See `cmd/demo/README.md` for detailed descriptions of each demo.

## Language overview

- **Core syntax**: S-expressions `(head arg1 arg2 ...)` with two sugars:
  - **Infix blocks**: `{ a + b * c }` → folded to prefix calls like `(add a (mul b c))`
  - **List literals**: `[1 2 3]` → `(list 1 2 3)`
- **Atoms**: numbers (`123`), strings (`"hello"`), names (`x`, `print`, `fib`), special `*` (unwrap), and `nil`.
- **Evaluation**: Names resolve to the current frame or parse as literals; lists evaluate to function applications.
- **Functions**: `(lambda x y body)` creates a closure capturing the current frame (excluding parameters).
- **Binding**: `(let name1 expr1 name2 expr2 ... body)` binds in a new local frame, then evaluates `body`.
- **Matching**: `(match cond k1 v1 k2 v2 ... default)` compares `cond` to keys in order, returning the first matching value; otherwise returns `default`.
- **Unwrapping**: Use `*` to splice list arguments into calls: `(f *[1 2 3])` becomes `(f 1 2 3)`.

## Syntax and parsing

### Tokens and comments
- **whitespace** separates tokens
- **comments** start with `#` and extend to end of line

### Strings
- Double-quoted with JSON escapes: `"\n"`, `"\t"`, `"\\"`, etc.

### Numbers
- Decimal integers (for now): `0`, `42`, `-7`.

### Lists and infix sugar
- `[a b c]` desugars to `(list a b c)`.
- `{a op b op c}` desugars left-associatively: `{1 + 2 - 3}` → `(sub (add 1 2) 3)`.

## Core forms and semantics

### `let`
Bind names in a fresh local frame and evaluate a final body.

```lisp
(let
  x 1
  y 2
  (add x y))       ; => 3
```

Rules:
- Requires an odd number of args: `name expr` pairs followed by a body.
- Names must be bare identifiers.

### `lambda`
Create a function with positional parameters and lexical closure.

```lisp
(let
  inc (lambda x (add x 1))
  (inc 41))        ; => 42
```

Rules:
- At least one argument: zero or more parameter names, then a body.
- Closure captures the surrounding frame at creation time, excluding parameters.
- Arity is enforced: calling with fewer arguments errors with "not enough arguments".

### Function application

```lisp
((lambda x y (add x y)) 1 2)  ; => 3
```

Application evaluates the operator to a function, evaluates arguments (with unwrapping), creates a call frame from the current frame plus closure and parameter bindings, then evaluates the body in that frame.

### `match`
Select among branches by comparing a value to keys.

```lisp
(match n
  0 "zero"
  1 "one"
  "many")
```

Rules:
- Even number of args after the condition: `k1 v1 k2 v2 ... default`.
- Keys are evaluated; comparison uses Go equality on underlying objects when comparable.
- For non-comparable objects (e.g., functions), the runtime guards to prevent panics (errors or safe fallback, depending on build).

## Data types

- **int**: `1`, `2`, `-5`. Prints as decimal.
- **string**: `"hello"`. Prints the raw string (without quotes).
- **list**: `[1 2 3]`. Prints as `[1 2 3]`.
- **function**: `(lambda ...)`. Prints as `(closure{...}; params -> body)`.
- **type**: Type objects (e.g., `(type 1)` → `int`). Types are first-class; repeated `type` calls increase level.
- **nil**: Dedicated `nil` value with its own type.
- **unwrap**: The special `*` used only to splice list arguments in calls.

### `type`

```lisp
(type 1)           ; => int
(type [1 2])       ; => list
(type (lambda x x)); => function
```

## Builtins and extensions

### Core builtins
- **type**: `(type x)` → returns the type of `x`.
- **let**: `(let name1 expr1 ... body)` → lexical bindings.
- **lambda**: `(lambda x y body)` → function creation with closure.
- **match**: `(match cond k1 v1 ... default)` → conditional selection.

### Standard extensions (installed by `runtime_ext.NewBasicRuntime()`)

- **Arithmetic**: `add`, `sub`, `mul`, `div`, `mod`
  - Examples: `(add 1 2 3)` → `6`, `(sub 10 3 2)` → `5`
- **Comparison**: `eq`, `ne`, `lt`, `le`, `gt`, `ge`
  - Examples: `(eq 1 1)` → `1` (true), `(lt 1 2)` → `1`
  - Truth values: `true` is `1`, `false` is `0`.
- **Lists**: `list`, `len`, `slice`, `range`
  - `(list a b c)` → `[a b c]`
  - `(len [1 2 3])` → `3`
  - `(slice [10 11 12 13] [0 2])` → `[10 12]`
  - `(range 3 6)` → `[3 4 5]`
- **I/O**: `print` prints its arguments separated by a space and returns `nil`.

### Operator/template helpers

`runtime_ext.WithTemplate` injects handy aliases and helpers:

```lisp
; operator aliases
+ add - sub x mul / div % mod
== eq != ne <= le < lt > gt >= ge

; lists
get (lambda l i (unit * (slice l (range i (add i 1)))))
head (lambda l (get l 0))
rest (lambda l (slice l (range 1 (len l))))

; map
map (lambda l f (match (len l)
  0 []
  (let
    first_elem (head l)
    first_elem2 (f first_elem)
    rest_elems (rest l)
    rest_elems2 (map rest_elems f)
    (list first_elem2 *rest_elems2))))
```

## Unwrapping with `*`

Use `*` to splice elements of a list into an argument list.

```lisp
(print [1 2 3] *[1 2 3])   ; prints: [1 2 3] 1 2 3
```

Unwrapping only works in call argument positions. Nested `*` is supported; unwrapping a non-list errors.

## Examples

### Basic Syntax
```lisp
(let
  x 42
  y "hello world"
  z [1 2 3]
  (print "x =" x)
  (print "y =" y)
  (print "z =" z)
  nil)
```

### Fibonacci
```lisp
(let
  fib (lambda n (match {n <= 1}
    true n
    (let p (fib {n - 1}) q (fib {n - 2}) {p + q})))
  (print "fib(20)=" (fib 20))
  nil)
```

### Mutual Recursion
```lisp
(let
  even (lambda n (match {n <= 0} true true (odd {n - 1})))
  odd  (lambda n (match {n <= 0} true false (even {n - 1})))
  (print "evens and odds:" [(odd 10) (even 10) (odd 11) (even 11)])
  nil)
```

### List Operations
```lisp
(let
  numbers [1 2 3 4 5]
  doubled (map numbers (lambda x {x * 2}))
  sum (reduce numbers + 0)
  (print "numbers:" numbers)
  (print "doubled:" doubled)
  (print "sum:" sum)
  nil)
```

### Type Introspection
```lisp
(let
  x 1
  y (list 1 2 3)
  z (lambda x {x + 1})
  (print (list (type x) (type y) (type z))))
```

### Argument Unwrapping
```lisp
(let
  numbers [1 2 3 4 5]
  (print "sum of numbers:" (add *numbers))
  (print "product of numbers:" (mul *numbers))
  nil)
```

### Advanced Features
```lisp
(let
  # Object-oriented programming simulation
  make_point (lambda x y (let
    get_x (lambda () x)
    get_y (lambda () y)
    move (lambda dx dy (make_point {x + dx} {y + dy}))
    (lambda method (match method
      "get_x" (get_x)
      "get_y" (get_y)
      "move" move
      "Error: unknown method"
    ))
  ))
  
  point (make_point 3 4)
  (print "Point x:" (point "get_x"))
  (print "Point y:" (point "get_y"))
  nil)
```

## Equality notes

- Comparable objects use Go equality in `match`.
- Non-comparable objects (e.g., functions) are guarded against to avoid panics (error or safe fallback, depending on build).

## Runtime and implementation

- Evaluation via `Runtime.Step(context, frame, expr)`.
- Names resolve from the current frame first; otherwise parsed as literals.
- Frames: persistent ordered maps; lists: persistent sequences.
- Tail calls are not optimized; consider `MaxStackDepth` if you need recursion limits.

## Project Structure

```
el/
├── ast/           # Abstract Syntax Tree definitions
├── parser/        # Lexer and parser implementation
├── runtime/       # Core runtime and evaluation engine
├── runtime_ext/   # Built-in extensions and standard library
├── cmd/
│   ├── basic/     # Basic example program
│   └── demo/      # Comprehensive demonstration programs
├── vendor/        # External dependencies
├── go.mod         # Go module definition
└── README.md      # This file
```

## Development

### Packages
- **`ast`**: Abstract Syntax Tree definitions (`Name`, `Lambda`)
- **`parser`**: Lexer and parser with support for infix expressions and list literals
- **`runtime`**: Core evaluation engine with frame-based execution
- **`runtime_ext`**: Built-in extensions (arithmetic, lists, I/O) and template system

### Running Examples
```bash
# Basic example
go run ./cmd/basic

# Individual demos
go run ./cmd/demo/01_basic_syntax.go

# All demos
for demo in cmd/demo/*.go; do go run "$demo"; done
```

### Building
```bash
# Build the project
go build ./...

# Run tests (if any)
go test ./...
```

## Language Design Philosophy

EL is designed as a minimal but expressive Lisp-like language with the following principles:

1. **Simplicity**: Minimal syntax with maximum expressiveness
2. **Functional**: First-class functions, closures, and functional programming patterns
3. **Extensible**: Built-in extension mechanism for adding new functionality
4. **Practical**: Includes practical features like infix syntax and argument unwrapping
5. **Educational**: Clear implementation that demonstrates language design concepts

## Contributing

The language is designed to be educational and extensible. Key areas for contribution:

- Additional built-in functions and extensions
- Performance optimizations
- More comprehensive error handling
- Additional demo programs
- Documentation improvements

## License

MIT
