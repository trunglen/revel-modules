package app

import (
	mgo "gopkg.in/mgo.v2"

	"github.com/revel/revel"
)

// Database connection variables
var (
	Db     *mgo.Database
	Driver string
	Spec   string
)

// Init method used to initialize DB module on `OnAppStart`
func InitDB() {
	// Read configuration.
	MaxPool = revel.Config.IntDefault("mongo.maxPool", 0)
	PATH, _ = revel.Config.String("mongo.path")
	DBNAME, _ = revel.Config.String("mongo.database")
	CheckAndInitServiceConnection()
}

func init() {
	revel.RegisterModuleInit(func(module *revel.Module) {
		InitDB()
	})
}
