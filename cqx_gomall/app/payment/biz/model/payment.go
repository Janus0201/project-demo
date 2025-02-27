package model

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type PaymentLog struct {
	gorm.Model
	UserID        uint32    `json:"user_id"`
	OrderID       string    `json:"order_id"`
	TransactionId string    `json:"transaction_id"`
	Amount        float64   `json:"amount"`
	PayAt         time.Time `json:"pay_at"`
}

func (p *PaymentLog) TableName() string {
	return "payment_log"
}

/*
GORM 默认会根据结构体名称（小写复数形式）来推断数据库中的表名。例如：
如果结构体名为 PaymentLog，GORM 默认会将其映射到名为 payment_logs 的表。
如果结构体名为 User，GORM 默认会将其映射到名为 users 的表。
在这种情况下，你不需要显式指定表名，因为 GORM 会自动处理。

如果数据库中的表名与结构体的默认映射规则不一致，则需要显式指定表名。例如：
数据库中的表名为 payment_log（单数形式），而结构体名为 PaymentLog（默认映射为 payment_logs）。
在这种情况下，可以通过 TableName 方法或结构体标签指定表名。

.Model() 主要用于指定操作的模型（结构体），并自动推断该模型对应的表名。
如果模型定义了 TableName 方法或使用了自定义表名映射规则，则 .Model() 会根据这些规则来确定表名。

.Table() 直接指定要操作的表名，而不依赖任何模型或结构体。
它是一个更底层的操作，适用于需要直接操作数据库表的场景。
*/

func CreatePaymentLog(c context.Context, db *gorm.DB, p *PaymentLog) error {
	return db.WithContext(c).Model(&PaymentLog{}).Create(&p).Error
}
