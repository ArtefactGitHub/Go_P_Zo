module github.com/ArtefactGitHub/Go_P_Zo/main

go 1.17

replace github.com/ArtefactGitHub/Go_P_Zo/internal/config => ../internal/config

require (
	github.com/ArtefactGitHub/Go_P_Zo/internal/api/v1/user v0.0.0-00010101000000-000000000000
	github.com/ArtefactGitHub/Go_P_Zo/internal/api/v1/zo v0.0.0-20211226064812-51c406b24abe
	github.com/ArtefactGitHub/Go_P_Zo/internal/config v0.0.0-20211226064812-51c406b24abe
	github.com/ArtefactGitHub/Go_P_Zo/internal/platform/mydb v0.0.0-20211226064812-51c406b24abe
	github.com/ArtefactGitHub/Go_P_Zo/internal/platform/myrouter v0.0.0-00010101000000-000000000000
	github.com/go-sql-driver/mysql v1.6.0
)

require (
	github.com/ArtefactGitHub/Go_P_Zo/internal/platform/myerror v0.0.0-20211226064812-51c406b24abe // indirect
	github.com/ArtefactGitHub/Go_P_Zo/internal/platform/myhttp v0.0.0-20211226064812-51c406b24abe // indirect
	github.com/ArtefactGitHub/Go_P_Zo/pkg/common v0.0.0-00010101000000-000000000000 // indirect
	github.com/julienschmidt/httprouter v1.3.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)

replace github.com/ArtefactGitHub/Go_P_Zo/internal/platform/myerror => ../internal/platform/myerror

replace github.com/ArtefactGitHub/Go_P_Zo/internal/platform/myhttp => ../internal/platform/myhttp

replace github.com/ArtefactGitHub/Go_P_Zo/internal/platform/mydb => ../internal/platform/mydb

replace github.com/ArtefactGitHub/Go_P_Zo/internal/platform/myrouter => ../internal/platform/myrouter

replace github.com/ArtefactGitHub/Go_P_Zo/internal/api/v1/zo => ../internal/api/v1/zo

replace github.com/ArtefactGitHub/Go_P_Zo/internal/api/v1/user => ../internal/api/v1/user

replace github.com/ArtefactGitHub/Go_P_Zo/internal/test => ../internal/test

replace github.com/ArtefactGitHub/Go_P_Zo/pkg/common => ../pkg/common
