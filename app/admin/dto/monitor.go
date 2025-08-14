package dto

import (
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
)

type ServerInfoVo struct {
	Host       *host.InfoStat         `json:"host"`
	Memory     *mem.VirtualMemoryStat `json:"memory"`
	CPU        []cpu.InfoStat         `json:"cpu"`
	CPUPercent float64                `json:"cpuPercent"`
	Disk       []*disk.UsageStat      `json:"disk"`
	RemoteIp   string                 `json:"remoteIp"`
}
