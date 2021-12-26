package user

import (
	"github.com/julienschmidt/httprouter"
)

func Routing(router *httprouter.Router) {
	c := userController{}
	router.GET("/api/v1/users", c.getAll)
	router.GET("/api/v1/users/:user_id", c.get)
	router.POST("/api/v1/users", c.post)
	router.PUT("/api/v1/users/:user_id", c.update)
	router.DELETE("/api/v1/users/:user_id", c.delete)
}
