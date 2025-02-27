package utils

import "github.com/cloudwego/hertz/pkg/common/hlog"

func MustHandleError(err error) {
	if err != nil {
		// 在日志中以致命错误的形式记录 err ，并调用 os.Exit(1) 退出程序
		hlog.Fatal(err)
	}
}
