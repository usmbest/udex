package services

import (
	"io/ioutil"

	"github.com/citypayorg/udex/backend/daos"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/dbtest"
)

var server dbtest.DBServer
var db *mgo.Session

func init() {
	temp, _ := ioutil.TempDir("", "test")
	server.SetPath(temp)

	session := server.Session()
	if _, err := daos.InitSession(session); err != nil {
		panic(err)
	}
	db = session
}
