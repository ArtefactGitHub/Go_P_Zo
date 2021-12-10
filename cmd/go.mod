module github.com/Go_T_TestDBAccess/main

go 1.17

replace github.com/ArtefactGitHub/Go_T_TestDBAccess/pkg/config => ../pkg/config

require (
	github.com/go-sql-driver/mysql v1.6.0
	gopkg.in/yaml.v2 v2.4.0 // indirect
)

require github.com/ArtefactGitHub/Go_T_TestDBAccess/pkg/config v0.0.0-00010101000000-000000000000
