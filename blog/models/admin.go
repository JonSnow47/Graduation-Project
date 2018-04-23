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

const AdminCollection = "admin"

type adminServiceProvide struct{}

var AdminService *adminServiceProvide

// Admin represent admin information.
type Admin struct {
	Id      bson.ObjectId `bson:"_id,omitempty"`
	Name    string        `bson:"name"`
	Pwd     string        `bson:"pwd"`
	Created time.Time     `bson:"created"`
	State   bool          `bson:"state"`
}

func CollectionAdmin() *mgo.Session {
	url := conf.MongoURL + "/" + consts.Database

	s, err := mgo.Dial(url)
	if err != nil {
		panic(err)
	}

	s.SetMode(mgo.Monotonic, true)

	s.DB(consts.Database).C("admin").EnsureIndex(mgo.Index{
		Key:        []string{"name"},
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
	c := s.DB(consts.Database).C(AdminCollection)
	defer s.Close()

	if len(pwd) < 6 || len(pwd) > 20 {
		return "", errors.New("Password length error.")
	}

	id := bson.NewObjectId()

	b, err := util.GenerateHash(pwd)
	if err != nil {
		return "", err
	}
	pwd = string(b)

	a := &Admin{
		Id:      id,
		Name:    name,
		Pwd:     pwd,
		Created: time.Now(),
		State:   true,
	}

	err = c.Insert(a)
	if err != nil {
		return "", err
	}

	return id.Hex(), err
}

func (sp *adminServiceProvide) Login(name, pwd string) string {
	s := CollectionAdmin()
	c := s.DB(consts.Database).C(AdminCollection)
	defer s.Close()

	var a Admin
	q := bson.M{"name": name}
	err := c.Find(q).One(&a)
	if err != nil {
		return ""
	} else if !util.CompareHash([]byte(a.Pwd), pwd) {
		return ""
	}
	return a.Id.Hex()
}
