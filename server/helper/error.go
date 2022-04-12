package helper

import (
	"fmt"
)

func HandlePanicIfError(err error) {
	if err != nil {
		panic(err)
	}
}

func HandleError() {
	err := recover()
	if err != nil {
		fmt.Println("Ada error nih,", err)
	}
}
