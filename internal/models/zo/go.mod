module github.com/ArtefactGitHub/Go_T_TestDBAccess/internal/models/zo

go 1.17

replace github.com/ArtefactGitHub/Go_T_TestDBAccess/internal/models => ../

replace github.com/ArtefactGitHub/Go_T_TestDBAccess/internal/config => ../../config

require github.com/ArtefactGitHub/Go_T_TestDBAccess/internal/models v0.0.0-00010101000000-000000000000

require github.com/ArtefactGitHub/Go_T_TestDBAccess/internal/config v0.0.0-00010101000000-000000000000 // indirect

require gopkg.in/yaml.v2 v2.4.0 // indirect
