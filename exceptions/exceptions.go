package exceptions

import (
	"fmt"
)

func ErrorFy(err error) {
	if err != nil {
		panic(fmt.Errorf(err.Error()))
	}
}
