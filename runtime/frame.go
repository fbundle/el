package runtime

import (
	"github.com/fbundle/lab_public/lab/go_util/pkg/persistent/ordered_map"
)

type Name string

type Frame = ordered_map.OrderedMap[Name, Object]
