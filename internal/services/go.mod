module github.com/ArtefactGitHub/Go_P_Zo/internal/services

go 1.17

require github.com/ArtefactGitHub/Go_P_Zo/internal/models/zo v0.0.0-00010101000000-000000000000

require (
	github.com/ArtefactGitHub/Go_P_Zo/internal/config v0.0.0-00010101000000-000000000000 // indirect
	github.com/ArtefactGitHub/Go_P_Zo/internal/models v0.0.0-00010101000000-000000000000 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)

replace github.com/ArtefactGitHub/Go_P_Zo/internal/models/zo => ../models/zo

replace github.com/ArtefactGitHub/Go_P_Zo/internal/models => ../models

replace github.com/ArtefactGitHub/Go_P_Zo/internal/config => ../config
