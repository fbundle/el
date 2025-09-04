package main

import (
	"context"
	"el/ast"
	"el/parser"
	"el/runtime"
	"el/runtime_ext"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

var (
	verbose     = flag.Bool("v", false, "verbose output - show parsed expressions")
	timeout     = flag.Duration("timeout", 30*time.Second, "execution timeout")
	showHelp    = flag.Bool("help", false, "show help message")
	showVersion = flag.Bool("version", false, "show version information")
	repl        = flag.Bool("repl", false, "start interactive REPL")
	debug       = flag.Bool("debug", false, "enable debug mode with detailed error information")
)

const version = "1.0.0"

func main() {
	flag.Parse()

	if *showHelp {
		showHelpMessage()
		return
	}

	if *showVersion {
		fmt.Printf("EL Programming Language Interpreter v%s\n", version)
		return
	}

	if *repl {
		startREPL()
		return
	}

	// Read from stdin or file
	var input io.Reader
	if len(flag.Args()) > 0 {
		file, err := os.Open(flag.Args()[0])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error opening file: %v\n", err)
			os.Exit(1)
		}
		defer file.Close()
		input = file
	} else {
		input = os.Stdin
	}

	// Read input
	content, err := io.ReadAll(input)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
		os.Exit(1)
	}

	// Execute the program
	if err := executeProgram(string(content)); err != nil {
		if *debug {
			fmt.Fprintf(os.Stderr, "Execution error: %+v\n", err)
		} else {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		}
		os.Exit(1)
	}
}

func executeProgram(program string) error {
	// Add template with common utilities
	programWithTemplate := withTemplate(program)

	// Tokenize
	tokens := parser.Tokenize(programWithTemplate)
	if *verbose {
		fmt.Printf("Tokens: %v\n", tokens)
	}

	// Create runtime
	r, s := runtime_ext.NewBasicRuntime()

	// Create context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), *timeout)
	defer cancel()

	var e ast.Expr
	var o runtime.Object
	var err error

	// Parse and execute each expression
	for len(tokens) > 0 {
		e, tokens, err = parser.Parse(tokens)
		if err != nil {
			return fmt.Errorf("parse error: %w", err)
		}

		if *verbose {
			fmt.Printf("Expression: %s\n", e)
		}

		if err := r.Step(ctx, s, e).Unwrap(&o); err != nil {
			return fmt.Errorf("execution error: %w", err)
		}

		if *verbose {
			fmt.Printf("Result: %s\n", o)
		} else {
			// Only print non-nil results in non-verbose mode
			if !isNil(o) {
				fmt.Println(o)
			}
		}
	}

	return nil
}

func startREPL() {
	fmt.Printf("EL Programming Language REPL v%s\n", version)
	fmt.Println("Type 'help' for commands, 'quit' or 'exit' to exit")
	fmt.Println()

	r, s := runtime_ext.NewBasicRuntime()

	ctx := context.Background()

	for {
		fmt.Print("el> ")

		var input string
		if _, err := fmt.Scanln(&input); err != nil {
			if err == io.EOF {
				fmt.Println("\nGoodbye!")
				return
			}
			fmt.Printf("Error reading input: %v\n", err)
			continue
		}

		input = strings.TrimSpace(input)
		if input == "" {
			continue
		}

		// Handle special commands
		switch input {
		case "quit", "exit":
			fmt.Println("Goodbye!")
			return
		case "help":
			showREPLHelp()
			continue
		case "clear":
			// Clear screen (works on most terminals)
			fmt.Print("\033[2J\033[H")
			continue
		}

		// Execute the input
		if err := executeREPLInput(ctx, r, s, input); err != nil {
			fmt.Printf("Error: %v\n", err)
		}
	}
}

func executeREPLInput(ctx context.Context, r runtime.Runtime, s runtime.Frame, input string) error {
	// Add template for REPL
	inputWithTemplate := withTemplate(input)

	tokens := parser.Tokenize(inputWithTemplate)
	if len(tokens) == 0 {
		return nil
	}

	var e ast.Expr
	var o runtime.Object
	var err error

	for len(tokens) > 0 {
		e, tokens, err = parser.Parse(tokens)
		if err != nil {
			return fmt.Errorf("parse error: %w", err)
		}

		if err := r.Step(ctx, s, e).Unwrap(&o); err != nil {
			return fmt.Errorf("execution error: %w", err)
		}

		// Print result
		if !isNil(o) {
			fmt.Printf("=> %s\n", o)
		}
	}

	return nil
}

func isNil(o runtime.Object) bool {
	_, ok := o.(runtime.Nil)
	return ok
}

func showHelpMessage() {
	fmt.Printf(`EL Programming Language Interpreter v%s

USAGE:
    go run cmd/el/main.go [OPTIONS] [FILE]

OPTIONS:
    -v              verbose output - show parsed expressions
    -timeout DUR    execution timeout (default: 30s)
    -max-stack N    maximum stack depth (default: 10000)
    -repl           start interactive REPL
    -debug          enable debug mode with detailed error information
    -help           show this help message
    -version        show version information

EXAMPLES:
    # Run a file
    go run cmd/el/main.go examples/1_hello.el
    
    # Run from stdin
    echo '(print "hello world")' | go run cmd/el/main.go
    
    # Start interactive REPL
    go run cmd/el/main.go -repl
    
    # Verbose execution
    go run cmd/el/main.go -v examples/demo1.el

DESCRIPTION:
    EL is a functional programming language with:
    - S-expression syntax
    - Lambda functions with closures
    - Pattern matching
    - List operations
    - Arithmetic and comparison operators
    - Tail-call optimization

For more information, see docs/README.md
`, version)
}

func showREPLHelp() {
	fmt.Println(`
REPL Commands:
    help            show this help message
    quit, exit      exit the REPL
    clear           clear the screen

REPL Examples:
    (print "hello world")
    (let x 42 (print x))
    (lambda x {x + 1})
    [1 2 3 4 5]
    (map [1 2 3] (lambda x {x * 2}))
`)
}

// withTemplate - add common utilities and functions to the code
func withTemplate(s string) string {
	return fmt.Sprintf(`
(let

# identity - identity function
unit (lambda x x) 

# get - get element from list
get (lambda l i (unit $(slice l (range i (add i 1)))))			# get l[i]
head (lambda l (get l 0))							# get l[0]
rest (lambda l (slice l (range 1 (len l))))			# get l[1:]
last (lambda l (get l (sub (len l) 1)))				# get last element
init (lambda l (slice l (range 0 (sub (len l) 1))))	# get all but last

# operator aliases
+ add - sub * mul / div %% mod
== eq != ne <= le < lt > gt >= ge

# conditional
if (lambda cond then else (match cond
	true then
	else
))

# list constructors
cons (lambda x xs (list x $xs))						# cons x xs
append (lambda xs ys (match (len xs)
	0 ys
	(let
		first (head xs)
		rest_xs (rest xs)
		rest_result (append rest_xs ys)
		(cons first rest_result)
	)
))

# map
map (lambda l f (match (len l)
	0 []					# if len l == 0 then return empty list
	(let
		first_elem (head l)
		first_elem2 (f first_elem)
		rest_elems (rest l)
		rest_elems2 (map rest_elems f)	# recursive call
		(cons first_elem2 rest_elems2)
	)
))

# filter
filter (lambda l pred (match (len l)
	0 []
	(let
		first_elem (head l)
		rest_elems (rest l)
		rest_filtered (filter rest_elems pred)
		(match (pred first_elem)
			true (cons first_elem rest_filtered)
			rest_filtered
		)
	)
))

# fold operations
foldl (lambda l init f (match (len l)
	0 init
	(let
		first_elem (head l)
		rest_elems (rest l)
		new_init (f init first_elem)
		(foldl rest_elems new_init f)
	)
))

# utility functions
sum (lambda l (foldl l 0 add))
product (lambda l (foldl l 1 mul))
max_list (lambda l (foldl l (head l) (lambda acc n (if {n > acc} n acc))))
min_list (lambda l (foldl l (head l) (lambda acc n (if {n < acc} n acc))))

# range generation
range (lambda start end (match {start >= end}
	true []
	(let
		rest_range (range {start + 1} end)
		(cons start rest_range)
	)
))

range_step (lambda start end step (match {start >= end}
	true []
	(let
		rest_range (range_step {start + step} end step)
		(cons start rest_range)
	)
))

# list utilities
reverse (lambda l (foldl l [] (lambda acc n (cons n acc))))
take (lambda n l (match {n <= 0}
	true []
	(match (len l)
		0 []
		(let
			first (head l)
			rest (rest l)
			rest_taken (take {n - 1} rest)
			(cons first rest_taken)
		)
	)
))

drop (lambda n l (match {n <= 0}
	true l
	(match (len l)
		0 []
		(drop {n - 1} (rest l))
	)
))

%s

)`, s)
}
