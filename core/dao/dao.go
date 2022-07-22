package dao

import "gorm.io/gorm"

type Dao struct {
	*gorm.DB
}

func New(engine *gorm.DB) *Dao {
	return &Dao{engine}
}
