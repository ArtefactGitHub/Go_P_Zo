module github.com/ArtefactGitHub/Go_P_Zo/internal/test

go 1.17

replace github.com/ArtefactGitHub/Go_P_Zo/internal/config => ../config

replace github.com/ArtefactGitHub/Go_P_Zo/internal/platform/mydb => ../platform/mydb

replace github.com/ArtefactGitHub/Go_P_Zo/internal/platform/myrouter => ../platform/myrouter

replace github.com/ArtefactGitHub/Go_P_Zo/internal/platform/myhttp => ../platform/myhttp

replace github.com/ArtefactGitHub/Go_P_Zo/internal/platform/myerror => ../platform/myerror

replace github.com/ArtefactGitHub/Go_P_Zo/internal/api/zo => ../api/zo

require (
	github.com/ArtefactGitHub/Go_P_Zo/internal/config v0.0.0-20211224055413-3bb2477ab245
	github.com/ArtefactGitHub/Go_P_Zo/internal/platform/mydb v0.0.0-20211224055413-3bb2477ab245
	github.com/go-sql-driver/mysql v1.6.0
)

require gopkg.in/yaml.v2 v2.4.0 // indirect
