package gateway

import (
	"fmt"
	"log/slog"
	"net/http"

	gateway "github.com/nabhdeep/gateway-cli/gateway/proxy"
	"github.com/nabhdeep/gateway-cli/pkg/config"
)

func Inti_API_gateway() {
	gateway_config, services_config := config.MustLoad()
	fmt.Println(gateway_config)
	router := http.NewServeMux()
	router.HandleFunc("GET /ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("heo"))
	})
	server := http.Server{
		Addr:    gateway_config.HttpServer.Address,
		Handler: router,
	}
	slog.Info("API Gateway started on", slog.String("address", gateway_config.HttpServer.Address))

	// Attaching services
	gateway.RunService(services_config, router)

	err := server.ListenAndServe()
	if err != nil {
		slog.Error("Error staring the server")
	}

}
