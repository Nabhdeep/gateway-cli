package service

import (
	"net/http"

	"github.com/nabhdeep/gateway-cli/pkg/config"
)

func Run_gateway_service() {
	gateway_config, _ := config.MustLoad()
	router := http.NewServeMux()
	router.HandleFunc("POST /", func(w http.ResponseWriter, r *http.Request) {})
	server := http.Server{
		Addr:    gateway_config.HttpServer.Address,
		Handler: router,
	}
	server.ListenAndServe()
}
