package info

import (
	"fmt"
	"time"

	"github.com/shirou/gopsutil/mem"
)

type MemInfo struct {
	Type        string  `json:"type"`
	Total       uint64  `json:"total"`
	Used        uint64  `json:"used"`
	Free        uint64  `json:"free"`
	Available   uint64  `json:"available"`
	UsedPercent float64 `json:"usedPercent"`
}

// 内存信息
func GetMemInfo() *MemInfo {
	memInfo := &MemInfo{}
	//获取内存信息
	Stat, _ := mem.VirtualMemory()
	memInfo.Type = "Mem"
	memInfo.Total = Stat.Total
	memInfo.Used = Stat.Used
	memInfo.Free = Stat.Free
	memInfo.UsedPercent = Stat.UsedPercent
	memInfo.Available = Stat.Available
	return memInfo

}

func GetSwapInfo() *MemInfo {
	memInfo := &MemInfo{}
	//获取交换分区信息
	Stat, _ := mem.SwapMemory()
	memInfo.Type = "Swap"
	memInfo.Total = Stat.Total
	memInfo.Used = Stat.Used
	memInfo.Free = Stat.Free
	memInfo.UsedPercent = Stat.UsedPercent
	return memInfo

}

func ContinueMem() {
	var a string
	fmt.Printf("disk io\n ")
	ticker := time.NewTicker(time.Second * 1) // 创建一个定时器对象
	//开启go协程
	go func() {
		for {
			select {
			case <-ticker.C: // 每隔一秒，会执行一次
				fmt.Print(".....")
			}
			if a == "q" {
				ticker.Stop()
			}
		}

	}()

	for {
		fmt.Scanf("%s", &a)
		if a == "q" {
			time.Sleep(time.Second)
			fmt.Println("Bay Bay ")
			break
		}
	}
}
