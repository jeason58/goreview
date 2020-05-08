package string

import (
	"runtime"
	"runtime/debug"
	"testing"
)

func TestSimplifyPath(t *testing.T) {

	//设置"M"最大数量（内核线程的最大数量）
	debug.SetMaxThreads(10000)

	//设置"P"的最大数量（G运行的上下文环境P）
	runtime.GOMAXPROCS(runtime.NumCPU())

	go func() {
		//将当前G与当前M（核心线程）绑定
		runtime.LockOSThread()
	}()

	runtime.GC()
}
