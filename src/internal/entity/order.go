package entity

import "time"

type Order struct {
	OrderID    int64  		`gorm:"primaryKey;type:INTEGER; NOT NULL" json:"order_id"`
	CustomerID int64  		`gorm:"type:INTEGER;foreignKey:CustomerID;AssociationForeignKey:CustomerID; NOT NULL" json:"customer_id"`
	ProductID  int64 		`gorm:"type:INTEGER;foreignKey:ProductID;AssociationForeignKey:ProductID; NOT NULL" json:"product_id"`
	Quantity   int    		`gorm:"type:INTEGER; NOT NULL" json:"quantity"`
	Status     string 		`gorm:"type:VARCHAR(20); NOT NULL" json:"status"` // shipped, processed, sent, cancelled
	OrderDate  time.Time	`gorm:"type:DATETIME; NOT NULL" json:"order_date"`
	ShippedDate time.Time	`gorm:"type:DATETIME; NOT NULL" json:"shipped_date"`
}