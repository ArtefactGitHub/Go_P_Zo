package zo

import (
	"github.com/julienschmidt/httprouter"
)

func Routing(router *httprouter.Router) {
	zc := zoController{}
	router.GET("/api/v1/zos", zc.getAll)
	router.GET("/api/v1/zos/:zo_id", zc.get)
	router.POST("/api/v1/zos", zc.post)
	router.PUT("/api/v1/zos/:zo_id", zc.update)
	router.DELETE("/api/v1/zos/:zo_id", zc.delete)
}
