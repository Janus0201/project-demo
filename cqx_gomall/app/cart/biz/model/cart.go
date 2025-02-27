package model

import (
	"context"
	"errors"

	cart "github.com/MrLittle05/Gomall/rpc_gen/kitex_gen/cart"
	"gorm.io/gorm"
)

type Cart struct {
	gorm.Model
	// index: 为字段添加索引，加快查询速度，尤其是在需要按 UserId 进行查找时
	UserId    uint32 `gorm:"type:int(11);not null;index:idx_user_id"`
	ProductId uint32 `gorm:"type:int(11);not null"`
	Quantity  uint32 `gorm:"type:int(11);not null"`
}

func (c Cart) TableName() string {
	return "cart"
}

type CartQuery struct {
	ctx context.Context
	// 由 gorm.Open() 返回的数据库连接对象
	db *gorm.DB
}

func NewCartQuery(ctx context.Context, db *gorm.DB) *CartQuery {
	return &CartQuery{
		ctx: ctx,
		db:  db,
	}
}

func (q CartQuery) AddItem(item *Cart) error {
	err := q.db.WithContext(q.ctx).
		Model(&Cart{}).
		Where(&Cart{UserId: item.UserId, ProductId: item.ProductId}).
		First(&Cart{}).Error
	if err == nil {
		return q.db.WithContext(q.ctx).
			Model(&Cart{}).
			Where(&Cart{UserId: item.UserId, ProductId: item.ProductId}).
			UpdateColumn("quantity", gorm.Expr("quantity + ?", item.Quantity)).Error
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	} else if item.UserId == 0 {
		return errors.New("user id is required")
	}
	return q.db.WithContext(q.ctx).Create(&item).Error
}

func (q CartQuery) EmptyCart(userId uint32) error {
	if userId == 0 {
		return errors.New("user id is required")
	}
	return q.db.WithContext(q.ctx).
		Where("user_id = ?", userId).
		Delete(&Cart{}).Error
}

func (q CartQuery) GetCart(userId uint32) (*cart.Cart, error) {
	if userId == 0 {
		return nil, errors.New("user id is required")
	}
	var adds []*Cart
	err := q.db.WithContext(q.ctx).
		Where("user_id = ?", userId).
		Find(&adds).Error
	items := make([]*cart.Item, len(adds))
	if err != nil {
		return nil, err
	}

	for i, c := range adds {
		items[i] = &cart.Item{
			ProductId: c.ProductId,
			Quantity:  c.Quantity,
		}
	}
	return &cart.Cart{
		UserId: userId,
		Items:  items,
	}, err
}
