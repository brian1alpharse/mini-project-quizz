package Repository

import (
	"context"
	"gorm.io/gorm"
	"myAPI/internal/domians"
)

type (
	EmployeeRepo struct {
		db *gorm.DB
	}

	EmployeeRepoInterface interface {
		StoreEmployee(
			ctx context.Context,
			emp *domians.Employee,
		) (*domians.Employee, error)
	}
)

func NewEmployeeRepo(db *gorm.DB) EmployeeRepoInterface {
	return EmployeeRepo{db: db}
}

func (repo EmployeeRepo) StoreEmployee(
	ctx context.Context,
	emp *domians.Employee,
) (*domians.Employee, error) {
	err := repo.db.WithContext(ctx).
		Create(&emp).
		Error
	return emp, err
}
