package info

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"time"
)

const (
	B = iota
	KB
	MB
	GB
	TB
	PB
)

func unitLoop(i int, val float64) (int, float64) {
	if val > 1024 {
		v := val / 1024
		return unitLoop(i+1, v)
	}
	return i, val
}

func convertUnit(u int, val float64) string {
	if val > 1024 {
		u, v := unitLoop(u, val)
		switch u {
		case B:
			return fmt.Sprintf("%.2fB", v)
		case KB:
			return fmt.Sprintf("%.2fKB", v)
		case MB:
			return fmt.Sprintf("%.2fMB", v)
		case GB:
			return fmt.Sprintf("%.2fGB", v)
		case TB:
			return fmt.Sprintf("%.2fTB", v)
		case PB:
			return fmt.Sprintf("%.2fPB", v)
		default:
			return fmt.Sprintf("%.2f", v)
		}
	}
	return fmt.Sprintf("%.2f", val)

}

var ros = runtime.GOOS

func clear() {

	if ros == "linux" {
		cmd := exec.Command("clear") //Linux example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else if ros == "windows" {
		cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}

}

var ZERO_TIME = time.Unix(0, 0)

func fmtSinceTime(t time.Time) string {
	d := t.Day() - ZERO_TIME.Day()
	h := t.Hour() - ZERO_TIME.Hour()
	// 如果计算出数小时为负数，向天数借一，给到小时数加24小时
	if h < 0 {
		h += 24
		d -= 1
	}

	return fmt.Sprintf("%02d天  %02d 时 %02d 分 %02d 秒", d, h, t.Minute(), t.Second())
}
