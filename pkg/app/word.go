package app

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type Words struct {
	gorm.Model
	Word string `json:"word"`
}

func ListWord(ctx context.Context, db *gorm.DB) ([]Words, error) {
	var words []Words
	result := db.WithContext(ctx).Find(&words)
	if result.Error != nil {
		return nil, fmt.Errorf("app.ListWord : %w", result.Error)
	}
	return words, nil
}

func (w *Words) Create(ctx context.Context, db *gorm.DB) error {
	return db.WithContext(ctx).Create(w).Error
}
