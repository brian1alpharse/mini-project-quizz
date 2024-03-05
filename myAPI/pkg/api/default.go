package api

import (
	"fmt"
	"log"
	"myAPI/internal/services/user"
	"myAPI/pkg/db"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Default() *Api {
	server := gin.Default()
	server.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*", "http://localhost:5173", "http://localhost:3000", "https://6yd68t.csb.app"},
		AllowHeaders:     []string{"X-Requested-with, Content-Type, Authorization, Access-Control-Allow-Origin"},
		AllowMethods:     []string{"POST, OPTIONS, GET, PUT, DELETE, OPTIONS"},
		ExposeHeaders:    []string{"Content.Length"},
		AllowCredentials: true,
	}))
	sqlConn, err := db.Default()
	if err != nil {
		log.Println(err)
		panic(fmt.Sprintf("panic at db connection: %s", err.Error()))
	}
	var routers = []Router{
		user.NewRoute(sqlConn),
	}
	return &Api{
		server:  server,
		routers: routers,
	}
}
