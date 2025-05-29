package middleware

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"sync"
	"time"

	"github.com/nabhdeep/gateway-cli/pkg/constants"
	"golang.org/x/time/rate"
)

type client struct {
	limiter  *rate.Limiter
	lastSeen time.Time
}

var (
	mu      sync.Mutex
	clients = make(map[string]*client)
)

// RateLimiterMiddleware applies rate limtationson bases on IP
// newLimiter first arg is tokens added per second to the bucket, second arg is how many tokens a buket can store
// each req takes 1 token,
func RateLimiterMiddleware(perSecond int) Middleware {
	return func(next http.HandlerFunc) http.HandlerFunc {
		go func() {
			for {
				time.Sleep(time.Second)
				mu.Lock()
				for ip, c := range clients {
					if time.Since(c.lastSeen) > 3*time.Minute {
						delete(clients, ip)
					}
				}
				mu.Unlock()
				fmt.Println(clients)
			}
		}()
		return func(w http.ResponseWriter, r *http.Request) {
			ip, _, err := net.SplitHostPort(r.RemoteAddr)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			mu.Lock()
			if _, found := clients[ip]; !found {
				clients[ip] = &client{
					limiter: rate.NewLimiter(rate.Limit(perSecond), perSecond),
				}
			}
			clients[ip].lastSeen = time.Now()

			if !clients[ip].limiter.Allow() {
				mu.Unlock()
				w.Header().Set("Content-Type", "application/json; charset=utf-8")
				w.WriteHeader(http.StatusTooManyRequests)
				json.NewEncoder(w).Encode(constants.Message{
					Success: false,
					Message: "too many requests, try again later.",
				})
				return
			}
			mu.Unlock()

			// Continue to the next handler
			next(w, r)
		}
	}
}

// Implement Redis based rate limiter
