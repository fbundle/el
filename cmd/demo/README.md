# El Language Demos

This directory contains comprehensive demo examples showcasing the features and capabilities of the El programming language.

## Running Demos

### Using the Demo Runner
```bash
# Run all demos
go run cmd/demo/main.go all

# Run a specific demo
go run cmd/demo/main.go hello
go run cmd/demo/main.go arithmetic
go run cmd/demo/main.go functions
# ... etc
```

### Available Demos

1. **hello** - Hello World and basic output
2. **arithmetic** - Basic arithmetic operations and comparisons
3. **functions** - Function definitions, calls, and higher-order functions
4. **lists** - List operations, creation, and manipulation
5. **recursion** - Recursive functions and algorithms
6. **closures** - Closure behavior and lexical scoping
7. **matching** - Pattern matching and conditional logic
8. **types** - Type system and introspection
9. **advanced** - Advanced features and complex examples
10. **algorithms** - Common algorithms (sorting, searching, etc.)
11. **data_structures** - Data structure implementations
12. **math_problems** - Mathematical problems and solutions

### Individual Demo Files

Each demo is also available as a standalone `.el` file:

- `01_hello_world.el` - Basic "Hello World" program
- `02_arithmetic.el` - Arithmetic operations and expressions
- `03_functions.el` - Function definitions and calls
- `04_lists.el` - List operations and manipulation
- `05_recursion.el` - Recursive algorithms
- `06_closures.el` - Closure examples
- `07_pattern_matching.el` - Pattern matching examples
- `08_types.el` - Type system demonstrations
- `09_advanced.el` - Advanced language features
- `10_algorithms.el` - Algorithm implementations
- `11_data_structures.el` - Data structure examples
- `12_math_problems.el` - Mathematical problem solutions

## Demo Categories

### Basic Features
- **Hello World**: Introduction to the language
- **Arithmetic**: Basic math operations and expressions
- **Functions**: Function definitions and calls
- **Lists**: List operations and data structures

### Intermediate Features
- **Recursion**: Recursive functions and algorithms
- **Closures**: Lexical scoping and closure behavior
- **Pattern Matching**: Conditional logic and pattern matching
- **Types**: Type system and introspection

### Advanced Features
- **Advanced**: Complex examples and advanced language features
- **Algorithms**: Common algorithms and data structures
- **Data Structures**: Custom data structure implementations
- **Math Problems**: Mathematical problem solving

## Learning Path

For beginners, we recommend following this order:

1. Start with `hello` to understand basic syntax
2. Try `arithmetic` to learn expressions and operators
3. Explore `functions` to understand function definitions
4. Learn `lists` for data manipulation
5. Practice `recursion` for algorithmic thinking
6. Study `closures` for advanced function concepts
7. Use `matching` for conditional logic
8. Explore `types` for type system understanding
9. Try `advanced` for complex examples
10. Study `algorithms` and `data_structures` for practical applications
11. Solve `math_problems` for mathematical programming

## Example Usage

### Running a Single Demo
```bash
go run cmd/demo/main.go functions
```

### Running All Demos
```bash
go run cmd/demo/main.go all
```

### Viewing Demo Source Code
```bash
cat cmd/demo/03_functions.el
```

## Customizing Demos

You can modify any of the `.el` files to experiment with the language. The demo runner will execute your modified code.

## Contributing

Feel free to add new demo examples by:
1. Creating a new `.el` file with a descriptive name
2. Adding the demo to the main demo runner
3. Updating this README with the new demo description

## Notes

- All demos are designed to be educational and demonstrate language features
- Some algorithms are simplified for demonstration purposes
- The language is still evolving, so some features may change
- Comments in the code explain the concepts being demonstrated
