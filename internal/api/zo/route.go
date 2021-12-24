package zo

import (
	"net/http"
)

func Routing(mux *http.ServeMux) {
	zc := zoController{}
	mux.HandleFunc("/zo", zc.handle)
	mux.HandleFunc("/zo/", zc.handle)
}
