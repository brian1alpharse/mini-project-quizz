package Repository

import (
	"context"
	"myAPI/internal/domians"

	"gorm.io/gorm"
)

type (
	CustomerRepo struct {
		db *gorm.DB
	}

	CustomerRepoInterface interface {
		StoreCustomer(
			ctx context.Context,
			cust *domians.Customer,
		) (*domians.Customer, error)

		GetAllEmployee(
			ctx context.Context,
		) ([]domians.Customer, error)

		LoginCustomer(
			ctx context.Context,
			email string,
		) (*domians.Customer, error)

		DeleteCustomer(
			ctx context.Context,
			id uint,
		) (*domians.Customer, error)

		UpdateCustomer(
			ctx context.Context,
			id uint,
			updateData map[string]interface{},
		) (*domians.Customer, error)
	}
)

func NewCustomerRepo(db *gorm.DB) CustomerRepoInterface {
	return CustomerRepo{db: db}
}

func (repo CustomerRepo) StoreCustomer(
	ctx context.Context,
	cust *domians.Customer,
) (*domians.Customer, error) {
	err := repo.db.WithContext(ctx).
		Create(&cust).
		Error
	return cust, err
}

func (repo CustomerRepo) GetAllEmployee(
	ctx context.Context,
) ([]domians.Customer, error) {
	var cust []domians.Customer
	err := repo.db.WithContext(ctx).Find(&cust).
		Error
	return cust, err
}

func (repo CustomerRepo) LoginCustomer(
	ctx context.Context,
	email string,
) (*domians.Customer, error) {
	customer := &domians.Customer{}

	err := repo.db.WithContext(ctx).
		Model(&domians.Customer{}).
		Where("email = ? ", email).
		First(customer).
		Error

	return customer, err
}

func (repo CustomerRepo) DeleteCustomer(
	ctx context.Context,
	id uint,
) (*domians.Customer, error) {
	customer := &domians.Customer{}

	err := repo.db.WithContext(ctx).
		Where("id = ?", id).
		Delete(customer).
		Error

	return customer, err
}

func (repo CustomerRepo) UpdateCustomer(
	ctx context.Context,
	id uint,
	updateData map[string]interface{},
) (*domians.Customer, error) {
	customer := &domians.Customer{}

	// Update data customer
	err := repo.db.WithContext(ctx).
		Model(&domians.Customer{}).
		Where("id = ?", id).
		Updates(updateData).
		Error

	if err != nil {
		return nil, err
	}

	// Ambil customer setelah diupdate
	err = repo.db.WithContext(ctx).
		First(customer, id).
		Error

	if err != nil {
		return nil, err
	}

	return customer, nil
}
