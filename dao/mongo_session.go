package dao

import mgo "gopkg.in/mgo.v2"

var MgoSession *mgo.Session

func init() {
	if MgoSession == nil {
		var err error
		MgoSession, err = mgo.Dial("localhost")
		if err != nil {
			panic(err)
		}
	}
}
