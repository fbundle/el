package runtime_ext

import (
	"el/runtime"
	"fmt"

	"github.com/fbundle/lab_public/lab/go_util/pkg/adt"
)

// helpers
// result helpers
func resultObj(o Object) adt.Result[Object] {
	return adt.Ok(o)
}

func resultTypedData(data TypedData) adt.Result[Object] {
	return adt.Ok(makeTypedData(data))
}

func resultErr(err error) adt.Result[Object] {
	return adt.Err[Object](err)
}
func resultErrStrf(format string, args ...any) adt.Result[Object] {
	return adt.Err[Object](fmt.Errorf(format, args...))
}
func makeTypedData(data TypedData) Object {
	return runtime.MakeData(data, runtime.MakeType(data.TypeName()))
}
