package handler

import (
	"context"

	"github.com/jakub-dzon/system-mcp/pkg/api"
	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
)

func HandleOverview(
	ctx context.Context,
	_ *mcp.CallToolRequest,
	_ api.SystemOverviewRequest,
) (*mcp.CallToolResult, api.SystemOverviewResponse, error) {
	var out api.SystemOverviewResponse

	info, err := host.InfoWithContext(ctx)
	if err != nil {
		return nil, out, err
	}

	cores, err := cpu.CountsWithContext(ctx, true)

	if err != nil {
		return nil, out, err
	}

	avg, err := load.AvgWithContext(ctx)
	if err != nil {
		return nil, out, err

	}

	memory, err := mem.VirtualMemoryWithContext(ctx)
	if err != nil {
		return nil, out, err

	}

	out = api.SystemOverviewResponse{
		UptimeSeconds: info.Uptime,
		CPUCores:      cores,
		Load1:         avg.Load1,
		Load5:         avg.Load5,
		Load15:        avg.Load15,
		MemoryTotal:   memory.Total,
		MemoryUsed:    memory.Used,
	}
	return nil, out, nil
}
