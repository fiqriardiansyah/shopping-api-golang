package main

import (
	"fmt"

	"github.com/fiqriardiansyah/shopping-api-golang/db/seeders"
)

func main() {
	fmt.Println("START SEEDING ...")
	seeders.CategoryProductSeed()
	fmt.Println("SEEDING FINISH 🌱")
}
