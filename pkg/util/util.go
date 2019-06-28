package util

import "log"

func ChkErr(err error) {
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
}
