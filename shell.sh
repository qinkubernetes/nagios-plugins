#!/bin/bash
# 获取CPU使用率

# top 信息
#PID: 每个进程的唯一标识符。
#USER: 启动该进程的用户。
#PR: 进程优先级。
#NI: nice值，影响进程优先级。
#VIRT: 虚拟内存使用量。
#RES: 常驻内存使用量（实际物理内存）。
#SHR: 与其他进程共享的内存量。
#S: 当前进程状态（如 R 表示运行，S 表示睡眠）。
#%CPU: 当前进程占用CPU的百分比。
#%MEM: 当前进程占用物理内存的百分比。
#TIME+: 累计使用CPU时间。
#COMMAND: 启动该进程的命令。

cpu_usage=$(top -bn1 | grep "Cpu(s)" | sed "s/.*, *\([0-9.]*\)%* id.*/\1/" | awk '{print 100 - $1}')

# 设置阈值
warning_threshold=20
cpu_threshold=80

# 检查CPU使用率
if (( $(echo "$cpu_usage < $warning_threshold" | bc -l) )); then
    echo "CRITICAL - CPU usage is below ${warning_threshold}%: ${cpu_usage}%"
    exit 2
elif (( $(echo "$cpu_usage > $warning_threshold" | bc -l) ) && ( $(echo "$cpu_usage > $cpu_threshold" | bc -l) ) ); then
    echo "CPU usage is below ${warning_threshold}%: ${cpu_usage}%"
    exit 0
fi
