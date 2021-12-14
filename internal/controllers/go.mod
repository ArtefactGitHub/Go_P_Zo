module github.com/ArtefactGitHub/Go_P_Zo/internal/controllers

go 1.17

replace github.com/ArtefactGitHub/Go_P_Zo/internal/config => ../config

replace github.com/ArtefactGitHub/Go_P_Zo/internal/services => ../services

replace github.com/ArtefactGitHub/Go_P_Zo/internal/models => ../models

replace github.com/ArtefactGitHub/Go_P_Zo/internal/models/zo => ../models/zo

require (
	github.com/ArtefactGitHub/Go_P_Zo/internal/models/zo v0.0.0-00010101000000-000000000000
	github.com/ArtefactGitHub/Go_P_Zo/internal/services v0.0.0-00010101000000-000000000000
)

require (
	github.com/ArtefactGitHub/Go_P_Zo/internal/config v0.0.0-00010101000000-000000000000 // indirect
	github.com/ArtefactGitHub/Go_P_Zo/internal/models v0.0.0-00010101000000-000000000000 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)
