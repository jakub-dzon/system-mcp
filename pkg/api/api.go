package api

type SystemOverviewRequest struct{}

type SystemOverviewResponse struct {
	UptimeSeconds uint64  `json:"uptimeSeconds"`
	CPUCores      int     `json:"cpuCores"`
	Load1         float64 `json:"load1"`
	Load5         float64 `json:"load5"`
	Load15        float64 `json:"load15"`
	MemoryTotal   uint64  `json:"memoryTotalBytes"`
	MemoryUsed    uint64  `json:"memoryUsedBytes"`
}
