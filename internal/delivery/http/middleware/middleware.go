package middleware

import (
	"github.com/fiqriardiansyah/shopping-api-golang/internal/helper"
	"gorm.io/gorm"
)

type Middleware struct {
	db     *gorm.DB
	config *helper.Config
}

func NewMiddleware(db *gorm.DB, config *helper.Config) *Middleware {
	return &Middleware{
		db,
		config,
	}
}
