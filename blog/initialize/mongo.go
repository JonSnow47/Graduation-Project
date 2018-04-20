package initialize

import (
	"gopkg.in/mgo.v2"

	"github.com/JonSnow47/Graduation-Project/blog/conf"
	"github.com/JonSnow47/Graduation-Project/blog/consts"
)

type Mongodb struct {
	S *mgo.Session
	Db *mgo.Database
	Con *mgo.Collection
}

func ConnectMongo(collection string) (M Mongodb) {
	var err error
	url := conf.MongoURL + "/" + consts.Database
	M.S, err = mgo.Dial(url)
	if err != nil {
		panic(err)
	}

	M.S.SetMode(mgo.Monotonic, true)

	M.Db = M.S.DB(consts.Database)
	M.Con = M.Db.C(collection)
	return
}
