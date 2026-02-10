package main

import (
	"fmt"
	"os"

	"hot-coffee/internal/config"
	"hot-coffee/internal/server"
)

func main() {
	cfg, err := config.Parse(os.Args[1:])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	if cfg.Help {
		fmt.Print(config.UsageText())
		return
	}

	if err := server.RunServer(cfg); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}