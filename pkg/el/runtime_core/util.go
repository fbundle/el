package runtime

import (
	"el/pkg/el/expr"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	"github.com/fbundle/lab_public/lab/go_util/pkg/adt"
)

const (
	MAX_STACK_DEPTH = 1000
)

var NameNotFoundError = func(name Name) error {
	return fmt.Errorf("object not found %s", name)
}
var InterruptError = func(err error) error {
	return fmt.Errorf("interrupted: %s", err)
}
var TimeoutError = func(err error) error {
	return fmt.Errorf("timeout: %s", err)
}
var StackOverflowError = errors.New("stack overflow")
var EmptyLiteralError = errors.New("empty literal")
var UnknownExpression = func(e expr.Expr) error {
	return fmt.Errorf("unknown expression type %s", e.String())
}
var CannotExecuteExpression = func(e expr.Expr) error {
	return fmt.Errorf("expression cannot be executed: %s", e.String())
}
var NotEnoughArguments = errors.New("not enough arguments")
var InternalError = errors.New("internal")

func unwrapArgsOpt(args []Object) adt.Option[[]Object] {
	return adt.Wrap(func() ([]Object, error) {
		return unwrapArgs(args)
	})()
}

func unwrapArgs(args []Object) ([]Object, error) {
	var unwrapArgsLoop func(args []Object) ([]Object, bool, error)
	unwrapArgsLoop = func(args []Object) ([]Object, bool, error) {
		unwrapped := false
		unwrappedArgs := make([]Object, 0, len(args))
		for len(args) > 0 {
			head := args[0]
			if _, ok := head.(Unwrap); ok {
				if len(args) <= 1 {
					return unwrappedArgs, unwrapped, errors.New("unwrapping argument empty")
				}
				switch next := args[1].(type) {
				case List:
					unwrappedArgs = append(unwrappedArgs, next...)
					args = args[2:]
					unwrapped = true
				case Unwrap: // nested unwrap
					unwrappedArgs = append(unwrappedArgs, head)
					args = args[1:]
				default:
					return unwrappedArgs, unwrapped, errors.New("unwrapping argument must be a list or an unwrap")
				}
			} else {
				unwrappedArgs = append(unwrappedArgs, head)
				args = args[1:]
			}
		}
		return unwrappedArgs, unwrapped, nil
	}
	var unwrapped bool
	var err error
	for { // keep unwrapping
		args, unwrapped, err = unwrapArgsLoop(args)
		if err != nil {
			return nil, err
		}
		if !unwrapped {
			return args, nil
		}
	}
}

func parseLiteral(lit string) (Object, error) {
	if len(lit) == 0 {
		return nil, EmptyLiteralError
	}
	if lit == "_" {
		return Wildcard{}, nil
	}
	if lit == "*" {
		return Unwrap{}, nil
	}
	if lit[0] == '"' && lit[len(lit)-1] == '"' {
		str := ""
		if err := json.Unmarshal([]byte(lit), &str); err != nil {
			return nil, err
		}
		strList := List{}
		for _, ch := range []rune(str) {
			strList = append(strList, Int(ch))
		}
		return strList, nil
	}
	i, err := strconv.Atoi(lit)
	return Int(i), err
}

func searchOnStack(s Stack, name Name) (Object, bool) {
	for _, frame := range s.Iter {
		if o, ok := frame.Get(name); ok {
			return o, true
		}
	}
	return nil, false
}

func object(o Object) adt.Option[Object] {
	return adt.Some[Object](o)
}

func errorObject(err error) adt.Option[Object] {
	return adt.Error[Object](err)
}

func errorObjectString(msg string) adt.Option[Object] {
	return errorObject(errors.New(msg))
}
