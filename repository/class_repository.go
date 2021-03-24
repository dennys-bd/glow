package repository

import "github.com/dennys-bd/glow/entity"

type ClassRepository struct{}

func ProvideClassRepo() ClassRepository {
	return ClassRepository{}
}

func (r ClassRepository) Create(c *entity.Class) error {
	return db.Create(c).Error
}

func (r ClassRepository) Find(id uint) (*entity.Class, error) {
	var class entity.Class
	err := db.First(&class, id).Error
	return &class, err
}

func (r ClassRepository) List(params map[string]interface{}) ([]*entity.Class, error) {
	var classes []*entity.Class
	err := db.Find(&classes, params).Error
	return classes, err
}
