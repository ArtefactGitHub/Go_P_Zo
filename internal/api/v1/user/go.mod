module github.com/ArtefactGitHub/Go_P_Zo/internal/api/v1/user

go 1.17

replace github.com/ArtefactGitHub/Go_P_Zo/internal/config => ../../../config

replace github.com/ArtefactGitHub/Go_P_Zo/internal/platform/mydb => ../../../platform/mydb

require github.com/ArtefactGitHub/Go_P_Zo/internal/platform/mydb v0.0.0-00010101000000-000000000000

require (
	github.com/ArtefactGitHub/Go_P_Zo/internal/config v0.0.0-00010101000000-000000000000 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)
