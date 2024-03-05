package user

import (
	"myAPI/internal/adapter/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

type (
	RequestHandler struct {
		ctrl ControllerInterface
	}
)

func (rh RequestHandler) Register(ctx *gin.Context) {
	var payload = RegisterUser{}
	err := ctx.Bind(&payload)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.DefaultErrorInvalidDataWithMessage(err.Error()))
		return
	}
	res, err := rh.ctrl.Register(ctx, payload)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.DefaultErrorResponseWithMessage(err.Error()))
	}

	ctx.JSON(http.StatusOK, res)
}

func (rh RequestHandler) GetAll(ctx *gin.Context) {
	res, err := rh.ctrl.GetAll(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.DefaultErrorResponseWithMessage(err.Error()))
	}
	ctx.JSON(http.StatusOK, res)

}

func (rh RequestHandler) LoginUser(ctx *gin.Context) {
	var payload = LoginParam{}
	err := ctx.Bind(&payload)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.DefaultErrorInvalidDataWithMessage(err.Error()))
		return
	}
	res, err := rh.ctrl.LoginUser(ctx, payload.Email, payload.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.DefaultErrorResponseWithMessage(err.Error()))
	}

	ctx.JSON(http.StatusOK, res)
}
func (rh RequestHandler) Remove(ctx *gin.Context) {
	var payload RemoveUser
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, dto.DefaultErrorInvalidDataWithMessage(err.Error()))
		return
	}

	res, err := rh.ctrl.Remove(ctx, payload.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.DefaultErrorResponseWithMessage(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, res)
}
func (rh RequestHandler) UpdateUser(ctx *gin.Context) {
	var payload UpdateUserPayload
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, dto.DefaultErrorInvalidDataWithMessage(err.Error()))
		return
	}

	// Lakukan validasi payload di sini jika diperlukan

	// Panggil controller untuk melakukan update
	updatedUser, err := rh.ctrl.UpdateUser(ctx, payload.ID, payload.UpdateData)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.DefaultErrorResponseWithMessage(err.Error()))
		return
	}

	// Kirim respons dengan data customer yang telah diupdate
	ctx.JSON(http.StatusOK, updatedUser)
}
