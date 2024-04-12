package database

import "time"

type Customer struct {
	CID       string    `gorm:"primaryKey;not null" json:"cid"`
	Name      string    `gorm:"not null" json:"name"`
	Email     string    `gorm:"unique;not null" json:"email"`
	Address   string    `gorm:"not null" json:"address"`
	City      string    `gorm:"not null" json:"city"`
	Zip       string    `gorm:"not null" json:"zip"`
	Orders    []Order   `gorm:"foreignKey:CID;constraint:OnDelete:CASCADE" json:"orders"` // ความสัมพันธ์ One-to-Many กับตาราง Order
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type Product struct {
	PID         string      `gorm:"primaryKey;not null" json:"pid"`
	Name        string      `gorm:"not null" json:"name"`
	Description string      `gorm:"not null" json:"description"`
	UnitPrice   float64     `gorm:"not null" json:"unit_price"`
	CreatedAt   time.Time   `json:"createdAt"`
	UpdatedAt   time.Time   `json:"updatedAt"`
	OrderLines  []OrderLine `gorm:"foreignKey:PID;constraint:OnDelete:SET NULL"` // ความสัมพันธ์ One-to-Many กับตาราง OrderLine
}

type Admin struct {
	AID       string    `gorm:"primaryKey;not null" json:"aid"`
	Email     string    `gorm:"unique;not null" json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type Order struct {
	OID        string      `gorm:"primaryKey;not null" json:"oid"`
	CID        string      `gorm:"not null" json:"cid"`                                           // Foreign key ไปยัง Customer
	OrderLines []OrderLine `gorm:"foreignKey:OID;constraint:OnDelete:CASCADE" json:"order_lines"` // ความสัมพันธ์ One-to-Many กับตาราง OrderLine
	CreatedAt  time.Time   `json:"createdAt"`
	UpdatedAt  time.Time   `json:"updatedAt"`
	// Customer   Customer    `gorm:"foreignKey:CID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type OrderLine struct {
	OID       string    `gorm:"primaryKey;not null" json:"oid"`
	PID       string    `gorm:"primaryKey;not null" json:"pid"`
	Quantity  int       `gorm:"not null" json:"quantity"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	// Order     Order     `gorm:"foreignKey:OID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"` // กำหนด foreign key constraint
	// Product   Product   `gorm:"foreignKey:PID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"` // กำหนด foreign key constraint

}
