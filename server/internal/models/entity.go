package models

import "time"

type User struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	Username string `gorm:"unique;not null" json:"username"`
	Password string `json:"-"`
	Role     string `json:"role"`
}

type Supplier struct {
	ID      uint   `gorm:"primaryKey" json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Address string `json:"address"`
}

type Item struct {
	ID    uint   `gorm:"primaryKey" json:"id"`
	Name  string `json:"name"`
	Stock int    `json:"stock"`
	Price int64  `json:"price"`
}

type Purchasing struct {
	ID         uint      `gorm:"primaryKey" json:"id"`
	Date       time.Time `json:"date"`
	SupplierID uint      `json:"supplier_id"`
	UserID     uint      `json:"user_id"`
	GrandTotal int64     `json:"grand_total"`
	Supplier   Supplier           `gorm:"foreignKey:SupplierID" json:"supplier"`
	User       User               `gorm:"foreignKey:UserID" json:"user"`
	Details    []PurchasingDetail `gorm:"foreignKey:PurchasingID" json:"details"`
}

type PurchasingDetail struct {
	ID           uint  `gorm:"primaryKey" json:"id"`
	PurchasingID uint  `json:"purchasing_id"`
	ItemID       uint  `json:"item_id"`
	Qty          int   `json:"qty"`
	SubTotal     int64 `json:"sub_total"`
	Item         Item  `gorm:"foreignKey:ItemID" json:"item"`
}