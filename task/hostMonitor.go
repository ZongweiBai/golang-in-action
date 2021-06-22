package task

import (
	"github.com/ZongweiBai/learning-go/config"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/net"
	"time"
)

// cpu info
func GetCpuInfo() {
	cpuInfos, err := cpu.Info()
	if err != nil {
		config.LOG.Errorf("get cpu info failed", err)
	}
	for _, ci := range cpuInfos {
		config.LOG.Info(ci)
	}
	// CPU使用率
	percent, _ := cpu.Percent(time.Second, false)
	config.LOG.Infof("cpu percent:%v", percent)
}

// 获取CPU负载信息
func GetCpuLoad() {
	info, _ := load.Avg()
	config.LOG.Infof("CPU load:%v", info)
}

// mem info
func GetMemInfo() {
	memInfo, _ := mem.VirtualMemory()
	config.LOG.Infof("mem info:%v", memInfo)
}

// host info
func GetHostInfo() {
	hInfo, _ := host.Info()
	config.LOG.Infof("host info:%v uptime:%v boottime:%v", hInfo, hInfo.Uptime, hInfo.BootTime)
}

// disk info
func GetDiskInfo() {
	parts, err := disk.Partitions(true)
	if err != nil {
		config.LOG.Errorf("get Partitions failed, err:%v", err)
		return
	}
	for _, part := range parts {
		config.LOG.Infof("part:%v\n", part.String())
		diskInfo, _ := disk.Usage(part.Mountpoint)
		config.LOG.Infof("disk info:used:%v free:%v", diskInfo.UsedPercent, diskInfo.Free)
	}

	ioStat, _ := disk.IOCounters()
	for k, v := range ioStat {
		config.LOG.Infof("iostat:%v:%v", k, v)
	}
}

func GetNetInfo() {
	info, _ := net.IOCounters(true)
	for index, v := range info {
		config.LOG.Infof("%v:%v send:%v recv:%v", index, v, v.BytesSent, v.BytesRecv)
	}
}
