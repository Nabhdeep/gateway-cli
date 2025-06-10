package proxy

import (
	"fmt"
	"log/slog"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"

	"github.com/nabhdeep/gateway-cli/pkg/config"
	"github.com/nabhdeep/gateway-cli/pkg/gateway/middleware"
)

type ServiceProxy struct {
	Proxy      *httputil.ReverseProxy
	Middleware middleware.Chain
	TargetUrl  *url.URL
	Endpoint   string
}

func NewService(resource config.Service) (*ServiceProxy, error) {
	targetUrl, err := url.Parse(resource.Baseurl)
	if err != nil {
		// skips bad urls
		slog.Error("Bad BaseUrl Config file skipping service", slog.String("BaseUrl", resource.Baseurl))
		return nil, err
	}
	proxy := NewProxy(targetUrl)
	chain := DefineServiceMiddlewares(resource)

	return &ServiceProxy{
		Proxy:      proxy,
		Middleware: chain,
		TargetUrl:  targetUrl,
		Endpoint:   resource.Service_Endpoint, // just to id the service
	}, nil
}

// RunServiceV2 is an upgraded version of RunService.
// In the previous version, applying different middleware to different services was restricted.
// RunServiceV2 removes that limitation by allowing each service to define its own middleware chain.
// Use DefineServiceMiddlewares to assign custom middleware per service.
func RunServiceV2(services_config *config.ServicesConfig, router *http.ServeMux) {
	for _, resource := range services_config.Services {
		// if enbaled in the config file
		if resource.Enabled {
			serviceProxy, err := NewService(resource)
			if err != nil {
				slog.Error("Invalid BaseURL, skipping service", slog.String("BaseUrl", resource.Baseurl))
				continue
			}
			// creates the router
			for _, routes := range resource.Routes {
				fmt.Println(routes.Method + resource.Service_Endpoint + routes.Endpoint)
				fullPath := routes.Method + resource.Service_Endpoint + routes.Endpoint
				handler := serviceProxy.Middleware.Then(ProxyRequestHandler(serviceProxy.Proxy, serviceProxy.TargetUrl, resource.Service_Endpoint))
				// adding more middleware create a Middleware type func and add here
				router.HandleFunc(fullPath, handler)
			}
		}
	}
}

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

				// adding more middleware create a Middleware type func and add here
				chain := middleware.CreateMiddlewareChain(middleware.RateLimiterMiddleware(resource.Rate_Limits))
				router.HandleFunc(routes.Method+resource.Service_Endpoint+routes.Endpoint, chain.Then(ProxyRequestHandler(proxy, url, resource.Service_Endpoint)))
			}
		}
	}
}

// DefineServiceMiddleware is a function where we can define different middleware Chain based on the service endpoint
// This fuction takes the resource and returns middlewareChain
// Can Edit this fuction as per your own needs.
func DefineServiceMiddlewares(resource config.Service) middleware.Chain {
	switch resource.Service_Endpoint {
	case "/service1":
		return middleware.CreateMiddlewareChain(middleware.RateLimiterMiddleware(resource.Rate_Limits))
	default:
		return middleware.CreateMiddlewareChain(middleware.RateLimiterMiddleware(resource.Rate_Limits))
	}
}

func NewProxy(target *url.URL) *httputil.ReverseProxy {
	return httputil.NewSingleHostReverseProxy(target)
}

func ProxyRequestHandler(proxy *httputil.ReverseProxy, url *url.URL, endpoint string) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// slog.Info("=================")
		// slog.Info("Request Received")
		// slog.Info("", slog.String("Url", r.URL.String()))
		// slog.Info("=================")
		r.URL.Host = url.Host
		r.URL.Scheme = url.Scheme
		r.Header.Set("X-Forwarded-Host", r.Header.Get("Host"))
		r.Host = url.Host
		path := r.URL.Path
		r.URL.Path = strings.TrimLeft(path, endpoint)
		// slog.Info("Redirecting Request")
		// slog.Info("", slog.String("Url", r.URL.String()))
		// slog.Info("=================")
		proxy.ServeHTTP(w, r)
	}

}
