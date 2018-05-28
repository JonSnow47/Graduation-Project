package mongo

import (
	"gopkg.in/mgo.v2"

	"github.com/JonSnow47/Graduation-Project/blog/conf"
)

type Mongodb struct {
	S *mgo.Session
	D *mgo.Database
	C *mgo.Collection
}

func ConnectMongo(collection string) (M Mongodb) {
	var err error
	url := conf.MongoURL + "/" + conf.MongoDatabase
	M.S, err = mgo.Dial(url)
	if err != nil {
		panic(err)
	}

	M.S.SetMode(mgo.Monotonic, true)

	M.D = M.S.DB(conf.MongoDatabase)
	M.C = M.D.C(collection)
	return
}
