package runtime_ext

import (
	"el/ast"
	"el/runtime"
	"encoding/json"
	"errors"
	"strconv"

	"github.com/fbundle/lab_public/lab/go_util/pkg/adt"
)

type Runtime = runtime.Runtime
type Frame = runtime.Frame

func NewBasicRuntime() (Runtime, Frame) {
	r := Runtime{
		ParseLiteral: func(lit string) adt.Result[Object] {
			val, err := parseLiteral(lit)
			return adt.Result[Object]{
				Val: makeTypedData(val),
				Err: err,
			}
		},
		UnwrapArgs: func(argsOpt adt.Result[[]Object]) adt.Result[[]Object] {
			var args []Object
			if err := argsOpt.Unwrap(&args); err != nil {
				return adt.Result[[]Object]{
					Err: err,
				}
			}
			unwrappedArgs, err := unwrapArgs(args)
			return adt.Result[[]Object]{
				Val: unwrappedArgs,
				Err: err,
			}
		},
	}
	f :=
		(&frameHelper{frame: runtime.Builtin}).
			LoadExtension(listExtension, lenExtension, sliceExtension, rangeExtension).
			Load("true", makeTypedData(True)).Load("false", makeTypedData(False)).
			Load("int_type", runtime.MakeType("int")).
			Load("list_type", runtime.MakeType("list")).
			Load("string_type", runtime.MakeType("string")).
			Load("names", runtime.MakeData(namesFunc, runtime.BuiltinType)).
			LoadExtension(eqExtension, neExtension, ltExtension, leExtension, gtExtension, geExtension).
			LoadExtension(addExtension, subExtension, mulExtension, divExtension, modExtension).
			LoadExtension(printExtension, inspectExtension)

	return r, f.frame
}

type frameHelper struct {
	frame Frame
}

func (sh *frameHelper) Load(name Name, value Object) *frameHelper {
	sh.frame = sh.frame.Set(name, value)
	return sh
}

func (sh *frameHelper) LoadExtension(exts ...Extension) *frameHelper {
	for _, ext := range exts {
		sh.Load(ext.Name, runtime.MakeData(ext.Module(), runtime.BuiltinType))
	}
	return sh
}
func unwrapArgs(args []Object) ([]Object, error) {
	var unwrapArgsLoop func(args []Object) ([]Object, bool, error)
	unwrapArgsLoop = func(args []Object) ([]Object, bool, error) {
		unwrapped := false
		unwrappedArgs := make([]Object, 0, len(args))
		for len(args) > 0 {
			head := args[0]
			if _, ok := head.Data().(Unwrap); ok {
				if len(args) <= 1 {
					return unwrappedArgs, unwrapped, errors.New("unwrapping argument empty")
				}
				switch next := args[1].Data().(type) {
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

func parseLiteral(lit string) (TypedData, error) {
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
