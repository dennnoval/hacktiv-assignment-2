package entity

import (
	t "time"
	// e "hacktiv-assignment-2/entity"
)

type Items struct {
	ItemID uint64 `gorm:"primaryKey;column:item_id;autoIncrement" json:"-"`
	ItemCode uint64 `gorm:"column:item_code" json:"itemCode"`
	Description string `json:"description"`
	Quantity int `json:"quantity"`
	OrderID uint64 `gorm:"column:order_id" json:"-"`
}

type Orders struct {
	OrderID uint64 `gorm:"primaryKey;column:order_id;autoIncrement" json:"-"`
	CustomerName string `gorm:"column:customer_name" json:"customerName"`
	OrderedAt t.Time `gorm:"column:ordered_at;autoCreateTime;autoUpdateTime" json:"orderedAt"`
	Items *[]Items `gorm:"foreignKey:OrderID" json:"items"`
}
