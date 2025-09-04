package ts

const (
	DefaultSortName = "type"
)

type Data interface {
	String() string
}

type Sort interface {
	Level() int
	String() string
	Parent() Sort
	Len() int
	LE(dst Sort) bool

	prepend(param Sort) Sort
}
