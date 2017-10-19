package data

import (
	"database/sql"
	"sync"
)

var instance *TDRDatabase
var once sync.Once

type TDRDatabase struct {
	Db *sql.DB
}

func GetTDRDatabase() *TDRDatabase {
	once.Do(func() {
		instance = new(TDRDatabase)
		var err error
		instance.Db, err = sql.Open("mysql", "devuser:devuser@tcp(localhost:3306)/circlepix?parseTime=true")
		if err != nil {
			panic(err.Error())
		}
	})
	return instance
}

func (tdrDb TDRDatabase) close() {
	tdrDb.Db.Close()
}
