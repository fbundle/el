package runtime_ext

import (
	"el/pkg/el/runtime_core"
	"encoding/json"
	"errors"
	"strconv"

	"github.com/fbundle/lab_public/lab/go_util/pkg/adt"
)

type Runtime = runtime_core.Runtime

var InitRuntime = Runtime{
	MAX_STACK_DEPTH: 1000,
	ParseLiteralOpt: parseLiteralOpt,
	UnwrapArgsOpt:   unwrapArgsOpt,
}

var ErrorEmptyLiteral = errors.New("empty literal")

func unwrapArgsOpt(args []Object) adt.Option[[]Object] {
	return adt.Wrap(func() ([]Object, error) {
		return unwrapArgs(args)
	})()
}

func parseLiteralOpt(s string) adt.Option[Object] {
	return adt.Wrap(func() (Object, error) {
		return parseLiteral(s)
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
		return nil, ErrorEmptyLiteral
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
