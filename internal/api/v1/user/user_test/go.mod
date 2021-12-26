module github.com/ArtefactGitHub/Go_P_Zo/internal/api/v1/user/user_test

go 1.17

replace github.com/ArtefactGitHub/Go_P_Zo/internal/config => ../../../../config

replace github.com/ArtefactGitHub/Go_P_Zo/internal/platform/mydb => ../../../../platform/mydb

replace github.com/ArtefactGitHub/Go_P_Zo/internal/test => ../../../../test

replace github.com/ArtefactGitHub/Go_P_Zo/internal/api/v1/user => ../

require (
	github.com/ArtefactGitHub/Go_P_Zo/internal/api/v1/user v0.0.0-00010101000000-000000000000
	github.com/ArtefactGitHub/Go_P_Zo/internal/platform/mydb v0.0.0-20211224055413-3bb2477ab245
	github.com/ArtefactGitHub/Go_P_Zo/internal/test v0.0.0-00010101000000-000000000000
)

require (
	github.com/ArtefactGitHub/Go_P_Zo/internal/config v0.0.0-20211224055413-3bb2477ab245 // indirect
	github.com/go-sql-driver/mysql v1.6.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)
