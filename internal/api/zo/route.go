package zo

import (
	"net/http"
)

func Routing() {
	zc := zoController{}
	http.HandleFunc("/zo", zc.handle)
	http.HandleFunc("/zo/", zc.handle)
}
