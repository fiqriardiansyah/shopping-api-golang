package seeders

import (
	"fmt"
	"github.com/fiqriardiansyah/shopping-api-golang/internal/config"
	"github.com/fiqriardiansyah/shopping-api-golang/internal/entity"
	"gorm.io/gorm/clause"
)

func CategoryProductSeed() {
	categories := []entity.Category{
		{Name: "Electronics", Description: "Devices, gadgets, and accessories"},
		{Name: "Fashion", Description: "Clothing, shoes, and accessories"},
		{Name: "Home & Living", Description: "Furniture, appliances, and decor"},
		{Name: "Books", Description: "Fiction, non-fiction, and educational books"},
		{Name: "Sports & Outdoors", Description: "Sporting goods, outdoor gear, and fitness equipment"},
		{Name: "Beauty & Health", Description: "Cosmetics, skincare, and wellness products"},
		{Name: "Toys & Games", Description: "Kids toys, puzzles, and hobby items"},
	}

	db := config.NewDB()

	if err := db.Clauses(clause.OnConflict{DoNothing: true}).Create(&categories).Error; err != nil {
		panic(fmt.Sprintf("Failed seeding Role: %v", err))
	}

	fmt.Println("SEEDING CATEGORY PRODUCT SUCCESS ✅")
}
