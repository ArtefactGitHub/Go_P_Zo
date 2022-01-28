module github.com/ArtefactGitHub/Go_P_Zo/main

go 1.17

replace github.com/ArtefactGitHub/Go_P_Zo/internal/config => ../internal/config

require (
	github.com/ArtefactGitHub/Go_P_Zo/internal/api/v1/auth v0.0.0-20220128073915-f9fbcaf60eb4
	github.com/ArtefactGitHub/Go_P_Zo/internal/api/v1/session v0.0.0-20220128073915-f9fbcaf60eb4
	github.com/ArtefactGitHub/Go_P_Zo/internal/api/v1/user v0.0.0-20220128073915-f9fbcaf60eb4
	github.com/ArtefactGitHub/Go_P_Zo/internal/api/v1/zo v0.0.0-20220128073915-f9fbcaf60eb4
	github.com/ArtefactGitHub/Go_P_Zo/internal/config v0.0.0-20220108142333-59a19c0068a6
	github.com/ArtefactGitHub/Go_P_Zo/internal/middleware v0.0.0-00010101000000-000000000000
	github.com/ArtefactGitHub/Go_P_Zo/internal/platform/mydb v0.0.0-20220108142333-59a19c0068a6
	github.com/go-sql-driver/mysql v1.6.0
)

require (
	github.com/ArtefactGitHub/Go_P_Zo/internal/platform/myauth v0.0.0-00010101000000-000000000000 // indirect
	github.com/ArtefactGitHub/Go_P_Zo/internal/platform/mycontext v0.0.0-00010101000000-000000000000 // indirect
	github.com/ArtefactGitHub/Go_P_Zo/internal/platform/myerror v0.0.0-20220109012703-3851fd0a803f // indirect
	github.com/ArtefactGitHub/Go_P_Zo/internal/platform/myhttp v0.0.0-20220109012703-3851fd0a803f // indirect
	github.com/ArtefactGitHub/Go_P_Zo/internal/platform/myrouter v0.0.0-20220109012703-3851fd0a803f // indirect
	github.com/ArtefactGitHub/Go_P_Zo/pkg/common v0.0.0-20220108142333-59a19c0068a6 // indirect
	github.com/golang-jwt/jwt v3.2.2+incompatible // indirect
	github.com/julienschmidt/httprouter v1.3.0 // indirect
	golang.org/x/crypto v0.0.0-20220112180741-5e0467b6c7ce // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)

replace github.com/ArtefactGitHub/Go_P_Zo/internal/platform/myerror => ../internal/platform/myerror

replace github.com/ArtefactGitHub/Go_P_Zo/internal/platform/mycontext => ../internal/platform/mycontext

replace github.com/ArtefactGitHub/Go_P_Zo/internal/platform/myhttp => ../internal/platform/myhttp

replace github.com/ArtefactGitHub/Go_P_Zo/internal/platform/myauth => ../internal/platform/myauth

replace github.com/ArtefactGitHub/Go_P_Zo/internal/platform/mydb => ../internal/platform/mydb

replace github.com/ArtefactGitHub/Go_P_Zo/internal/platform/myrouter => ../internal/platform/myrouter

replace github.com/ArtefactGitHub/Go_P_Zo/internal/middleware => ../internal/middleware

replace github.com/ArtefactGitHub/Go_P_Zo/internal/test => ../internal/test

replace github.com/ArtefactGitHub/Go_P_Zo/pkg/common => ../pkg/common
