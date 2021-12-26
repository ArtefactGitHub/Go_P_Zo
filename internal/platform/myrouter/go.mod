module github.com/ArtefactGitHub/Go_P_Zo/internal/platform/myrouter

go 1.17

replace github.com/ArtefactGitHub/Go_P_Zo/internal/api/v1/zo => ../../api/v1/zo

replace github.com/ArtefactGitHub/Go_P_Zo/internal/api/v1/user => ../../api/v1/user

replace github.com/ArtefactGitHub/Go_P_Zo/internal/platform/myhttp => ../myhttp

replace github.com/ArtefactGitHub/Go_P_Zo/internal/platform/mydb => ../mydb

replace github.com/ArtefactGitHub/Go_P_Zo/internal/config => ../../config

replace github.com/ArtefactGitHub/Go_P_Zo/internal/platform/myerror => ../../platform/myerror

require (
	github.com/ArtefactGitHub/Go_P_Zo/internal/api/v1/user v0.0.0-20211226064812-51c406b24abe
	github.com/ArtefactGitHub/Go_P_Zo/internal/api/v1/zo v0.0.0-20211226064812-51c406b24abe
	github.com/julienschmidt/httprouter v1.3.0
)

require (
	github.com/ArtefactGitHub/Go_P_Zo/internal/config v0.0.0-20211226064812-51c406b24abe // indirect
	github.com/ArtefactGitHub/Go_P_Zo/internal/platform/mydb v0.0.0-20211226064812-51c406b24abe // indirect
	github.com/ArtefactGitHub/Go_P_Zo/internal/platform/myerror v0.0.0-20211226064812-51c406b24abe // indirect
	github.com/ArtefactGitHub/Go_P_Zo/internal/platform/myhttp v0.0.0-20211226064812-51c406b24abe // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)
