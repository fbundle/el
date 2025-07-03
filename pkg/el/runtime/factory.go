package runtime

import (
	"el/pkg/el/obj"
	"encoding/json"
	"errors"
	"strconv"
)

// NewCoreRuntime - Runtime and core control flow extensions
func NewCoreRuntime() *Runtime {
	return (&Runtime{
		ParseLiteral: func(lit string) (obj.Object, error) {
			if len(lit) == 0 {
				return nil, errors.New("empty literal")
			}
			if lit == "_" {
				return obj.Wildcard{}, nil
			}
			if lit == "*" {
				return obj.Unwrap{}, nil
			}
			if lit[0] == '"' && lit[len(lit)-1] == '"' {
				str := ""
				if err := json.Unmarshal([]byte(lit), &str); err != nil {
					return nil, err
				}
				strList := obj.List{}
				for _, ch := range []rune(str) {
					strList = append(strList, obj.Int(ch))
				}
				return strList, nil
			}
			i, err := strconv.Atoi(lit)
			return obj.Int(i), err
		},
		Stack: obj.NewFrameStack(),
	}).LoadModule(letModule, lambdaModule, matchModule)
}

// NewBasicRuntime - NewCoreRuntime and minimal set of arithmetic extensions for Turing completeness
func NewBasicRuntime() *Runtime {
	return NewCoreRuntime().
		// list extension
		LoadExtension(listExtension, lenExtension, rangeExtension, sliceExtension).
		// arithmetic extension
		LoadConstant("true", obj.True).LoadConstant("false", obj.False).
		LoadExtension(eqExtension, neExtension, ltExtension, leExtension, gtExtension, geExtension).
		LoadExtension(addExtension, subExtension, mulExtension, divExtension, modExtension)
	// extra
}
