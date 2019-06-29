package compose

import (
	"log"
	"net/http"
)

type MiddlewareFn func(next http.HandlerFunc) http.HandlerFunc

func Compose(mdws ...interface{}) http.HandlerFunc {
	last := mdws[len(mdws)-1]
	log.Printf("xxxx.. %s", last)
	fn, ok := last.(*http.HandlerFunc)

	if !ok {
		panic("last element must be Handler fn")
	}

	var tmp = *fn
	for i := 0; i < len(mdws)-1; i++ {
		mdw, ok := mdws[i].(MiddlewareFn)
		if !ok {
			panic("!!")
		}
		tmp = mdw(tmp)
	}

	return tmp
}
