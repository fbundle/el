package runtime_ext

import (
	"el/ast"
	"encoding/json"
	"errors"
	"strconv"

	"github.com/fbundle/lab_public/lab/go_util/pkg/adt"
)

// helpers
func value(o Object) adt.Result[Object] {
	return adt.Ok[Object](o)
}

func errValue(err error) adt.Result[Object] {
	return adt.Err[Object](err)
}

func errValueString(msg string) adt.Result[Object] {
	return errValue(errors.New(msg))
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

func parseLiteral(lit string) (Object, error) {
	if len(lit) == 0 {
		return nil, ErrorEmptyLiteral
	}
	if lit == ast.TokenUnwrap {
		return Unwrap{}, nil
	}
	if string(lit[0]) == ast.TokenStringBeg && string(lit[len(lit)-1]) == ast.TokenStringEnd {
		str := ""
		if err := json.Unmarshal([]byte(lit), &str); err != nil {
			return nil, err
		}
		return String{Val: str}, nil

	}
	i, err := strconv.Atoi(lit)
	return Int{Val: i}, err
}
