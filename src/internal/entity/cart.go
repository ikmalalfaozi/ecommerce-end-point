package entity

type Cart struct {
	CustomerID int64 `gorm:"type:INTEGER;foreignKey:CustomerID;AssociationForeignKey:CustomerID; NOT NULL" json:"customer_id"`
	ProductID  int64 `gorm:"type:INTEGER;foreignKey:ProductID;AssociationForeignKey:ProductID; NOT NULL" json:"product_id"`
	Quantity   int   `gorm:"type:INTEGER; NOT NULL" json:"quantity"`
}