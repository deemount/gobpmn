package common

import (
	"fmt"
	"runtime"
)

func NewError(err error) error {
	_, file, line, _ := runtime.Caller(1)
	return fmt.Errorf("\r[%s][%d]: %s", file, line, err)
}
