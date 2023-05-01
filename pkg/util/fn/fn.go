package fn

import (
	"path/filepath"
	"runtime"
)

func Name(skips ...int) string {
	skip := 1
	if len(skips) > 0 {
		skip = skips[0]
	}
	pc, _, _, _ := runtime.Caller(skip)
	return filepath.Base(runtime.FuncForPC(pc).Name())
}
