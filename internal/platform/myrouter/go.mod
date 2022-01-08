module github.com/ArtefactGitHub/Go_P_Zo/internal/platform/myrouter

go 1.17

replace github.com/ArtefactGitHub/Go_P_Zo/internal/api/v1/zo => ../../api/v1/zo

replace github.com/ArtefactGitHub/Go_P_Zo/internal/api/v1/user => ../../api/v1/user

replace github.com/ArtefactGitHub/Go_P_Zo/internal/platform/myhttp => ../myhttp

replace github.com/ArtefactGitHub/Go_P_Zo/internal/platform/mydb => ../mydb

replace github.com/ArtefactGitHub/Go_P_Zo/internal/config => ../../config

replace github.com/ArtefactGitHub/Go_P_Zo/internal/platform/myerror => ../../platform/myerror

replace github.com/ArtefactGitHub/Go_P_Zo/pkg/common => ../../../pkg/common

require (
	github.com/ArtefactGitHub/Go_P_Zo/pkg/common v0.0.0-00010101000000-000000000000
	github.com/julienschmidt/httprouter v1.3.0
)
