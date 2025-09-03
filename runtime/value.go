package runtime

import (
	"context"
	"el/ast"
	"fmt"

	"github.com/fbundle/lab_public/lab/go_util/pkg/adt"
)

type Exec = func(r Runtime, ctx context.Context, s Stack, argList []ast.Node) adt.Result[Value]

type Value interface {
	Type() Type
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

func (t Type) Type() Type {
	return Type{
		level: t.level + 1,
	}
}

type Module struct {
	repr string
	exec Exec
}

func (m Module) Type() Type {
	return DataType("module")
}
func (m Module) String() string {
	return m.repr
}
