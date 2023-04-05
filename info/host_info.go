package info

import (
	"fmt"

	"github.com/shirou/gopsutil/host"
)

type HostInfo struct {
	Hostname        string `json:"hostname"`
	Uptime          uint64 `json:"uptime"`
	BootTime        uint64 `json:"bootTime"`
	Procs           uint64 `json:"procs"`           // number of processes
	OS              string `json:"os"`              // ex: freebsd, linux
	Platform        string `json:"platform"`        // ex: ubuntu, linuxmint
	PlatformFamily  string `json:"platformFamily"`  // ex: debian, rhel
	PlatformVersion string `json:"platformVersion"` // version of the complete OS
	KernelVersion   string `json:"kernelVersion"`   // version of the OS kernel (if available)
	KernelArch      string `json:"kernelArch"`      // native cpu architecture queried at runtime, as returned by `uname -m` or empty string in case of error
	// VirtualizationSystem string `json:"virtualizationSystem"`
	// VirtualizationRole   string `json:"virtualizationRole"` // guest or host
	HostID string `json:"hostid"` // ex: uuid
}

func GetHostInfo() *HostInfo {
	hostInfo := &HostInfo{}
	info, _ := host.Info()
	fmt.Printf("info:%v", info)

	hostInfo.Hostname = info.Hostname
	hostInfo.OS = info.OS
	hostInfo.HostID = info.HostID
	hostInfo.Platform = info.Platform
	hostInfo.PlatformFamily = info.PlatformFamily
	hostInfo.PlatformVersion = info.PlatformVersion
	hostInfo.KernelArch = info.KernelArch
	hostInfo.KernelVersion = info.KernelVersion
	hostInfo.Uptime = info.Uptime
	hostInfo.BootTime = info.BootTime

	return hostInfo

}

func GetUpTime() (uint64, error) {
	return host.Uptime()
}
