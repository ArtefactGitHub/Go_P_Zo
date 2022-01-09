package myrouter

import (
	"net/http"

	"github.com/ArtefactGitHub/Go_P_Zo/internal/platform/mycontext"
	"github.com/ArtefactGitHub/Go_P_Zo/internal/platform/myerror"
	"github.com/ArtefactGitHub/Go_P_Zo/internal/platform/myhttp"
	"github.com/ArtefactGitHub/Go_P_Zo/pkg/common"
	"github.com/julienschmidt/httprouter"
)

type router struct {
	r      httprouter.Router
	routes map[RouteKey]func(w http.ResponseWriter, r *http.Request, ps common.QueryMap)
}

func (mr *router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	mr.r.ServeHTTP(w, req)
}

func NewMyRouter() *router {
	return &router{r: *httprouter.New()}
}

func NewMyRouterWithRoutes(routes ...map[RouteKey]func(
	w http.ResponseWriter,
	r *http.Request,
	ps common.QueryMap)) *router {

	result := &router{r: *httprouter.New()}

	// https://go.dev/ref/spec#Passing_arguments_to_..._parameters
	merged := merge(routes...)
	result.SetRoutes(merged)
	return result
}

func (mr *router) SetRoutes(routes map[RouteKey]func(
	w http.ResponseWriter,
	r *http.Request,
	ps common.QueryMap)) {

	addRoutes := getNewRoutes(mr.routes, routes)

	for key, handlerFunc := range addRoutes {
		needAuth := key.NeedAuth
		switch key.Method {
		case http.MethodGet:
			mr.r.GET(key.Path, createHandle(handlerFunc, needAuth))
		case http.MethodPost:
			mr.r.POST(key.Path, createHandle(handlerFunc, needAuth))
		case http.MethodPut, http.MethodPatch:
			mr.r.PUT(key.Path, createHandle(handlerFunc, needAuth))
		case http.MethodDelete:
			mr.r.DELETE(key.Path, createHandle(handlerFunc, needAuth))
		}
	}

	merge(mr.routes, addRoutes)
}

func createHandle(f func(
	w http.ResponseWriter,
	req *http.Request,
	params common.QueryMap),
	needAuth bool) httprouter.Handle {

	return func(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
		// StatusUnauthorized
		if needAuth {
			isSuccess, err := mycontext.FromContextBool(req.Context(), mycontext.AuthorizedKey)
			if err != nil || !isSuccess {
				myhttp.WriteError(w, myerror.NewError(nil, "token error"), http.StatusUnauthorized, "")
				return
			}
		}

		zoParams := common.QueryMap{}
		for _, p := range params {
			if _, ok := zoParams[p.Key]; !ok {
				zoParams[p.Key] = p.Value
			}
		}

		f(w, req, zoParams)
	}
}

func merge(maps ...map[RouteKey]func(w http.ResponseWriter, req *http.Request, params common.QueryMap)) map[RouteKey]func(w http.ResponseWriter, req *http.Request, params common.QueryMap) {
	result := make(map[RouteKey]func(w http.ResponseWriter, req *http.Request, params common.QueryMap), 0)

	for _, m := range maps {
		for k, v := range m {
			if _, ok := result[k]; !ok {
				result[k] = v
			}
		}
	}

	return result
}

func getNewRoutes(
	base map[RouteKey]func(w http.ResponseWriter, req *http.Request, params common.QueryMap),
	addRoutes map[RouteKey]func(w http.ResponseWriter, req *http.Request, params common.QueryMap)) map[RouteKey]func(w http.ResponseWriter, req *http.Request, params common.QueryMap) {

	result := make(map[RouteKey]func(w http.ResponseWriter, req *http.Request, params common.QueryMap), 0)

	for k, v := range addRoutes {
		if _, ok := base[k]; !ok {
			result[k] = v
		}
	}

	return result
}
