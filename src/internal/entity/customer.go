package entity

type Customer struct {
	CustomerID int64  `gorm:"primaryKey;type:INTEGER" json:"customer_id"`
	Email      string `gorm:"type:VARCHAR(255);NOT NULL" json:"email"`
	Password   string `gorm:"type:VARCHAR(255);NOT NULL" json:"password"`
	FirstName  string `gorm:"type:VARCHAR(50);NOT NULL" json:"firstname"`
	LastName   string `gorm:"type:VARCHAR(50);NOT NULL" json:"lastname"`
	Address    string `gorm:"type:VARCHAR(255);NOT NULL" json:"address"`
	Zip        string `gorm:"type:VARCHAR(15);NOT NULL" json:"zip"`
	Phone      string `gorm:"type:VARCHAR(20);NOT NULL" json:"phone"`
}