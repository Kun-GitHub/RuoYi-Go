package persistence

import (
	"RuoYi-Go/internal/domain/model"
	"gorm.io/gorm"
)

type DemoRepository struct {
	db *gorm.DB
}

func NewDemoRepository() *DemoRepository {
	return &DemoRepository{}
}

func (r *DemoRepository) GetDemoByID(id uint) (*model.Demo, error) {
	var demo model.Demo
	if err := r.db.First(&demo, id).Error; err != nil {
		return nil, err
	}
	return &demo, nil
}
