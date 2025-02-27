package utils

import (
	"context"
	"strconv"
	// "strconv"

	"github.com/cloudwego/hertz/pkg/common/hlog"
)

func GetUserIdFromCtx(ctx context.Context) int32 {
	userId := ctx.Value(SessionUserIdKey)
	if userId == nil {
		hlog.Warn("User ID not found in context")
		return 0
	}

	// 根据实际情况进行类型转换
	switch v := userId.(type) {
	case int32:
		return v
	case int:
		return int32(v)
	case string:
		if id, err := strconv.ParseInt(v, 10, 32); err == nil {
			return int32(id)
		}
		hlog.Errorf("Failed to parse user_id from string: %v", v)
		return 0
	default:
		hlog.Errorf("Invalid user_id type in context: %T", userId)
		return 0
	}
}
