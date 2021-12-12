package main

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/ArtefactGitHub/Go_T_TestDBAccess/internal/config"
	"github.com/ArtefactGitHub/Go_T_TestDBAccess/internal/models"
	"github.com/ArtefactGitHub/Go_T_TestDBAccess/internal/models/zo"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// 設定ファイルの取得
	config, err := config.LoadConfig("../configs/config.yml")
	if err != nil {
		panic(err)
	}
	fmt.Printf("config: %v\n", config)

	err = models.Init(config)
	if err != nil {
		panic(err)
	}
	defer models.Finalize()

	zos, err := zo.FindAll()
	if err != nil {
		panic(err)
	}
	for _, zo := range zos {
		fmt.Println(zo)
	}

	z1, _ := zo.Find(1)
	fmt.Println(z1)
	z2, _ := zo.Find(2)
	fmt.Println(z2)
	z3, _ := zo.Find(3)
	fmt.Println(z3)

	z1.Exp = 222
	z1.Message = "updated"
	err = z1.Update()
	if err != nil {
		panic(err)
	}
	fmt.Println("update success")

	ac, _ := time.Parse(layout, "2021-12-12")
	fmt.Println(ac)
	newz := zo.Zo{AchievementDate: ac, Exp: 1000, CategoryId: 0, Message: "inserted", CreatedAt: time.Now(), UpdatedAt: sql.NullTime{}}
	id, err := zo.Create(&newz)
	if err != nil {
		panic(err)
	}
	fmt.Println("insert success")
	fmt.Printf("inserted id: %v\n", id)

	fmt.Printf("newz id: %v\n", newz.Id)
	err = zo.Delete(newz)
	if err != nil {
		panic(err)
	}
	fmt.Println("delete success")
}

var layout = "2006-01-02"
