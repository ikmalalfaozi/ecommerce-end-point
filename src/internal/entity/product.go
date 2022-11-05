package entity

type Product struct {
	ProductID   int64  `gorm:"primaryKey;type:INTEGER;NOT NULL" json:"product_id"`
	Name        string `gorm:"type:VARCHAR(255);NOT NULL" json:"name"`
	Price       int64  `gorm:"type:BIGINT;NOT NULL" json:"price"`
	Quantity    int    `gorm:"type:INTEGER;NOT NULL" json:"quantity"`
	ImageUrl    string `gorm:"type:VARCHAR(255);NOT NULL" json:"image_url"`
	Description string `gorm:"type:TEXT;NOT NULL" json:"description"`
	Category    string `gorm:"type:VARCHAR(255);NOT NULL" json:"category"`
	Sku         string `gorm:"type:VARCHAR(255);NOT NULL" json:"sku"`
}