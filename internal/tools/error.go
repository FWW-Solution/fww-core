package tools

import (
	"errors"
	"fmt"
	"runtime"
)

func ErrorBuilder(err error) error {
	_, file, line, _ := runtime.Caller(1)
	errorMessage := fmt.Sprintf("%s | occurred in file %s at line %d", err.Error(), file, line)
	return errors.New(errorMessage)
}
