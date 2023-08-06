package cl

import (
	"fmt"
)

func WriteErr(e error) {
	panic(e.Error())
}

func Write(a ...any) {
	fmt.Println(a)
}
