package middleware

import "net/http"

// Template Middleware is one you can copy to
// create your own middleware and same can be added to the chain method
// CreateMiddlewareChain(m1 , m2 , m3)
func template() Middleware {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			// handle the middleware here
			next(w, r)
		}
	}
}
