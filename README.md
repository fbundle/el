### EL

An experimental, minimal Lisp-like language with sugar for infix and arrow lambdas, implemented in Go.

### Features

- S-expressions with `(let)`, `(lambda)`, `(match)`
- Sugar blocks `{ ... }` for infix arithmetic, comparisons, arrow lambdas `{a b => expr}`, and type casts `{v : type}`
- Lists via `[a b c]` or `(list a b c)` and common list helpers
- First-class functions, closures, and currying
- Simple type objects with cast and arrow type construction

### Getting Started

Prerequisites: Go 1.25+.

```bash
git clone https://github.com/your/repo.git
cd el
go run ./cmd/basic
```

This runs the demo program embedded in `cmd/basic/main.go`.

### Try the examples

Each example is a standalone program demonstrating features:

```text
examples/1_hello.el
examples/2_basic_arithmetic.el
examples/3_lists_and_operations.el
examples/4_functions_and_lambdas.el
examples/0_comprehensive_demo.el
```

To run an example, copy its contents into the `program` string in `cmd/basic/main.go` or adapt the launcher to read files.

### Language overview

Basic forms:

- Function call: `(f a b)`
- Let binding: `(let name1 val1 ... body)`
- Lambda: `(lambda p1 p2 ... body)`
- Match: `(match cond v1 r1 ... default)`

Sugar:

- Infix: `{1 + 2 * 3}` => `(add 1 (mul 2 3))` (left fold)
- Arrow lambda: `{a b => expr}` => `(lambda a b expr)`
- Type cast: `{v : type}` => `(type_cast type v)`
- Brackets: `[1 2]` => `(list 1 2)`
- Unwrap: `$` spreads list arguments: `(print $[1 2 3])`

### Builtins (selection)

- Core: `let`, `lambda`, `match`
- Types: `type_of`, `type_cast`, `type_chain`
- Lists: `list`, `len`, `slice`, `range`
- Math: `add`, `sub`, `mul`, `div`, `mod`
- Cmp: `eq`, `ne`, `lt`, `le`, `gt`, `ge`
- IO: `print`, `inspect`

More details in `DOCS.md`.

### Development

```bash
go test ./...   # if tests are added later
go vet ./...
```

### License

See `LICENSE`.


