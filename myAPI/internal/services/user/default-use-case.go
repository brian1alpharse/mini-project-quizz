package user

import (
	"context"
	"myAPI/internal/adapter/Repository"
	"myAPI/internal/domains"
	"myAPI/utils/helper"
)

type (
	UseCase struct {
		userRepo Repository.UserRepoInterface
	}

	UsecaseInterface interface {
		RegisterUser(
			ctx context.Context,
			payload RegisterUser,
		) (result UseCaseRegisterResult, err error)

		GetAll(
			ctx context.Context,
		) ([]domains.User, error)

		LoginUser(
			ctx context.Context,
			email, password string,
		) (*domains.User, string, error)

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

func (uc UseCase) RegisterUser(
	ctx context.Context,
	payload RegisterUser,
) (result UseCaseRegisterResult, err error) {

	hashPass := helper.HasPass(payload.Password)
	user, err := uc.userRepo.StoreUser(
		ctx,
		&domains.User{
			Nama:     payload.Nama,
			Email:    payload.Email,
			Password: hashPass,
			Role:     payload.Role,
		})
	result.User = RegisterUser{
		Nama:     user.Nama,
		Email:    user.Email,
		Password: hashPass,
		Role:     user.Role,
	}
	return result, err
}

func (uc UseCase) GetAll(
	ctx context.Context,
) ([]domains.User, error) {
	return uc.userRepo.GetAllUser(ctx)
}

func (uc UseCase) LoginUser(
	ctx context.Context,
	email, password string,
) (*domains.User, string, error) {
	user, err := uc.userRepo.LoginUser(ctx, email)
	if err != nil {
		return nil, "", err
	}

	// verify hashed password
	comparePass := helper.ComparePass([]byte(user.Password), []byte(password))
	if !comparePass {
		return nil, "", err
	}

	//Generate token JWT
	tokenString, errToken := helper.GenerateToken(user.ID, email)
	if errToken != nil {
		return nil, "", err
	}

	return user, tokenString, nil

}

func (uc UseCase) DeleteUser(
	ctx context.Context,
	id uint,
) (*domains.User, error) {
	user, err := uc.userRepo.DeleteUser(ctx, id)
	if err != nil {
		return nil, err
	}
	return user, nil

}

func (uc UseCase) UpdateUser(
	ctx context.Context,
	id uint,
	updateData map[string]interface{},
) (*domains.User, error) {
	return uc.userRepo.UpdateUser(ctx, id, updateData)
}
