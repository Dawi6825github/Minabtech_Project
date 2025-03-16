package models

import (
	"time"


	"gorm.io/gorm"
    "gorm.io/driver/postgres"
)

// Bookmark represents a user's bookmark for a recipe
type Bookmark struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID    uint      `gorm:"not null" json:"user_id"`
	RecipeID  uint      `gorm:"not null" json:"recipe_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	User   User   `gorm:"foreignKey:UserID"`
	Recipe Recipe `gorm:"foreignKey:RecipeID"`
}

// CreateBookmark creates a new bookmark
func CreateBookmark(db *gorm.DB, bookmark *Bookmark) error {
	return db.Create(bookmark).Error
}

// GetBookmarksByUser fetches bookmarks for a user
func GetBookmarksByUser(db *gorm.DB, userID uint) ([]Bookmark, error) {
	var bookmarks []Bookmark
	err := db.Where("user_id = ?", userID).Find(&bookmarks).Error
	return bookmarks, err
}
