package main

import (
	"context"
	"log"

	"github.com/jakub-dzon/system-mcp/internal/handler"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

func main() {
	server := mcp.NewServer(
		&mcp.Implementation{
			Name:    "system-diagnostics",
			Version: "0.1.0",
		}, nil)

	mcp.AddTool(
		server,
		&mcp.Tool{
			Name:        "get_system_overview",
			Description: "Read Linux host health metrics",
		},
		handler.HandleOverview,
	)

	if err := server.Run(
		context.Background(),
		&mcp.StdioTransport{},
	); err != nil {
		log.Fatal(err)
	}

}
