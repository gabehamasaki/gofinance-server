package models

type Account struct {
	BaseModel
	Name         string        `json:"name"`
	Email        string        `gorm:"unique" json:"email"`
	Password     string        `json:"password"`
	Transactions []Transaction `json:"transactions,omitempty" gorm:"foreignKey:Owner;reference:ID"`
}
