package user

import "myAPI/internal/adapter/dto"

type (
	RegisterUser struct {
		Nama     string `json:"nama"`
		Email    string `json:"email"`
		Password string `json:"password"`
		Role     string `json:"role"`
	}

	UseCaseRegisterResult struct {
		User RegisterUser `json:"user"`
	}

	SuccessLoginUser struct {
		Response    dto.ResponseMeta
		Email       string `json:"email"`
		AccessToken string `json:"accessToken"`
	}

	SuccessRemoveUser struct {
		Response dto.ResponseMeta
		ID       uint `json:"id"`
	}

	LoginParam struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	RemoveUser struct {
		ID uint `json:"id"`
	}

	UpdateUserPayload struct {
		ID         uint                   `json:"id"`
		UpdateData map[string]interface{} `json:"update_data"`
	}
)
