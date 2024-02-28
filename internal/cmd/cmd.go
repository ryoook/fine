package cmd

import (
	"FINE/api"
	"context"
	// Autoload .dot file into ENV
	_ "github.com/joho/godotenv/autoload"
)

// Run a server
func Run() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Run daemon
	daemon(ctx)

	run(ctx, api.RouterHandler)
}
