module github.com/ArtefactGitHub/Go_P_Zo/internal/api/v1/session

go 1.17

replace github.com/ArtefactGitHub/Go_P_Zo/internal/config => ../../../config

replace github.com/ArtefactGitHub/Go_P_Zo/internal/platform/mydb => ../../../platform/mydb

replace github.com/ArtefactGitHub/Go_P_Zo/internal/platform/mycontext => ../../../platform/mycontext

replace github.com/ArtefactGitHub/Go_P_Zo/internal/platform/myhttp => ../../../platform/myhttp

replace github.com/ArtefactGitHub/Go_P_Zo/internal/platform/myerror => ../../../platform/myerror

replace github.com/ArtefactGitHub/Go_P_Zo/internal/platform/session => ../../../platform/session

replace github.com/ArtefactGitHub/Go_P_Zo/internal/platform/myrouter => ../../../platform/myrouter

replace github.com/ArtefactGitHub/Go_P_Zo/pkg/common => ../../../../pkg/common

require (
	github.com/ArtefactGitHub/Go_P_Zo/internal/platform/mydb v0.0.0-00010101000000-000000000000
	github.com/ArtefactGitHub/Go_P_Zo/internal/platform/myhttp v0.0.0-00010101000000-000000000000
	github.com/ArtefactGitHub/Go_P_Zo/internal/platform/myrouter v0.0.0-00010101000000-000000000000
	github.com/ArtefactGitHub/Go_P_Zo/pkg/common v0.0.0-00010101000000-000000000000
	golang.org/x/crypto v0.0.0-20220112180741-5e0467b6c7ce
)

require (
	github.com/ArtefactGitHub/Go_P_Zo/internal/config v0.0.0-00010101000000-000000000000 // indirect
	github.com/ArtefactGitHub/Go_P_Zo/internal/platform/mycontext v0.0.0-00010101000000-000000000000 // indirect
	github.com/ArtefactGitHub/Go_P_Zo/internal/platform/myerror v0.0.0-00010101000000-000000000000 // indirect
	github.com/julienschmidt/httprouter v1.3.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)