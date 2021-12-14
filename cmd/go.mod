module github.com/ArtefactGitHub/Go_P_Zo/main

go 1.17

replace github.com/ArtefactGitHub/Go_P_Zo/internal/config => ../internal/config

require (
	github.com/ArtefactGitHub/Go_P_Zo/internal/config v0.0.0-00010101000000-000000000000
	github.com/ArtefactGitHub/Go_P_Zo/internal/controllers v0.0.0-00010101000000-000000000000
	github.com/ArtefactGitHub/Go_P_Zo/internal/models v0.0.0-00010101000000-000000000000
	github.com/go-sql-driver/mysql v1.6.0
)

require (
	github.com/ArtefactGitHub/Go_P_Zo/internal/models/zo v0.0.0-00010101000000-000000000000 // indirect
	github.com/ArtefactGitHub/Go_P_Zo/internal/services v0.0.0-00010101000000-000000000000 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)

replace github.com/ArtefactGitHub/Go_P_Zo/internal/controllers => ../internal/controllers

replace github.com/ArtefactGitHub/Go_P_Zo/internal/services => ../internal/services

replace github.com/ArtefactGitHub/Go_P_Zo/internal/models => ../internal/models

replace github.com/ArtefactGitHub/Go_P_Zo/internal/models/zo => ../internal/models/zo
