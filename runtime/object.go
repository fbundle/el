package runtime

import (
	"context"
	"el/ast"
	"fmt"

	"github.com/fbundle/lab_public/lab/go_util/pkg/adt"
)

type Exec = func(r Runtime, ctx context.Context, s Stack, argList []ast.Expr) adt.Result[Object]

type Object interface {
	Type() Object
	String() string
}

func DataType(name string) Type {
	return Type{
		level: 0,
		name:  name,
	}
}

type Type struct {
	level int
	name  string
}

func (t Type) String() string {
	if t.level == 0 {
		return t.name
	}
	return fmt.Sprintf("type_%d", t.level)
}

func (t Type) Type() Object {
	return Type{
		level: t.level + 1,
	}
}

type Function struct {
	exec Exec
	repr string
}

func (f Function) Type() Object {
	return DataType("function")
}
func (f Function) String() string {
	return f.repr
}
