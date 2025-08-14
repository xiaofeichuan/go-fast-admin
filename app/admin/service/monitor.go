package service

import (
	"fmt"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
	"go-fast-admin/server/app/admin/dto"
	"go-fast-admin/server/common/tools"
	"io/ioutil"
	"net/http"
	"time"
)

type MonitorService struct{}

// GetCacheKeys 获取缓存key
func (monitorService *MonitorService) GetCacheKeys() []string {
	var keys = tools.Redis.GetAllKeys()
	return keys
}

// GetCacheValue 获取缓存value
func (monitorService *MonitorService) GetCacheValue(key string) (value interface{}) {
	tools.Redis.Get(key, &value)
	return value
}

// DeleteCache 删除缓存
func (monitorService *MonitorService) DeleteCache(key string) {
	tools.Redis.Del(key)
}

// GetServerInfo 获取服务器信息
func (monitorService *MonitorService) GetServerInfo() (server dto.ServerInfoVo) {

	// 获取主机信息
	hostInfo, err := host.Info()
	if err != nil {
		fmt.Printf("Failed to get host info: %v", err)
		return
	}
	server.Host = hostInfo
	interval := 1 * time.Second // 定义采样间隔
	// 获取CPU信息
	cpuInfo, err := cpu.Info()
	if err != nil {
		fmt.Printf("Failed to get CPU info: %v", err)
		return
	}
	cpuPercent, err := cpu.Percent(interval, false)
	server.CPU = cpuInfo
	server.CPUPercent = cpuPercent[0]

	// 获取内存信息
	memInfo, err := mem.VirtualMemory()
	if err != nil {
		fmt.Printf("Failed to get memory info: %v", err)
		return
	}
	server.Memory = memInfo

	// 获取磁盘分区信息
	partitions, err := disk.Partitions(true)
	if err != nil {
		fmt.Printf("Failed to get disk partitions: %v", err)
		return
	}
	fmt.Printf("Disk partitions: %+v\n", partitions)
	for _, partition := range partitions {
		// 获取每个磁盘分区的使用情况
		usage, err := disk.Usage(partition.Mountpoint)
		if err != nil {
			fmt.Printf("Failed to get disk usage for %s: %v", partition.Mountpoint, err)
			continue
		}
		server.Disk = append(server.Disk, usage)
	}

	// 获取公网ip
	resp, err := http.Get("http://myip.ipip.net/")
	if err != nil {
		fmt.Println(err)
		return
	}

	// 读取请求响应的数据
	ip, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 输出当前公网 IP
	server.RemoteIp = string(ip)

	return server
}
