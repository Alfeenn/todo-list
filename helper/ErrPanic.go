package helper

import "log"

func PanicIfErr(err error) {
	if err != nil {
		log.Print("error helper")
		panic(err)
	}
}
