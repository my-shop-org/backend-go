package main

import (
	"log"
	"math/rand"
	"product-service/config"
	"product-service/internal/entity"
	"strconv"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

func init() {
	if err := godotenv.Load(".env.local"); err != nil {
		log.Println("Warning: .env.local file not found, using environment variables")
	}
}

type DBSeeder struct {
	db *gorm.DB
}

func NewDBSeeder(db *gorm.DB) *DBSeeder {
	return &DBSeeder{db: db}
}

func (s *DBSeeder) SeedCategory() ([]*entity.Category, error) {
	categories := make([]*entity.Category, 11)

	for i := range categories {
		categories[i] = &entity.Category{
			Name:        "Category " + strconv.Itoa(i+1),
			Description: "Description for Category " + strconv.Itoa(i+1),
		}
	}

	s.db.Create(&categories)

	return categories, nil
}

func (s *DBSeeder) SeedProduct(categories []*entity.Category) error {

	tx := s.db.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	for i := 0; i < 50; i++ {
		p := entity.Product{
			Name:        "Product " + strconv.Itoa(i+1),
			Description: "Description for Product " + strconv.Itoa(i+1),
			BasePrice:   float64((i + 1) * 10),
		}

		if err := tx.Create(&p).Error; err != nil {
			tx.Rollback()
			return err
		}

		first := categories[i%len(categories)]
		second := categories[(i+1)%len(categories)]
		if err := tx.Model(&p).Association("Categories").Append(first, second); err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit().Error
}

func (s *DBSeeder) SeedAttributes() ([]*entity.Attribute, error) {
	attriubtes := make([]*entity.Attribute, 10)

	for i := range attriubtes {
		attriubtes[i] = &entity.Attribute{
			Name: "Attribute " + strconv.Itoa(i+1),
		}
	}

	s.db.Create(&attriubtes)

	return attriubtes, nil
}

func (s *DBSeeder) SeedAttributeValue(attributes []*entity.Attribute) error {
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 30; i++ {
		var attrID uint
		if len(attributes) > 0 {
			idx := rand.Intn(len(attributes))
			attrID = attributes[idx].ID
		} else {
			attrID = uint(rand.Intn(10) + 1)
		}

		av := &entity.AttributeValue{
			AttributeID: attrID,
			Value:       "Value " + strconv.Itoa(i+1),
		}
		if err := s.db.Create(&av).Error; err != nil {
			return err
		}
	}
	return nil
}

func (s *DBSeeder) Clear() error {
	if err := s.db.Exec("DELETE FROM product_categories").Error; err != nil {
		return err
	}
	if err := s.db.Exec("DELETE FROM products").Error; err != nil {
		return err
	}

	if err := s.db.Exec("DELETE FROM categories").Error; err != nil {
		return err
	}

	if err := s.db.Exec("DELETE FROM attributes").Error; err != nil {
		return err
	}
	if err := s.db.Exec("DELETE FROM attribute_values").Error; err != nil {
		return err
	}
	return nil
}

func (s *DBSeeder) Seed() error {
	categories, _ := s.SeedCategory()
	s.SeedProduct(categories)

	attributes, _ := s.SeedAttributes()
	s.SeedAttributeValue(attributes)
	return nil
}

func main() {
	db := config.ConnectDB()
	s := NewDBSeeder(db)

	s.Clear()
	s.Seed()
}
