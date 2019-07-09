package util

import (
	"log"
)

func ChkErr(err error) {
	if err != nil {
		log.Println("Error:", err)
		panic(err)
	}
}
