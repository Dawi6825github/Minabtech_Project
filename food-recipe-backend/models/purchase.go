package models

import (
	"time"

	"gorm.io/gorm"
)

// Purchase represents a purchase made by a user
type Purchase struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID    uint      `gorm:"not null" json:"user_id"`
	RecipeID  uint      `gorm:"not null" json:"recipe_id"`
	Amount    float32   `gorm:"not null" json:"amount"` // Amount in the respective currency
	Status    string    `gorm:"not null" json:"status"` // e.g., "pending", "completed"
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	User   User   `gorm:"foreignKey:UserID"`
	Recipe Recipe `gorm:"foreignKey:RecipeID"`
}

// CreatePurchase creates a new purchase
func CreatePurchase(db *gorm.DB, purchase *Purchase) error {
	return db.Create(purchase).Error
}

// GetPurchasesByUser fetches all purchases for a user
func GetPurchasesByUser(db *gorm.DB, userID uint) ([]Purchase, error) {
	var purchases []Purchase
	err := db.Where("user_id = ?", userID).Find(&purchases).Error
	return purchases, err
}
