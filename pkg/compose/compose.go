package compose

import (
	"net/http"
)

type MiddlewareFn func(http.HandlerFunc) http.HandlerFunc

func Compose(mdws ...MiddlewareFn) MiddlewareFn {
	return func(next http.HandlerFunc) http.HandlerFunc {
		for i := 0; i < len(mdws); i++ {
			mdw := mdws[i]
			next = mdw(next)
		}
		return next
	}
}
