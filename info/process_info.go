package info

import "github.com/shirou/gopsutil/process"

type ProcessStat struct {
	Pid        int32
	Ppid       int32
	Name       string
	UserName   string
	Status     string
	NumThreds  int32
	NumFDs     int32
	CPUPercent float64
	MemPercent float32
	IsRunning  bool
	Cmdline    string
	Exe        string
}

func GetProcInfo() []ProcessStat {
	procStatALL := []ProcessStat{}
	procs, _ := process.Processes()
	for _, proc := range procs {
		procStat := ProcessStat{}
		procStat.InitProcStat(proc)

		procStatALL = append(procStatALL, procStat)
		// fmt.Printf("%v-%v,%s,%v|%d, %d, %.2f,%.2f|%v||%v,%v\n", proc.Pid, ppid, name, status, numFDs, numThreds, cpuPercent, memPercent, isrunning, cmdline, userName)
	}
	return procStatALL
}

func (ps *ProcessStat) InitProcStat(proc *process.Process) {
	ps.Pid = proc.Pid
	ps.Ppid, _ = proc.Ppid()
	ps.Name, _ = proc.Name()
	ps.UserName, _ = proc.Username()
	ps.Status, _ = proc.Status()
	ps.NumThreds, _ = proc.NumThreads()
	ps.NumFDs, _ = proc.NumFDs()
	ps.CPUPercent, _ = proc.CPUPercent()
	ps.MemPercent, _ = proc.MemoryPercent()
	ps.IsRunning, _ = proc.IsRunning()
	ps.Cmdline, _ = proc.Cmdline()
	ps.Exe, _ = proc.Exe()
}

func GetProcForPid(pid int32) ProcessStat {
	proc, err := process.NewProcess(pid)
	procStat := ProcessStat{}
	if err != nil {
		return procStat

	}

	procStat.InitProcStat(proc)
	return procStat
}
