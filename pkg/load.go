package pkg

import (
	"fmt"
	"github.com/shirou/gopsutil/v4/cpu"
	"os"
	"time"
)

/*
Nagios插件需要返回特定的退出状态码，以便Nagios核心能够正确理解检查结果。这些状态码包括：
`0`：OK，表示检查结果正常。
`1`：WARNING，表示检查结果有警告。
`2`：CRITICAL，表示检查结果严重。
`3`：UNKNOWN，表示检查结果未知。
*/

func MonitorCpuLoad(n int) {
	percent, err := cpu.Percent(time.Second, false)
	if err != nil {
		fmt.Println("Error getting CPU percent:", err)
		os.Exit(3)
	}
	//  检查CPU使用率是否低于20%
	//  大于 20% 并且 小于 80%
	if percent[0] < float64(n) {
		fmt.Printf("CRITICAL - CPU usage is too low: %.2f%%\n", percent[0])
		os.Exit(2)
	} else if percent[0] > float64(n) && percent[0] < 80 {
		fmt.Printf("OK - CPU usage is normal: %.2f%%\n", percent[0])
		os.Exit(0)
	}

}
