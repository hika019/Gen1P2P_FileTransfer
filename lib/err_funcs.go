package lib

import (
	"fmt"
	"os"
)

func CheckErrExit(err error) {
	if err != nil {
		fmt.Println("Fail: ", err.Error())
		os.Exit(1)
	}
}

func CheckErr(err error) bool {
	if err != nil {
		fmt.Println("Fail: ", err.Error())
		return false
	}
	return true
}
