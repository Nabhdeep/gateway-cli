package main

import (
	"log/slog"
	"os"

	"github.com/nabhdeep/gateway-cli/cmd"
)

func main() {
	handler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	})
	logger := slog.New(handler)

	// Set the logger as default
	slog.SetDefault(logger)
	cmd.Execute()
}
