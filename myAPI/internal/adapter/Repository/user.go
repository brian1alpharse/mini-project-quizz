package Repository

import (
	"context"
	"myAPI/internal/domains"

	"gorm.io/gorm"
)

type (
	UserRepo struct {
		db *gorm.DB
	}

	UserRepoInterface interface {
		StoreUser(
			ctx context.Context,
			cust *domains.User,
		) (*domains.User, error)

		GetAllUser(
			ctx context.Context,
		) ([]domains.User, error)

		LoginUser(
			ctx context.Context,
			email string,
		) (*domains.User, error)

		DeleteUser(
			ctx context.Context,
			id uint,
		) (*domains.User, error)

		UpdateUser(
			ctx context.Context,
			id uint,
			updateData map[string]interface{},
		) (*domains.User, error)
	}
)

func NewUserRepo(db *gorm.DB) UserRepoInterface {
	return UserRepo{db: db}
}

func (repo UserRepo) StoreUser(
	ctx context.Context,
	cust *domains.User,
) (*domains.User, error) {
	err := repo.db.WithContext(ctx).
		Create(&cust).
		Error
	return cust, err
}

func (repo UserRepo) GetAllUser(
	ctx context.Context,
) ([]domains.User, error) {
	var cust []domains.User
	err := repo.db.WithContext(ctx).Find(&cust).
		Error
	return cust, err
}

func (repo UserRepo) LoginUser(
	ctx context.Context,
	email string,
) (*domains.User, error) {
	User := &domains.User{}

	err := repo.db.WithContext(ctx).
		Model(&domains.User{}).
		Where("email = ? ", email).
		First(User).
		Error

	return User, err
}

func (repo UserRepo) DeleteUser(
	ctx context.Context,
	id uint,
) (*domains.User, error) {
	User := &domains.User{}

	err := repo.db.WithContext(ctx).
		Where("id = ?", id).
		Delete(User).
		Error

	return User, err
}

func (repo UserRepo) UpdateUser(
	ctx context.Context,
	id uint,
	updateData map[string]interface{},
) (*domains.User, error) {
	User := &domains.User{}

	// Update data User
	err := repo.db.WithContext(ctx).
		Model(&domains.User{}).
		Where("id = ?", id).
		Updates(updateData).
		Error

	if err != nil {
		return nil, err
	}

	// Ambil User setelah diupdate
	err = repo.db.WithContext(ctx).
		First(User, id).
		Error

	if err != nil {
		return nil, err
	}

	return User, nil
}
