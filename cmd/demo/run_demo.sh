#!/bin/bash

# El Language Demo Runner Script
# This script provides an easy way to run individual demos

if [ $# -eq 0 ]; then
    echo "Usage: $0 <demo-name>"
    echo ""
    echo "Available demos:"
    echo "  hello       - Hello World"
    echo "  arithmetic  - Basic arithmetic operations"
    echo "  functions   - Function definitions and calls"
    echo "  lists       - List operations"
    echo "  recursion   - Recursive functions"
    echo "  closures    - Closure examples"
    echo "  matching    - Pattern matching"
    echo "  types       - Type system examples"
    echo "  advanced    - Advanced features"
    echo "  algorithms  - Algorithm implementations"
    echo "  data_structures - Data structure examples"
    echo "  math_problems - Mathematical problems"
    echo "  all         - Run all demos"
    echo ""
    echo "Examples:"
    echo "  $0 hello"
    echo "  $0 functions"
    echo "  $0 all"
    exit 1
fi

DEMO_NAME=$1

# Check if we're in the right directory
if [ ! -f "cmd/demo/main.go" ]; then
    echo "Error: Please run this script from the project root directory"
    exit 1
fi

# Run the demo
echo "Running El demo: $DEMO_NAME"
echo "=================================="
go run cmd/demo/main.go "$DEMO_NAME"
