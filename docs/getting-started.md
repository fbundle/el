# Getting Started with El

## What is El?

El is a functional programming language inspired by Lisp, designed to be simple yet powerful. It features:

- **Functional Programming**: First-class functions, closures, and lambda expressions
- **Lisp-like Syntax**: S-expressions with parentheses-based syntax
- **Pattern Matching**: Powerful conditional logic with `match` expressions
- **Dynamic Typing**: Runtime type checking with type introspection
- **Lexical Scoping**: Proper closure support with lexical scoping
- **Tail Recursion**: Optimized recursive algorithms

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

## Your First Program

Create a file called `hello.el`:

```lisp
(let
    _ (print "Hello, World!")
    _ (print "Welcome to El!")
    nil
)
```

Run it:
```bash
./el < hello.el
```

## Basic Syntax

### Literals
```lisp
42                    # integer
"hello"              # string
true                 # boolean
nil                  # nil value
[1 2 3]              # list
```

### Variables
```lisp
(let
    x 42
    y "hello"
    _ (print x y)
    nil
)
```

### Functions
```lisp
(let
    square (lambda x {x * x})
    _ (print (square 5))
    nil
)
```

### Arithmetic
```lisp
{1 + 2 * 3}          # infix notation
(add 1 2 3)          # function notation
```

## Language Features

### 1. Functions and Closures

Functions are first-class values in El:

```lisp
(let
    make-adder (lambda n (lambda x {x + n}))
    add-5 (make-adder 5)
    _ (print (add-5 3))  # prints 8
    nil
)
```

### 2. Pattern Matching

Use `match` for conditional logic:

```lisp
(let
    classify (lambda n (match n
        0 "zero"
        1 "one"
        "many"
    ))
    _ (print (classify 0))  # prints "zero"
    nil
)
```

### 3. Lists

Lists are fundamental data structures:

```lisp
(let
    numbers [1 2 3 4 5]
    _ (print (len numbers))  # prints 5
    _ (print (slice numbers [0 2 4]))  # prints [1 3 5]
    nil
)
```

### 4. Recursion

El supports tail recursion:

```lisp
(let
    factorial (lambda n (match {n <= 1}
        true 1
        {n * factorial {n - 1}}
    ))
    _ (print (factorial 5))  # prints 120
    nil
)
```

## Running Examples

### Using the Demo Runner
```bash
# Run all demos
go run cmd/demo/main.go all

# Run specific demos
go run cmd/demo/main.go hello
go run cmd/demo/main.go functions
go run cmd/demo/main.go recursion
```

### Using the Shell Script
```bash
# Make the script executable (first time only)
chmod +x cmd/demo/run_demo.sh

# Run demos
./cmd/demo/run_demo.sh hello
./cmd/demo/run_demo.sh all
```

## Learning Path

1. **Start with Basics**: Run `hello` and `arithmetic` demos
2. **Learn Functions**: Explore `functions` and `closures` demos
3. **Practice Data**: Try `lists` and `types` demos
4. **Master Logic**: Study `matching` and `recursion` demos
5. **Advanced Topics**: Explore `advanced`, `algorithms`, and `data_structures` demos
6. **Math Problems**: Solve `math_problems` for practical applications

## Common Patterns

### Function Composition
```lisp
(let
    compose (lambda f g (lambda x (f (g x))))
    square (lambda x {x * x})
    increment (lambda x {x + 1})
    square-then-increment (compose increment square)
    _ (print (square-then-increment 3))  # prints 10
    nil
)
```

### List Processing
```lisp
(let
    sum-list (lambda lst (match (len lst)
        0 0
        {*slice lst [0] + sum-list (slice lst (range 1 (len lst)))}
    ))
    _ (print (sum-list [1 2 3 4 5]))  # prints 15
    nil
)
```

### Type Checking
```lisp
(let
    is-int (lambda x (eq (type x) "int"))
    _ (print (is-int 42))  # prints true
    _ (print (is-int "hello"))  # prints false
    nil
)
```

## Built-in Functions

### Arithmetic
- `add`, `sub`, `mul`, `div`, `mod`
- `eq`, `ne`, `lt`, `le`, `gt`, `ge`

### Lists
- `list`, `len`, `slice`, `range`

### Control Flow
- `let` - variable binding
- `match` - pattern matching
- `lambda` - function definition

### I/O
- `print` - output values

## Error Handling

El provides runtime error checking:

```lisp
(let
    safe-div (lambda a b (match b
        0 "division by zero"
        (div a b)
    ))
    _ (print (safe-div 10 0))  # prints "division by zero"
    _ (print (safe-div 10 2))  # prints 5
    nil
)
```

## Next Steps

1. **Explore Demos**: Run through all the demo examples
2. **Read Documentation**: Check out the [Language Reference](language-reference.md)
3. **Write Programs**: Create your own El programs
4. **Contribute**: Add new features or examples

## Getting Help

- Check the demo examples in `cmd/demo/`
- Read the language reference documentation
- Look at the source code in the repository
- Experiment with the interactive demos

## Tips for Success

1. **Start Simple**: Begin with basic examples and gradually increase complexity
2. **Use Demos**: The demo examples are your best learning resource
3. **Experiment**: Try modifying the examples to understand how they work
4. **Practice**: Write your own programs to reinforce learning
5. **Read Code**: Study the implementation to understand the language internals

Happy programming with El!
