package info

import (
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/load"
)

type CpuBaseInfo struct {
	ModelName       string `json:"modelName"`
	CPUCores        int    `json:"cpuCores"`
	CPULogicalCores int    `json:"couLogicalCores"`
}

type CpuCurrentInfo struct {
	CPUAvgPercent      float64   `json:"cpuAvgPercent"`
	CPUTotalUsePercent float64   `json:"CPUTotalUsePercent"`
	CouesUsedPercent   []float64 `json:"CouesUsedPercent"`
}

type LoadCurrentInfo struct {
	Load1        float64 `json:"load1"`
	Load5        float64 `json:"load5"`
	Load15       float64 `json:"load15"`
	UsagePercent float64 `json:"usagePercent"`
}

func GetCpuInfo() *CpuBaseInfo {
	//1 CPU全部信息
	cpuBaseInfo := &CpuBaseInfo{}
	cpuInfos, err := cpu.Info()
	if err == nil {
		cpuBaseInfo.ModelName = cpuInfos[0].ModelName
	}
	// for _, ci := range cpuInfos {
	// 	fmt.Println("CPU基本信息 : \n", ci)
	// }

	cpuBaseInfo.CPUCores, _ = cpu.Counts(false)
	cpuBaseInfo.CPULogicalCores, _ = cpu.Counts(true)
	return cpuBaseInfo
}

func GetCpuCurrentInfo(c *CpuBaseInfo) *CpuCurrentInfo {
	cpuCurrentInfo := &CpuCurrentInfo{}
	percent, _ := cpu.Percent(0, false)
	if len(percent) == 1 {
		cpuCurrentInfo.CPUAvgPercent = percent[0] * 0.01
		cpuCurrentInfo.CPUTotalUsePercent = cpuCurrentInfo.CPUAvgPercent * float64(c.CPUCores)
	}
	cpuCurrentInfo.CouesUsedPercent, _ = cpu.Percent(0, true)
	return cpuCurrentInfo
}

func GetLoadCurrentInfo(c *CpuBaseInfo) *LoadCurrentInfo {
	loadCurrentInfo := &LoadCurrentInfo{}
	loadInfo, _ := load.Avg()
	loadCurrentInfo.Load1 = loadInfo.Load1
	loadCurrentInfo.Load5 = loadInfo.Load5
	loadCurrentInfo.Load15 = loadInfo.Load15
	loadCurrentInfo.UsagePercent = loadCurrentInfo.Load1 / (float64(c.CPULogicalCores) * 0.85) * 100
	return loadCurrentInfo
}
