module github.com/ArtefactGitHub/Go_P_Zo/internal/models/zo

go 1.17

replace github.com/ArtefactGitHub/Go_P_Zo/internal/myerror => ../../myerror

replace github.com/ArtefactGitHub/Go_P_Zo/internal/platform => ../../platform

replace github.com/ArtefactGitHub/Go_P_Zo/internal/models => ../

replace github.com/ArtefactGitHub/Go_P_Zo/internal/config => ../../config

require github.com/ArtefactGitHub/Go_P_Zo/internal/config v0.0.0-00010101000000-000000000000 // indirect

require (
	github.com/ArtefactGitHub/Go_P_Zo/internal/models v0.0.0-00010101000000-000000000000
	github.com/ArtefactGitHub/Go_P_Zo/internal/platform v0.0.0-00010101000000-000000000000
)

require (
	github.com/ArtefactGitHub/Go_P_Zo/internal/myerror v0.0.0-00010101000000-000000000000 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)
