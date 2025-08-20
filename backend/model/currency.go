package model

type Currency struct {
	ID       string `gorm:"primaryKey;type:varchar(3)" json:"id"`
	Name     string `gorm:"not null" json:"name"`
	Symbol   string `gorm:"not null" json:"symbol"`
	Country  string `gorm:"not null" json:"country"`
	Decimals int    `gorm:"not null" json:"decimals"`
}
