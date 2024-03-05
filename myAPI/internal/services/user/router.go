package user

import (
	"myAPI/internal/adapter/Repository"
	"myAPI/pkg/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type (
	Router struct {
		rq *RequestHandler
	}
)

func NewRoute(
	db *gorm.DB,
) *Router {
	return &Router{rq: &RequestHandler{
		ctrl: &Controller{
			Uc: UseCase{
				userRepo: Repository.NewUserRepo(db),
			},
		},
	},
	}

}

func (r Router) Route(router *gin.RouterGroup) {

	user := router.Group("/user")
	user.POST(
		"/register",
		r.rq.Register,
	)
	user.GET(
		"/all",
		// middleware.Authentication(),
		r.rq.GetAll,
	)

	user.POST(
		"/login",
		//middleware.Cors(),
		r.rq.LoginUser,
	)

	user.DELETE(
		"/remove",
		r.rq.Remove,
	)

	user.PUT(
		"/update",
		middleware.Authentication(),
		r.rq.UpdateUser)

}
