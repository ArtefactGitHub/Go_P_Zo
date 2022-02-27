module github.com/ArtefactGitHub/Go_P_Zo

go 1.17

replace internal/api/v1/client => ./internal/api/v1/client

replace internal/api/v1/session => ./internal/api/v1/session

replace internal/api/v1/user => ./internal/api/v1/user

replace internal/api/v1/zo => ./internal/api/v1/zo

replace internal/config => ./internal/config

replace internal/middleware => ./internal/middleware

replace internal/platform/myerror => ./internal/platform/myerror

replace internal/platform/mycontext => ./internal/platform/mycontext

replace internal/platform/myhttp => ./internal/platform/myhttp

replace internal/platform/myauth => ./internal/platform/myauth

replace internal/platform/mydb => ./internal/platform/mydb

replace internal/platform/myrouter => ./internal/platform/myrouter

replace internal/test => ./internal/test

replace pkg/common => ./pkg/common

require (
	github.com/go-sql-driver/mysql v1.6.0
	github.com/golang-jwt/jwt v3.2.2+incompatible
	github.com/julienschmidt/httprouter v1.3.0
	golang.org/x/crypto v0.0.0-20220214200702-86341886e292
	gopkg.in/yaml.v2 v2.4.0
)
