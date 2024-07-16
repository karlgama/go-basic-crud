package database

import (
	"context"

	"github.com/karlgama/go-microservices/internal/models"
)

func (c Client) GetAllProducts(ctx context.Context) ([]models.Product, error) {
	var products []models.Product
	result := c.DB.WithContext(ctx).Find(&products)
	return products, result.Error
}
