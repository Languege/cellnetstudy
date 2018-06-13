package main

import (
	"os"
	"runtime/pprof" // 引用pprof package
)

func main() {
	f, _ := os.Create("profile_file")
	pprof.StartCPUProfile(f)     // 开始cpu profile，结果写到文件f中
	defer pprof.StopCPUProfile() // 结束profile

	var n int
	for i := 0; i < 100; i++ {
		n++
	}
}
