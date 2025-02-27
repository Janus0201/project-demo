package model

import "gorm.io/gorm"

type User struct {
	// 这是一个嵌入结构体，它包含了以下默认字段：ID, CreatedAt, UpdatedAt, DeletedAt。这些字段提供了基本的时间戳管理和软删除功能。
	gorm.Model

	// 用 GORM 标签来指定字段的数据库映射规则
	// GORM 可以自动推断字段类型和列名
	// 使用 gorm:"uniqueIndex" 标签表示这个字段应该有唯一索引，确保数据库中的每个用户都有一个唯一的电子邮件地址
	// 使用 gorm:"type:varchar(255) not null" 标签指定数据库中的字段类型为 varchar(255) 并且不能为空
	Email string `gorm:"uniqueIndex; type:varchar(255) not null"`

	PasswordHashed string `gorm:"type:varchar(255) not null"`
}

func (user *User) TableName() string {
	return "user"
}

func Create(db *gorm.DB, user *User) error {
	return db.Create(user).Error
}

func GetByEmail(db *gorm.DB, email string) (*User, error) {
	var user User
	err := db.Where("email = ?", email).First(&user).Error
	return &user, err
}
