package middleware

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/nabhdeep/gateway-cli/pkg/config"
	"github.com/nabhdeep/gateway-cli/pkg/constants"
)

// Whitelist_ip is a fuction which checks if the incoming request from
// the remote address is whitelisted with the service (In service_confign file)
func Whitelist_ip(resouce config.Service) Middleware {
	var ip_map = make(map[string](bool))
	for _, ip := range resouce.Allow_List {
		ip_map[ip] = true
	}
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			incoming_ip := strings.Split(r.RemoteAddr, ":")[0]
			if !ip_map[incoming_ip] {
				w.Header().Set("Content-Type", "application/json; charset=utf-8")
				w.WriteHeader(http.StatusUnauthorized)
				json.NewEncoder(w).Encode(constants.Message{
					Success: false,
					Message: "unauthorized",
				})
				return
			}
			next(w, r)
		}
	}
}
