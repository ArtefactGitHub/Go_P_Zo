module github.com/ArtefactGitHub/Go_P_Zo/internal/platform/myrouter

go 1.17

replace github.com/ArtefactGitHub/Go_P_Zo/internal/api/zo => ../../api/zo

replace github.com/ArtefactGitHub/Go_P_Zo/internal/platform/myhttp => ../myhttp

replace github.com/ArtefactGitHub/Go_P_Zo/internal/platform/mydb => ../mydb

replace github.com/ArtefactGitHub/Go_P_Zo/internal/config => ../../config

replace github.com/ArtefactGitHub/Go_P_Zo/internal/platform/myerror => ../../platform/myerror

require (
	github.com/ArtefactGitHub/Go_P_Zo/internal/api/zo v0.0.0-20211224055413-3bb2477ab245
	github.com/julienschmidt/httprouter v1.3.0
)

require (
	github.com/ArtefactGitHub/Go_P_Zo/internal/config v0.0.0-20211224055413-3bb2477ab245 // indirect
	github.com/ArtefactGitHub/Go_P_Zo/internal/platform/mydb v0.0.0-20211224055413-3bb2477ab245 // indirect
	github.com/ArtefactGitHub/Go_P_Zo/internal/platform/myerror v0.0.0-20211224055413-3bb2477ab245 // indirect
	github.com/ArtefactGitHub/Go_P_Zo/internal/platform/myhttp v0.0.0-20211224055413-3bb2477ab245 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)
