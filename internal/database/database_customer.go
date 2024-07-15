package database

import (
	"context"

	"github.com/karlgama/go-microservices/internal/models"
)

func (c Client) GetAllCustomers(ctx context.Context, email string) ([]models.Customer, error) {
	var customers []models.Customer
	result := c.DB.WithContext(ctx).
		Where(models.Customer{Email: email}).
		Find(&customers)
	return customers, result.Error
}
