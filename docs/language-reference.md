# El Language Reference

## Table of Contents
1. [Introduction](#introduction)
2. [Lexical Structure](#lexical-structure)
3. [Data Types](#data-types)
4. [Expressions](#expressions)
5. [Built-in Functions](#built-in-functions)
6. [Control Flow](#control-flow)
7. [Functions and Closures](#functions-and-closures)
8. [Advanced Features](#advanced-features)
9. [Error Handling](#error-handling)

## Introduction

El is a functional programming language with Lisp-like syntax. It features:
- Dynamic typing
- Lexical scoping with closures
- Tail recursion optimization
- Pattern matching
- First-class functions
- Immutable data structures

## Lexical Structure

### Tokens

#### Literals
- **Integers**: `123`, `-456`, `0`
- **Strings**: `"hello world"`, `"with \"escaped\" quotes"`
- **Booleans**: `true`, `false`
- **Nil**: `nil`

#### Identifiers
- Start with letter or underscore
- Can contain letters, digits, underscores, hyphens
- Examples: `x`, `my-var`, `_private`, `func123`

#### Special Characters
- `(` `)` - Parentheses for expressions
- `[` `]` - Square brackets for list literals
- `{` `}` - Curly braces for infix expressions
- `"` - String delimiters
- `#` - Comment prefix
- `*` - Unwrap operator

#### Comments
```lisp
# This is a single-line comment
(let x 42  # inline comment
    x)
```

## Data Types

### Primitive Types

#### Integer
- 64-bit signed integers
- Examples: `1`, `-42`, `0`

#### String
- UTF-8 encoded strings
- Escape sequences: `\"`, `\\`, `\n`, `\t`
- Examples: `"hello"`, `"line1\nline2"`

#### Boolean
- `true` and `false`
- Used in conditionals and comparisons

#### Nil
- Represents absence of value
- Type: `nil`

### Composite Types

#### List
- Immutable sequences
- Heterogeneous (can contain different types)
- Examples: `[1 2 3]`, `["a" 42 true]`

#### Function
- First-class values
- Can be stored in variables, passed as arguments
- Examples: `(lambda x {x + 1})`, `add`

#### Type
- Runtime type information
- Obtained with `(type value)`
- Examples: `"int"`, `"list"`, `"string"`

## Expressions

### Basic Expressions

#### Literals
```lisp
42                    # integer
"hello"              # string
true                 # boolean
nil                  # nil value
```

#### Variable References
```lisp
x                    # reference to variable x
my-var               # reference to variable my-var
```

#### Function Calls
```lisp
(function arg1 arg2)  # call function with arguments
(add 1 2)            # call add with 1 and 2
```

### Special Syntax

#### Infix Arithmetic
```lisp
{1 + 2 * 3}          # equivalent to (add 1 (mul 2 3))
{x - y / z}          # equivalent to (sub x (div y z))
```

#### Arrow Functions
```lisp
{x y => {x + y}}     # equivalent to (lambda x y (add x y))
{n => {n * n}}       # equivalent to (lambda n (mul n n))
```

#### List Literals
```lisp
[1 2 3]              # equivalent to (list 1 2 3)
["a" "b" "c"]        # equivalent to (list "a" "b" "c")
```

#### Unwrap Operator
```lisp
*[1 2 3]             # unwraps list as arguments: 1 2 3
(add *[1 2])         # equivalent to (add 1 2)
```

## Built-in Functions

### Arithmetic Functions

#### add
```lisp
(add a b ...)        # Addition
(add 1 2 3)         # => 6
```

#### sub
```lisp
(sub a b ...)        # Subtraction
(sub 10 3 2)        # => 5
```

#### mul
```lisp
(mul a b ...)        # Multiplication
(mul 2 3 4)         # => 24
```

#### div
```lisp
(div a b ...)        # Division (integer)
(div 20 4 2)        # => 2
```

#### mod
```lisp
(mod a b)            # Modulo
(mod 17 5)          # => 2
```

### Comparison Functions

#### eq
```lisp
(eq a b)             # Equality
(eq 5 5)            # => true
(eq "a" "b")        # => false
```

#### ne
```lisp
(ne a b)             # Not equal
(ne 1 2)            # => true
```

#### lt, le, gt, ge
```lisp
(lt a b)             # Less than
(le a b)             # Less than or equal
(gt a b)             # Greater than
(ge a b)             # Greater than or equal
```

### List Functions

#### list
```lisp
(list a b ...)       # Create list
(list 1 2 3)        # => [1 2 3]
```

#### len
```lisp
(len list)           # Get length
(len [1 2 3])       # => 3
```

#### slice
```lisp
(slice list indices) # Get elements by indices
(slice [1 2 3 4] [0 2])  # => [1 3]
```

#### range
```lisp
(range start end)    # Create range
(range 0 5)         # => [0 1 2 3 4]
```

### I/O Functions

#### print
```lisp
(print value ...)    # Print values
(print "hello" 42)  # prints: hello 42
```

## Control Flow

### Variable Binding (let)

```lisp
(let var1 val1 var2 val2 ... result)
```

Creates local bindings and evaluates the result expression.

```lisp
(let
    x 10
    y 20
    {x + y}          # => 30
)
```

### Pattern Matching (match)

```lisp
(match value pattern1 result1 pattern2 result2 ... default)
```

Evaluates value and returns the result of the first matching pattern.

```lisp
(match 3
    1 "one"
    2 "two"
    3 "three"        # => "three"
    "other"
)
```

## Functions and Closures

### Lambda Functions

```lisp
(lambda param1 param2 ... body)
```

Creates a function with the given parameters and body.

```lisp
(lambda x {x * 2})                    # double function
(lambda x y {x + y})                  # add function
(lambda x (print x))                  # print function
```

### Function Application

```lisp
(function arg1 arg2 ...)
```

Calls a function with the given arguments.

```lisp
((lambda x {x + 1}) 5)               # => 6
(add 1 2)                            # => 3
```

### Closures

Functions capture their lexical environment:

```lisp
(let
    make-adder (lambda n (lambda x {x + n}))
    add-5 (make-adder 5)
    (add-5 10)                        # => 15
)
```

### Recursive Functions

```lisp
(let
    factorial (lambda n (match {n <= 1}
        true 1
        {n * factorial {n - 1}}
    ))
    (factorial 5)                     # => 120
)
```

## Advanced Features

### Type Introspection

```lisp
(type value)         # Get type of value
(type 42)           # => "int"
(type "hello")      # => "string"
(type [1 2 3])      # => "list"
```

### Higher-Order Functions

```lisp
(let
    apply-twice (lambda f x (f (f x)))
    increment (lambda x {x + 1})
    (apply-twice increment 5)         # => 7
)
```

### Complex Data Structures

```lisp
(let
    person [name "Alice" age 30 city "NYC"]
    name (slice person [0 1])         # => [name "Alice"]
    age (slice person [2 3])          # => [age 30]
    [name age]
)
```

### Mutual Recursion

```lisp
(let
    even (lambda n (match {n <= 0}
        true true
        (odd {n - 1})
    ))
    odd (lambda n (match {n <= 0}
        true false
        (even {n - 1})
    ))
    [(even 4) (odd 4)]                # => [true false]
)
```

## Error Handling

### Common Errors

#### Name Not Found
```lisp
undefined-var        # Error: object not found undefined-var
```

#### Type Mismatch
```lisp
(add "hello" 42)     # Error: add argument must be an integer
```

#### Wrong Number of Arguments
```lisp
(eq 1)               # Error: eq requires 2 arguments
```

#### Division by Zero
```lisp
(div 10 0)           # Error: division by zero
```

### Error Prevention

Always check types and argument counts:

```lisp
(let
    safe-div (lambda a b (match b
        0 "division by zero"
        (div a b)
    ))
    (safe-div 10 0)                   # => "division by zero"
    (safe-div 10 2)                   # => 5
)
```

## Best Practices

### Code Organization

1. **Use meaningful variable names**
2. **Break complex expressions into smaller parts**
3. **Use comments to explain complex logic**
4. **Prefer pure functions when possible**

### Performance

1. **Use tail recursion for iterative algorithms**
2. **Avoid deep nesting of function calls**
3. **Use pattern matching instead of nested conditionals**

### Style

1. **Indent consistently (2 or 4 spaces)**
2. **Use line breaks for readability**
3. **Group related expressions together**

## Examples

### Sorting Algorithm
```lisp
(let
    quicksort (lambda lst (match (len lst)
        0 []
        1 lst
        (let
            pivot (slice lst [0])
            rest (slice lst (range 1 (len lst)))
            smaller (filter rest (lambda x {x < *pivot}))
            larger (filter rest (lambda x {x >= *pivot}))
            [*quicksort smaller *pivot *quicksort larger]
        )
    ))
    (quicksort [3 1 4 1 5 9 2 6])
)
```

### Map and Filter
```lisp
(let
    map (lambda lst f (match (len lst)
        0 []
        [f *slice lst [0] *map (slice lst (range 1 (len lst))) f]
    ))
    filter (lambda lst pred (match (len lst)
        0 []
        (let
            head (slice lst [0])
            tail (slice lst (range 1 (len lst)))
            (match (pred *head)
                true [*head *filter tail pred]
                (filter tail pred)
            )
        )
    ))
    numbers [1 2 3 4 5 6 7 8 9 10]
    evens (filter numbers (lambda x {x % 2 == 0}))
    squares (map evens (lambda x {x * x}))
    squares
)
```
