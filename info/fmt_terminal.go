package info

import (
	"fmt"
	"time"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
)

type TableWriter struct {
	Writer table.Writer
}

func CreatTableWriter() *TableWriter {
	w := new(TableWriter)
	w.Writer = table.NewWriter()
	// newTable.SetAutoIndex(true)
	w.Writer.Style().Options.SeparateHeader = true
	w.Writer.Style().Options.SeparateRows = true
	w.Writer.Style().Options.SeparateFooter = true
	w.Writer.Style().Title.Align = text.AlignCenter
	// w.Writer.SetPageSize(20)
	return w
}
func (w *TableWriter) fmtPercent(name string) {

	//字体颜色
	warnColor := text.Colors{text.BgRed}
	warnTransFormer := text.Transformer(func(val interface{}) string {
		if val.(float64) > 80 {
			return warnColor.Sprintf("%.2f%%", val)
		}
		return fmt.Sprintf("%.2f%%", val)
	})
	w.Writer.Style().Format = table.FormatOptions{
		Header: text.FormatTitle,
	}
	w.Writer.SetColumnConfigs([]table.ColumnConfig{
		{
			Name:        name,
			Align:       text.AlignCenter,
			AlignHeader: text.AlignCenter,
			AlignFooter: text.AlignCenter,
			Transformer: warnTransFormer,
		},
	})

}

func (w *TableWriter) resetTable() {
	w.Writer.ResetRows()
	w.Writer.ResetHeaders()
}

func (w *TableWriter) updateHost() {
	hInfo := GetHostInfo()
	w.resetTable()
	w.Writer.SetTitle("系统信息")
	w.Writer.AppendRows([]table.Row{
		{"系统名称", hInfo.Hostname},
		{"系统类别", hInfo.OS},
		{"系统类型", hInfo.KernelArch},
		{"发行版本", fmt.Sprintf("%s-%s", hInfo.Platform, hInfo.PlatformVersion)},
		{"内核版本", hInfo.KernelVersion},
		{"启动时间", time.Unix(int64(hInfo.BootTime), 0).Format(time.DateTime)},
		{"运行时间", fmtSinceTime(time.Unix(int64(hInfo.Uptime), 0))},
	})
}
func (w *TableWriter) updateCPU() {
	cpuInfo := GetCpuInfo()
	cpuCurrentInfo := GetCpuCurrentInfo(cpuInfo)

	w.resetTable()
	w.Writer.SetTitle("cpu信息")
	w.Writer.AppendHeader(table.Row{fmt.Sprintf("Cores:%d, Logical Cores:%d", cpuInfo.CPUCores, cpuInfo.CPULogicalCores), cpuInfo.ModelName})
	w.Writer.AppendHeader(table.Row{"coreID", "Percent"})

	for idx, percent := range cpuCurrentInfo.CouesUsedPercent {
		w.Writer.AppendRow(table.Row{idx, percent})
	}
	w.Writer.AppendRow(table.Row{"Avg", cpuCurrentInfo.CPUAvgPercent})
	w.Writer.AppendRow(table.Row{"Total", cpuCurrentInfo.CPUTotalUsePercent})
	w.fmtPercent("Percent")

}

func (w *TableWriter) updateLoad() {
	cpuInfo := GetCpuInfo()
	loadInfo := GetLoadCurrentInfo(cpuInfo)
	w.resetTable()
	w.Writer.SetTitle("Load")
	w.Writer.AppendHeader(table.Row{"Load", "Percent"})

	w.Writer.AppendRow(table.Row{"Load1", loadInfo.Load1 * 100})
	w.Writer.AppendRow(table.Row{"Load5", loadInfo.Load5 * 100})
	w.Writer.AppendRow(table.Row{"Load15", loadInfo.Load15 * 100})
	w.Writer.AppendRow(table.Row{"Load1-UsagePercent", loadInfo.UsagePercent})
	w.fmtPercent("Percent")
}

func (w *TableWriter) updateMemory() {
	MemInfoLast := GetMemInfo()
	SwapMemInfoLast := GetSwapInfo()

	w.resetTable()
	w.Writer.SetTitle("内存信息")
	w.Writer.AppendHeader(table.Row{"type", "total", "used", "userPercent", "free", "available"})
	w.Writer.AppendRow(
		table.Row{
			MemInfoLast.Type,
			convertUnit(B, float64(MemInfoLast.Total)),
			convertUnit(B, float64(MemInfoLast.Used)),
			MemInfoLast.UsedPercent,
			convertUnit(B, float64(MemInfoLast.Free)),
			convertUnit(B, float64(MemInfoLast.Available)),
		})
	w.Writer.AppendRow(
		table.Row{
			SwapMemInfoLast.Type,
			convertUnit(B, float64(SwapMemInfoLast.Total)),
			convertUnit(B, float64(SwapMemInfoLast.Used)),
			SwapMemInfoLast.UsedPercent,
			convertUnit(B, float64(SwapMemInfoLast.Free)),
		})

	w.fmtPercent("userPercent")

}

func (w *TableWriter) updateDisk() {
	rows, err := GetDiskInfo()
	if err != nil {
		fmt.Printf("获取磁盘信息失败 , err:%v\n", err)
	}
	w.resetTable()
	// device
	w.Writer.SetTitle("硬盘信息-device")
	w.Writer.AppendHeader(table.Row{"Device", "fstype", "path", "total", "used", "free", "userPercent"})
	for _, row := range rows {
		w.Writer.AppendRow(
			table.Row{
				row.Device,
				row.Fstype,
				row.Path,
				convertUnit(B, float64(row.Total)),
				convertUnit(B, float64(row.Used)),
				convertUnit(B, float64(row.Free)),
				row.UsedPercent,
			})
	}

	w.fmtPercent("userPercent")
	w.render()
	// indoes
	w.resetTable()
	w.Writer.SetTitle("硬盘信息-Inodes")
	w.Writer.AppendHeader(table.Row{"Device", "total", "used", "free", "userPercent"})
	for _, row := range rows {
		w.Writer.AppendRow(
			table.Row{
				row.Device,
				convertUnit(B, float64(row.InodesTotal)),
				convertUnit(B, float64(row.InodesUsed)),
				convertUnit(B, float64(row.InodesFree)),
				row.InodesUsedPercent,
			})
	}
	w.fmtPercent("userPercent")

	fmt.Print("\n")
	fmt.Print(w.Writer.Render())
}

func (w *TableWriter) updateNet() {
	nets, err := GetNetIO()
	if err != nil {
		fmt.Printf("获取网卡信息失败 , err:%v\n", err)
	}
	w.Writer.SetTitle("网卡信息")
	w.resetTable()
	w.Writer.AppendHeader(table.Row{"Name", "BytesSent", "BytesRecv", "PacketsSent", "PacketsRecv"})
	for _, row := range nets {
		w.Writer.AppendRow(table.Row{row.Name, row.BytesSent, row.BytesRecv, row.PacketsSent, row.PacketsRecv})
	}

}
func (w *TableWriter) updateProc() {
	procs := GetProcInfo()

	w.Writer.SetTitle("进程信息")
	w.resetTable()

	w.Writer.AppendHeader(table.Row{"Name", "PID", "IsRunning", "CPU-Percent", "MEM-Percent", "threds", "UserName", "status"})
	for _, row := range procs {
		w.Writer.AppendRow(
			table.Row{
				row.Name,
				row.Pid,
				row.IsRunning,
				row.CPUPercent,
				float64(row.MemPercent),
				row.NumThreds,
				row.UserName,
				row.Status,
			})
	}

	w.fmtPercent("CPU-Percent")
	w.fmtPercent("MEM-Percent")
}

func (w *TableWriter) updateConnents() {
	conns := GetConnents()

	w.Writer.SetTitle("网络信息")
	w.resetTable()

	w.Writer.AppendHeader(table.Row{"Name", "PID", "type", "loacl IP", "local port", "remote IP", "remote port", "status"})
	for _, row := range conns {
		w.Writer.AppendRow(table.Row{GetProcForPid(row.Pid).Name, row.Pid, row.Type, row.LIP, row.LPort, row.RIP, row.RPort, row.Status})
	}

}

func (w *TableWriter) render() {
	// fmt.Fprintln(u.Write, t.Render())
	clear()
	fmt.Print(w.Writer.Render())
}

func (w *TableWriter) UITicker() {
	var a string
	ticker := time.NewTicker(time.Second * 1) // 创建一个定时器对象
	fmt.Print("start...\ncpu信息c,硬盘信息d,系统信息h,网卡信息i,负载信息l,网络信息n,内存信息m,内存信息p,退出q,")
	go func() {
		for {
			select {
			case <-ticker.C: // 每隔一秒，会执行一次
				switch a {
				case "c":
					w.updateCPU()
					w.render()
					ticker.Reset(time.Second * 1)
				case "d":
					w.updateDisk()
					ticker.Reset(time.Second * 2)
				case "h":
					w.updateHost()
					w.render()
					ticker.Reset(time.Second * 1)
				case "i":
					w.updateNet()
					w.render()
					ticker.Reset(time.Second * 3)
				case "l":
					w.updateLoad()
					w.render()
					ticker.Reset(time.Second * 2)
				case "n":
					w.updateConnents()
					w.render()
					ticker.Reset(time.Second * 5)
				case "m":
					w.updateMemory()
					w.render()
					ticker.Reset(time.Second * 1)
				case "p":
					w.updateProc()
					w.render()
					ticker.Reset(time.Second * 8)
				case "q":
					ticker.Stop()
				default:
					w.updateCPU()
					w.render()
				}
			}

		}

	}()

	for {
		fmt.Scanf("%s", &a)
		if a == "q" {
			fmt.Println("Bay Bay ")
			break
		}

	}
}
