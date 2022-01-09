module github.com/ArtefactGitHub/Go_P_Zo/internal/api/v1/zo

go 1.17

replace github.com/ArtefactGitHub/Go_P_Zo/internal/config => ../../../config

replace github.com/ArtefactGitHub/Go_P_Zo/internal/platform/myhttp => ../../../platform/myhttp

replace github.com/ArtefactGitHub/Go_P_Zo/internal/platform/myrouter => ../../../platform/myrouter

replace github.com/ArtefactGitHub/Go_P_Zo/internal/platform/mydb => ../../../platform/mydb

replace github.com/ArtefactGitHub/Go_P_Zo/internal/platform/myerror => ../../../platform/myerror

replace github.com/ArtefactGitHub/Go_P_Zo/pkg/common => ../../../../pkg/common

replace github.com/ArtefactGitHub/Go_P_Zo/internal/test => ../../../test

require (
	github.com/ArtefactGitHub/Go_P_Zo/internal/platform/mydb v0.0.0-20220108142333-59a19c0068a6
	github.com/ArtefactGitHub/Go_P_Zo/internal/platform/myhttp v0.0.0-20220108142333-59a19c0068a6
	github.com/ArtefactGitHub/Go_P_Zo/internal/platform/myrouter v0.0.0-20220108142333-59a19c0068a6
	github.com/ArtefactGitHub/Go_P_Zo/internal/test v0.0.0-20220108142333-59a19c0068a6
	github.com/ArtefactGitHub/Go_P_Zo/pkg/common v0.0.0-20220108142333-59a19c0068a6
)

require (
	github.com/ArtefactGitHub/Go_P_Zo/internal/config v0.0.0-20220108142333-59a19c0068a6 // indirect
	github.com/ArtefactGitHub/Go_P_Zo/internal/platform/myerror v0.0.0-20220108142333-59a19c0068a6 // indirect
	github.com/go-sql-driver/mysql v1.6.0 // indirect
	github.com/julienschmidt/httprouter v1.3.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)
