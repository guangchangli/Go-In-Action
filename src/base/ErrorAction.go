package main

import (
	"errors"
	"fmt"
)

func main() {

	println(errors.New("api exception"))
}

type apiException interface {
	Error() string
}
func BusinessException(format string,a ...interface{}) error{
	return errors.New(fmt.Sprintf(format,a...))
}
