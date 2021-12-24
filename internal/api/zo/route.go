package zo

import (
	"github.com/julienschmidt/httprouter"
)

func Routing(router *httprouter.Router) {
	zc := zoController{}
	router.GET("/zo", zc.getAll)
	router.GET("/zo/:id", zc.get)
	router.POST("/zo", zc.post)
	router.PUT("/zo/:id", zc.update)
	router.DELETE("/zo/:id", zc.delete)
}
