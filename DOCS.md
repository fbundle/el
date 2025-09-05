### EL Language Specification

This document specifies the EL programming language as implemented in this repository.

### 1. Lexical Structure

- **Comments**: Start with `#` and run to end of line.
- **Whitespace**: Separates tokens; otherwise insignificant.
- **Strings**: Double-quoted JSON strings. Escape sequences follow JSON rules. Examples: "hello", "a\"b".
- **Numbers**: Signed base-10 integers (e.g., 0, -12, 42).
- **Names**: Any non-whitespace token that is not a special token; used for identifiers, operators, and keywords.
- **Special tokens**:
  - Parentheses: `(` `)` denote S-expressions.
  - Braces: `{` `}` introduce sugar expressions (infix and arrow/lambda sugar, and type casts).
  - Brackets: `[` `]` are syntactic sugar for `(list ...)`.
  - Dollar: `$` as a standalone token denotes the argument unwrapping operator.

### 2. Syntax

EL has two expression forms: names and S-expressions.

- **Name**: A single token such as `foo`, `123`, `"hi"`, `add`.
- **S-expression**: `(` command arg1 arg2 ... `)` where each element is an expression.

Sugar expressions in `{ ... }` are desugared into S-expressions per rules below.

#### 2.1. Core forms

- Function application: `(f a b c)` applies value `f` to arguments `a b c`.
- Let binding: `(let name1 expr1 name2 expr2 ... body)` binds names to values in a new scope, then evaluates `body`.
- Lambda: `(lambda p1 p2 ... body)` creates a closure with parameters `p1 p2 ...` and body `body`. Supports currying.
- Match: `(match cond v1 r1 v2 r2 ... default)` evaluates `cond`, compares with `v1`, `v2`, ... (by value and type). If equal, returns corresponding result; otherwise returns `default`.

#### 2.2. Sugar blocks `{ ... }`

Inside `{ ... }`, EL supports:

- **Arrow functions**: `{a b => expr}` => `(lambda a b expr)`.
- **Type casts**: `{value : type}` => `(type.cast type value)`.
- **Infix operators**:
  - Left-associative by default: `{a op b op c}` => `((op a b) c)`.
  - A special right-associative arrow for types: `{a -> b -> c}` => `(-> a (-> b c))` which maps to `(type_chain a b c)` by alias in templates.

Brackets are mapped prior to tokenization: `[e1 e2 ...]` => `(list e1 e2 ...)`.

### 3. Semantics

- **Evaluation model**: Call-by-value. Arguments to a function are evaluated before the call; the runtime may unwrap arguments (see `$` below) after evaluation.
- **Environment/Scopes**: A `Frame` maps names to values. Name resolution first checks current frame, then attempts to parse as literal (number, string, `$`).
- **Closure**: Lambdas capture the defining frame excluding parameter names. On full application, the call frame is merged into the closure for free variables; on partial application, a curried function is returned.
- **Equality in match**: Only comparable native data may be matched; type mismatch or non-comparable values cause error.
- **Errors/Interrupts**: Runtime checks for context cancellation and deadline to signal interruption or timeout.

### 4. Literals and Values

- **Integers**: e.g., `1`, `-3`. Type: `int_type`.
- **Strings**: JSON strings, e.g., `"hello"`. Type: `string_type`.
- **Booleans**: `true`, `false` bound in the base environment.
- **Lists**: `(list v1 v2 ...)` or `[v1 v2 ...]`. Type: `list_type`.
- **Nil/Unit**: `nil` is provided; the empty expression `()` evaluates to `nil`.

### 5. Builtins

Core builtins are available as names in the base frame.

- `let`: `(let name1 val1 ... body)`
- `lambda`: `(lambda p1 ... body)`
- `match`: `(match cond v1 r1 ... default)`
- `type_of`: `(type_of v)` returns the type of `v`.
- `type_cast`: `(type_cast type v)` casts value `v` to new type parent `type` if allowed.
- `type_chain`: `(type_chain t1 t2 ... tn)` constructs an arrow type `t1 -> t2 -> ... -> tn`.

List and utility extensions provided by the basic runtime:

- `list`: `(list a b ...)` create list.
- `len`: `(len list)` length.
- `slice`: `(slice list indices)` indices is a list of integers; returns a list of selected elements.
- `range`: `(range m n)` produce list `[m, m+1, ..., n-1]`.

Arithmetic and comparisons (integer-based):

- Arithmetic: `add`, `sub`, `mul`, `div`, `mod`.
- Comparisons: `eq`, `ne`, `lt`, `le`, `gt`, `ge`.

I/O utilities:

- `print`: `(print v1 v2 ...)` prints values, returns `nil`.
- `inspect`: `(inspect msg v1 v2 ...)` prints with types for debugging.

### 6. Operators and Aliases

The host program typically defines aliases via a template loaded before user code. Common aliases:

- `+ add`, `- sub`, `* mul`, `/ div`, `% mod`
- `== eq`, `!= ne`, `<= le`, `< lt`, `> gt`, `>= ge`
- `-> type_chain`

Through sugar blocks, infix expressions desugar accordingly, e.g., `{1 + 2 * 3}` => `(add 1 (mul 2 3))` with left-to-right grouping; there is no built-in operator precedence beyond the folding order, so use explicit grouping via braces if needed.

### 7. Unwrapping operator `$`

- `$` appearing as a name literal is parsed as an `Unwrap` marker by the basic runtime literal parser.
- After argument evaluation, the runtime repeatedly unwraps any pair `($ some_list)` into the elements of `some_list` as positional arguments.
- Examples:
  - `(print $[1 2 3])` spreads the list into arguments as if `(print 1 2 3)`.
  - Nested unwrapping is processed until no more unwraps are present.

### 8. Standard Library (provided in examples)

Common helpers defined in templates (see `cmd/basic/main.go`):

- Identity: `unit (lambda a a)`
- List helpers: `get`, `head`, `rest` using `slice` and `range`.
- `map`: recursive higher-order function over lists.
- `curry2`: converts a binary function into chained unary functions.

### 9. Types and Casting

- Types are objects of kind `builtin_type` with a sort system backed by arrows. Examples loaded in basic runtime: `int_type`, `string_type`, `list_type`.
- Functions have weakest arrow types by arity and can be cast to more specific arrow types using `type_chain` + `type_cast` when permitted.
- `{value : type}` sugar desugars to `(type.cast type value)`.

### 10. Examples

See `examples/` for end-to-end programs, including:
- `1_hello.el` – Hello world
- `2_basic_arithmetic.el` – Arithmetic and variables
- `3_lists_and_operations.el` – List primitives and HOFs
- `4_functions_and_lambdas.el` – Lambdas, recursion, closures
- `0_comprehensive_demo.el` – A comprehensive showcase

### 11. Error Cases

- Wrong arity for builtins yields runtime errors.
- `match` requires comparable, same-typed values.
- `type_cast` fails when the current parent sort is not less-equal to the target type.
- Unwrapping requires the next argument to be a list.

### 12. Implementation Notes

- AST forms: `Name` and `Lambda`.
- Parser: tokenizes with string-awareness; `{...}` sugar block handled by `processSugar` (arrow, type cast, and infix fold with special `->`).
- Runtime: evaluates names by frame lookup or literal parse; executes lambdas by looking up callable in head position; closures and currying supported.


