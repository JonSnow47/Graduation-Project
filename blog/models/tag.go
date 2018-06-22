package models

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/JonSnow47/Graduation-Project/blog/mongo"
)

const collectionTag = "tag"

type tagServiceProvider struct{}

var TagService *tagServiceProvider

/*func init() (session *mgo.Session) {
	session, err := mgo.Dial(conf.MongoURL + "/" + consts.Database)
	if err != nil {
		panic(err)
	}
	session.SetMode(mgo.Monotonic, true)
	session.DB(consts.Database).C(consts.CollectionTag).EnsureIndex(mgo.Index{
		Key:        []string{"tag"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	})

	return
}*/

func CollectionTag() mongo.Mongodb {
	m := mongo.ConnectMongo(collectionTag)
	m.C.EnsureIndex(mgo.Index{
		Key:        []string{"_id"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	})
	return m
}

// Tag represents the tag of article.
type Tag struct {
	Id    string `bson:"_id"`   // 标签的Id即为标签
	Count int    `bson:"Count"` // 计数，拥有这个标签的文章数
	State bool   `bson:"State"` // 是否可用
}

// New insert a new article.
func (*tagServiceProvider) New(tag string) (err error) {
	m := CollectionTag()
	defer m.S.Close()

	t := &Tag{
		Id:    tag,
		State: true,
	}
	err = m.C.Insert(t)
	return
}

// Delete modify tag's status.
func (*tagServiceProvider) Delete(id string) error {
	m := CollectionTag()
	defer m.S.Close()

	err := m.C.UpdateId(id, bson.M{"$set": bson.M{"State": false}})
	return err
}

// Enable modify tag's status.
func (*tagServiceProvider) Enable(id string) error {
	m := CollectionTag()
	defer m.S.Close()

	err := m.C.UpdateId(id, bson.M{"$set": bson.M{"State": true}})
	return err
}

// Get Tag.Name.
func (*tagServiceProvider) Get(id string) (t *Tag, err error) {
	m := CollectionTag()
	defer m.S.Close()

	err = m.C.FindId(id).One(&t)
	return
}

// All get all tags list.
func (*tagServiceProvider) All() (tags []*Tag, err error) {
	m := CollectionTag()
	defer m.S.Close()

	err = m.C.Find(nil).All(&tags)
	return tags, err
}

func (*tagServiceProvider) Count(tags []string) (err error) {
	m := CollectionTag()
	defer m.S.Close()

	for i, _ := range tags {
		err = m.C.UpdateId(tags[i], bson.M{"$inc": bson.M{"Count": +1}})
		if err != nil {
			return err
		}
	}
	return
}
