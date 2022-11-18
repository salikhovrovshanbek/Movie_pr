package neccessaryFuncs

import "log"

func CheckERR(err error) {
	if err != nil {
		log.Println("error:", err.Error())
		return
	}
}
