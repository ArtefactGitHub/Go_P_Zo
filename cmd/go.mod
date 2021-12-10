module github.com/ArtefactGitHub/Go_T_TestDBAccess/main

go 1.17

replace github.com/ArtefactGitHub/Go_T_TestDBAccess/internal/config => ../internal/config

require (
	github.com/ArtefactGitHub/Go_T_TestDBAccess/internal/config v0.0.0-00010101000000-000000000000
	github.com/go-sql-driver/mysql v1.6.0
)

require gopkg.in/yaml.v2 v2.4.0 // indirect
