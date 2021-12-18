package zo

import (
	"net/http"
)

func Routing() {
	zc := ZoController{}
	http.HandleFunc("/zo", zc.Handle)
	http.HandleFunc("/zo/", zc.Handle)
}
