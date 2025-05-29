package middleware

import "net/http"

// Reference
// https://dev.to/karankumarshreds/middlewares-in-go-41j
type Middleware func(http.HandlerFunc) http.HandlerFunc
type Chain []Middleware

// To chain middleware in Go, each middleware function should take an http.Handler
// as an argument and return an http.Handler.
// Middleware is applied from the outermost to the innermost function.
// So the order of execution is: m1(m2(m3(originalHandler)))
// This means m1 runs first, and m3 runs just before the original handler.

// CreateMiddlewareChain will take middleware as arguments and at last then func will work to run the original handler
func CreateMiddlewareChain(middlewares ...Middleware) Chain {
	return Chain(middlewares)
}

func (c Chain) Then(originalHandler http.HandlerFunc) http.HandlerFunc {
	for i := len(c) - 1; i >= 0; i-- {
		originalHandler = c[i](originalHandler)
	}
	return originalHandler
}
