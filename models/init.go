package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
)

var O orm.Ormer

func init() {
	orm.RegisterModel(new(Object))
}

func NewORM() {
	O = orm.NewOrm()
	O.Using("default")
}
