package main

import (
	"fmt"

	"github.com/lempiy/chat_go/models"
	_ "github.com/lempiy/chat_go/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func init() {
	orm.RegisterDriver("postgres", orm.DRPostgres)
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		"anton", "q1w2e3r4", "chat_db")

	err := orm.RegisterDataBase("default", "postgres", dbinfo)
	beego.Info("DB errors", err)
}

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	models.NewORM()
	beego.Run()
}
