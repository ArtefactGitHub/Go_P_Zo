module github.com/ArtefactGitHub/Go_P_Zo/internal/platform/myrouter

go 1.17

replace github.com/ArtefactGitHub/Go_P_Zo/internal/api/zo => ../../api/zo

replace github.com/ArtefactGitHub/Go_P_Zo/internal/platform/myhttp => ../myhttp

replace github.com/ArtefactGitHub/Go_P_Zo/internal/platform/mydb => ../mydb

replace github.com/ArtefactGitHub/Go_P_Zo/internal/config => ../../config

replace github.com/ArtefactGitHub/Go_P_Zo/internal/myerror => ../../myerror

require github.com/ArtefactGitHub/Go_P_Zo/internal/api/zo v0.0.0-00010101000000-000000000000

require (
	github.com/ArtefactGitHub/Go_P_Zo/internal/config v0.0.0-20211217011354-173852cfa445 // indirect
	github.com/ArtefactGitHub/Go_P_Zo/internal/myerror v0.0.0-00010101000000-000000000000 // indirect
	github.com/ArtefactGitHub/Go_P_Zo/internal/platform/mydb v0.0.0-00010101000000-000000000000 // indirect
	github.com/ArtefactGitHub/Go_P_Zo/internal/platform/myhttp v0.0.0-00010101000000-000000000000 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)
