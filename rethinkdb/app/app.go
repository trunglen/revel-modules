package app

import (
	"github.com/revel/revel"
)

var (
	DB *RethinkDB
)

func InitDB() {
	if DB == nil {
		DB = NewRethinkDB()
	}
}
func init() {
	revel.RegisterModuleInit(func(module *revel.Module) {
		InitDB()
	})
}
