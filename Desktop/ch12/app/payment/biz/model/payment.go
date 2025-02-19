// Copyright 2024 CloudWeGo Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package model

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type PaymentLog struct {
	gorm.Model
	UserId        uint32    `json:"user_id"`
	OrderId       string    `json:"order_id"`
	TransactionId string    `json:"transaction_id"`
	Amount        float32   `json:"amount"`
	PayAt         time.Time `json:"pay_at"`
}

func (p PaymentLog) TableName() string {
	return "payment_log"
}

func CreatePaymentLog(db *gorm.DB, ctx context.Context, payment *PaymentLog) error {
	return db.WithContext(ctx).Model(&PaymentLog{}).Create(payment).Error
}

type CancelPaymentLog struct {
    gorm.Model
    UserId        uint32    `json:"user_id"`
    OrderId       string    `json:"order_id"`
    TransactionId string    `json:"transaction_id"`
    CancelAt      time.Time `json:"cancel_at"`
}

func (c CancelPaymentLog) TableName() string {
    return "cancel_payment_log"
}

func CreateCancelPaymentLog(db *gorm.DB, ctx context.Context, cancelPayment *CancelPaymentLog) error {
    return db.WithContext(ctx).Model(&CancelPaymentLog{}).Create(cancelPayment).Error
}

type ScheduleCancelPaymentLog struct {
    gorm.Model
    UserId        uint32    `json:"user_id"`
    OrderId       string    `json:"order_id"`
    TransactionId string    `json:"transaction_id"`
    ScheduleAt    time.Time `json:"schedule_at"`
    CancelAt      time.Time `json:"cancel_at"`
}

func (s ScheduleCancelPaymentLog) TableName() string {
    return "schedule_cancel_payment_log"
}

func CreateScheduleCancelPaymentLog(db *gorm.DB, ctx context.Context, scheduleCancelPayment *ScheduleCancelPaymentLog) error {
    return db.WithContext(ctx).Model(&ScheduleCancelPaymentLog{}).Create(scheduleCancelPayment).Error
}