package main

import (
	"fmt"

	"github.com/ArtefactGitHub/Go_T_TestDBAccess/pkg/config"
)

func main() {
	config, err := config.LoadConfig("../configs/config.yml")
	if err != nil {
		panic(err)
	}

	fmt.Printf("config: %v\n", config)
}
