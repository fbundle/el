package el

import (
	"encoding/json"
	"errors"
	"strconv"
)

// NewCoreRuntime - Runtime + core control flow extensions
func NewCoreRuntime() *Runtime {
	return (&Runtime{
		ParseLiteral: func(lit string) (Object, error) {
			if len(lit) == 0 {
				return nil, errors.New("empty literal")
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
		},
		Stack: newFrameStack(),
	}).LoadModule(letModule, lambdaModule, matchModule)
}

// NewBasicRuntime - NewCoreRuntime + minimal set of arithmetic extensions for Turing completeness
func NewBasicRuntime() *Runtime {
	return NewCoreRuntime().
		// list extension
		LoadExtension(listExtension, lenExtension, rangeExtension, sliceExtension, unitExtension).
		// compare extension
		LoadConstant("true", Int(1)).LoadConstant("false", Int(0)).
		LoadExtension(eqExtension, neExtension, ltExtension, leExtension, gtExtension, geExtension).
		// arithmetic extension
		LoadExtension(addExtension, subExtension, mulExtension, divExtension, modExtension).
		// extra
		LoadExtension(ifExtension)
}
