package projectpath

import (
	"runtime"
	"strings"
)

func Root() string {

	_, file_path, _, _ := runtime.Caller(0)
	index := strings.Index(file_path, "/ops/projectpath")
	return file_path[:index]
}
