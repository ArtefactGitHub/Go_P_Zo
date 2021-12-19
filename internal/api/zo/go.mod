module github.com/ArtefactGitHub/Go_P_Zo/internal/api/zo

go 1.17

replace github.com/ArtefactGitHub/Go_P_Zo/internal/config => ../../config

replace github.com/ArtefactGitHub/Go_P_Zo/internal/platform/myhttp => ../../platform/myhttp

replace github.com/ArtefactGitHub/Go_P_Zo/internal/platform/mydb => ../../platform/mydb

replace github.com/ArtefactGitHub/Go_P_Zo/internal/platform/myerror => ../../platform/myerror

replace github.com/ArtefactGitHub/Go_P_Zo/internal/test => ../../test

require (
	github.com/ArtefactGitHub/Go_P_Zo/internal/platform/mydb v0.0.0-00010101000000-000000000000
	github.com/ArtefactGitHub/Go_P_Zo/internal/platform/myhttp v0.0.0-00010101000000-000000000000
	github.com/ArtefactGitHub/Go_P_Zo/internal/test v0.0.0-00010101000000-000000000000
)

require (
	github.com/ArtefactGitHub/Go_P_Zo/internal/config v0.0.0-20211217011354-173852cfa445 // indirect
	github.com/ArtefactGitHub/Go_P_Zo/internal/platform/myerror v0.0.0-00010101000000-000000000000 // indirect
	github.com/go-sql-driver/mysql v1.6.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)
