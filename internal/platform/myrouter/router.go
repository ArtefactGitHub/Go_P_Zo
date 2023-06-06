package myrouter

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/ArtefactGitHub/Go_P_Zo/internal/platform/mycontext"
	"github.com/ArtefactGitHub/Go_P_Zo/internal/platform/mydb"
	"github.com/ArtefactGitHub/Go_P_Zo/internal/platform/myerror"
	"github.com/ArtefactGitHub/Go_P_Zo/internal/platform/myhttp"
	"github.com/ArtefactGitHub/Go_P_Zo/internal/test_v2"
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
		log.Printf("[handle]req: %v", req)
		log.Printf("[handle]params: %v", params)

		// StatusUnauthorized
		if needAuth {
			isSuccess, err := mycontext.FromContextBool(req.Context(), mycontext.AuthorizedKey)
			if err != nil || !isSuccess {
				myhttp.WriteError(w, myerror.NewError(errors.New("access token not found"), ""), http.StatusUnauthorized, "Unauthorized")
				return
			}
		}

		queryMap := common.QueryMap{}
		for _, p := range params {
			if _, ok := queryMap[p.Key]; !ok {
				queryMap[p.Key] = p.Value
			}
		}

		// TODO: リファクタ
		// データベースオブジェクトをcontextに詰めておく
		tx, err := mydb.Db.BeginTx(req.Context(), nil)
		if err != nil {
			myhttp.WriteError(w, myerror.NewError(errors.New(fmt.Sprintf("createHandle BeginTx error: %#v \n", err)), ""), http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
			return
		}
		ctx := test_v2.WithDBContext(req.Context(), mydb.Db)
		ctx = test_v2.WithTXContext(ctx, tx)
		req = req.WithContext(ctx)
		defer func() {
			if rec := recover(); rec != nil {
				log.Print("recover in defer")
				msg := fmt.Sprintf("cause panic: %#v", rec)
				if txErr := tx.Rollback(); txErr != nil {
					msg = fmt.Sprintf("%s and rollback error in recover: %#v \n", msg, txErr)
				}
				log.Print(msg)
				myhttp.WriteError(w, myerror.NewError(errors.New(msg), ""), http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
				return
			}
		}()

		f(w, req, queryMap)
		if txErr := tx.Commit(); txErr != nil {
			log.Printf("commit error: %#v \n", txErr)
			myhttp.WriteError(w, myerror.NewError(errors.New(fmt.Sprintf("commit error: %#v \n", err)), ""), http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
			return
		}
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
