package models

import (
	"errors"
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/JonSnow47/Graduation-Project/blog/conf"
	"github.com/JonSnow47/Graduation-Project/blog/consts"
	"github.com/JonSnow47/Graduation-Project/blog/util"
)

const collectionAdmin = "admin"

type adminServiceProvide struct{}

var AdminService *adminServiceProvide

// Admin represent admin information.
type Admin struct {
	Id      bson.ObjectId `bson:"_id,omitempty"`
	Name    string        `bson:"Name"`
	Pwd     string        `bson:"Pwd"`
	Created time.Time     `bson:"Created"`
	State   bool          `bson:"State"`
}

func CollectionAdmin() *mgo.Session {
	url := conf.MongoURL + "/" + consts.Database

	s, err := mgo.Dial(url)
	if err != nil {
		panic(err)
	}

	s.SetMode(mgo.Monotonic, true)

	s.DB(consts.Database).C("admin").EnsureIndex(mgo.Index{
		Key:        []string{"Name"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	})
	return s
}

// New create a new admin.
func (sp *adminServiceProvide) New(name, pwd string) (string, error) {
	s := CollectionAdmin()
	c := s.DB(consts.Database).C(collectionAdmin)
	defer s.Close()

	if len(pwd) < 6 || len(pwd) > 20 {
		return "", errors.New("Password length error.")
	}

	b, err := util.GenerateHash(pwd)
	if err != nil {
		return "", err
	}

	a := &Admin{
		Id:      bson.NewObjectId(),
		Name:    name,
		Pwd:     string(b),
		Created: time.Now(),
		State:   true,
	}

	err = c.Insert(a)
	if err != nil {
		return "", err
	}

	return a.Id.Hex(), err
}

func (sp *adminServiceProvide) Login(name, pwd string) string {
	s := CollectionAdmin()
	c := s.DB(consts.Database).C(collectionAdmin)
	defer s.Close()

	var a Admin
	q := bson.M{"Name": name}
	err := c.Find(q).One(&a)
	if err != nil {
		return ""
	} else if !util.CompareHash([]byte(a.Pwd), pwd) {
		return ""
	}
	return a.Id.Hex()
}
