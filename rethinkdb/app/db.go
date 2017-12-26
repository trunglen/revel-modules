package app

import (
	"github.com/golang/glog"
	"github.com/revel/revel"
	r "gopkg.in/dancannon/gorethink.v2"
)

type RethinkDB struct {
	Session *r.Session
	dbName  string
}

func (db *RethinkDB) GetTable(name string) r.Term {
	return r.DB(db.dbName).Table(name)
}
func NewRethinkDB() *RethinkDB {
	addr, _ := revel.Config.String("rethinkdb.address")
	initialCap, _ := revel.Config.Int("rethinkdb.initial_cap")
	maxOpen, _ := revel.Config.Int("rethinkdb.max_open")
	database, _ := revel.Config.String("rethinkdb.database")
	var err error
	session, err := r.Connect(r.ConnectOpts{
		Address:    addr,
		InitialCap: initialCap,
		MaxOpen:    maxOpen,
		Database:   database,
	})
	if err != nil {
		revel.AppLog.Error("loi rethinkdb " + err.Error())
	}
	return &RethinkDB{
		Session: session,
		dbName:  database,
	}
}

func (db *RethinkDB) QueryBuilder() r.Term {
	if db == nil {
		glog.Fatal("nil db instance")
	}
	return r.DB(db.dbName)
}

func (db *RethinkDB) IsErrEmpty(err error) bool {
	return err == r.ErrEmptyResult
}

func (db *RethinkDB) Table(name string) r.Term {
	var cursor, err = r.DB(db.dbName).TableList().Run(db.Session)
	if err != nil {
		panic(err)
	}
	var names []string
	if err := cursor.All(&names); err != nil {
		panic(err)
	}
	for _, table := range names {
		if table == name {
			return db.GetTable(name)
		}
	}

	// create table
	{
		var _, err = r.DB(db.dbName).TableCreate(name).RunWrite(db.Session)
		if err != nil {
			panic(err)
		}
		return db.GetTable(name)
	}

}
