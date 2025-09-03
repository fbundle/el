package runtime_ext

import (
	"encoding/json"
	"errors"
	"strconv"

	"github.com/fbundle/lab_public/lab/go_util/pkg/adt"
)

// helpers
func value(o Value) adt.Result[Value] {
	return adt.Ok[Value](o)
}

func errValue(err error) adt.Result[Value] {
	return adt.Err[Value](err)
}

func errValueString(msg string) adt.Result[Value] {
	return errValue(errors.New(msg))
}

func unwrapArgs(args []Value) ([]Value, error) {
	var unwrapArgsLoop func(args []Value) ([]Value, bool, error)
	unwrapArgsLoop = func(args []Value) ([]Value, bool, error) {
		unwrapped := false
		unwrappedArgs := make([]Value, 0, len(args))
		for len(args) > 0 {
			head := args[0]
			if _, ok := head.(Unwrap); ok {
				if len(args) <= 1 {
					return unwrappedArgs, unwrapped, errors.New("unwrapping argument empty")
				}
				switch next := args[1].(type) {
				case List:
					unwrappedArgs = append(unwrappedArgs, next.Repr()...)
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

var ErrorEmptyLiteral = errors.New("empty literal")

func parseLiteral(lit string) (Value, error) {
	if len(lit) == 0 {
		return nil, ErrorEmptyLiteral
	}
	if lit == "*" {
		return Unwrap{}, nil
	}
	if lit[0] == '"' && lit[len(lit)-1] == '"' {
		str := ""
		if err := json.Unmarshal([]byte(lit), &str); err != nil {
			return nil, err
		}
		return String{str}, nil

	}
	i, err := strconv.Atoi(lit)
	return Int{i}, err
}
