package util

import (
	"log"
)

// Gin style shortcut
type H map[string]interface{}

func ChkErr(err error) {
	if err != nil {
		log.Println("Error:", err)
		panic(err)
	}
}
