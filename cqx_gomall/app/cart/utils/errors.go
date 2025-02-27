package utils

import "github.com/cloudwego/kitex/pkg/klog"

func MustHandleError(err error) {
	if err != nil {
		// klog 是 kitex 提供的日志组件
		// 它会在日志中以致命错误的形式记录 err ，并调用 os.Exit(1) 退出程序
		klog.Fatal(err)
	}
}
