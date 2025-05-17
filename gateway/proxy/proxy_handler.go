package gateway

import (
	"fmt"
	"log/slog"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"

	"github.com/nabhdeep/gateway-cli/pkg/config"
)

// RunService is a function which creates routes from the gateway endpoint to proxy to the proxy server
func RunService(services_config *config.ServicesConfig, router *http.ServeMux) {
	for _, resource := range services_config.Services {

		// if enbaled in the config file
		if resource.Enabled {
			// parse url
			url, err := url.Parse(resource.Baseurl)
			proxy := NewProxy(url)
			if err != nil {
				// skips bad urls
				slog.Error("Bad BaseUrl Config file skipping service", slog.String("BaseUrl", resource.Baseurl))
				continue
			}
			// creates the router
			for _, routes := range resource.Routes {
				fmt.Println(routes.Method + resource.Service_Endpoint + routes.Endpoint)
				router.HandleFunc(routes.Method+resource.Service_Endpoint+routes.Endpoint, ProxyRequestHandler(proxy, url, resource.Service_Endpoint))
			}
		}
	}
}

func NewProxy(target *url.URL) *httputil.ReverseProxy {
	return httputil.NewSingleHostReverseProxy(target)
}

func ProxyRequestHandler(proxy *httputil.ReverseProxy, url *url.URL, endpoint string) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		slog.Info("=================")
		slog.Info("Request Received")
		slog.Info("", slog.String("Url", r.URL.String()))
		slog.Info("=================")
		r.URL.Host = url.Host
		r.URL.Scheme = url.Scheme
		r.Header.Set("X-Forwarded-Host", r.Header.Get("Host"))
		r.Host = url.Host
		path := r.URL.Path
		r.URL.Path = strings.TrimLeft(path, endpoint)
		slog.Info("Redirecting Request")
		slog.Info("", slog.String("Url", r.URL.String()))
		slog.Info("=================")
		proxy.ServeHTTP(w, r)
	}

}
