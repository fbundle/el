package runtime_ext

import (
	"encoding/json"
	"errors"
	"strconv"

	"github.com/fbundle/lab_public/lab/go_util/pkg/persistent/seq"
)

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
					lst := seq.Seq[Object](next)
					unwrappedArgs = append(unwrappedArgs, lst.Repr()...)
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
		strList := seq.Seq[Object]{}
		for _, ch := range []rune(str) {
			strList = strList.Ins(strList.Len(), Int(ch))
		}
		return List(strList), nil
	}
	i, err := strconv.Atoi(lit)
	return Int(i), err
}
