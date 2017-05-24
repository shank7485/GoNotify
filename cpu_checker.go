package main


import (
	"fmt"
	"github.com/shirou/gopsutil/mem"
)

func bits_to_bytes(mem_bits uint64)  uint64{
	bytes := mem_bits/8
	return bytes
}

func bytes_to_mega_bytes(mem_bytes uint64)  uint64{
	mega_bytes := mem_bytes/1024
	return mega_bytes
}

func mega_bytes_to_giga_bytes(mem_mega_bytes uint64) uint64{
	giga_bytes := mem_mega_bytes/1024
	return giga_bytes
}

func find_CPU_usage() float64{
	a, _ := mem.VirtualMemory()
	return a.UsedPercent
}


func threshold_checker(threshold_percent int) bool{
	if find_CPU_usage() >= threshold_percent {
		return true
	} else {
		return false
	}
}


func main() {
	fmt.Print(find_CPU_usage())
}
