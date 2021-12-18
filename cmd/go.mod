module github.com/ArtefactGitHub/Go_P_Zo/main

go 1.17

replace github.com/ArtefactGitHub/Go_P_Zo/internal/config => ../internal/config

require (
	github.com/ArtefactGitHub/Go_P_Zo/internal/api/zo v0.0.0-00010101000000-000000000000
	github.com/ArtefactGitHub/Go_P_Zo/internal/config v0.0.0-20211217011354-173852cfa445
	github.com/ArtefactGitHub/Go_P_Zo/internal/models v0.0.0-20211217011354-173852cfa445
	github.com/go-sql-driver/mysql v1.6.0
)

require (
	github.com/ArtefactGitHub/Go_P_Zo/internal/myerror v0.0.0-00010101000000-000000000000 // indirect
	github.com/ArtefactGitHub/Go_P_Zo/internal/platform v0.0.0-00010101000000-000000000000 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)

replace github.com/ArtefactGitHub/Go_P_Zo/internal/models => ../internal/models

replace github.com/ArtefactGitHub/Go_P_Zo/internal/myerror => ../internal/myerror

replace github.com/ArtefactGitHub/Go_P_Zo/internal/platform => ../internal/platform

replace github.com/ArtefactGitHub/Go_P_Zo/internal/api/zo => ../internal/api/zo
