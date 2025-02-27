package utils

import (
	"github.com/cloudwego/kitex/pkg/klog"
	"gorm.io/gorm"
)

func checkConnectionPoolStatus(db *gorm.DB) {
	// 获取底层 sql.DB 实例
	sqlDB, err := db.DB()
	if err != nil {
		klog.Fatalf("Failed to get sql.DB instance: %v", err)
	}

	// 获取连接池状态
	stats := sqlDB.Stats()

	// 打印连接池状态
	klog.Info("Connection Pool Stats:")
	klog.Infof("- OpenConnections: %d", stats.OpenConnections)       // 当前打开的连接数
	klog.Infof("- InUse: %d", stats.InUse)                           // 当前正在使用的连接数
	klog.Infof("- Idle: %d", stats.Idle)                             // 当前空闲的连接数
	klog.Infof("- MaxOpenConnections: %d", stats.MaxOpenConnections) // 最大打开连接数
}
