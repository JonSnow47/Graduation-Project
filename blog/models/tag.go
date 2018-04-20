package models

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/JonSnow47/Graduation-Project/blog/consts"
	"github.com/JonSnow47/Graduation-Project/blog/initialize"
)

type tagServiceProvider struct{}

var TagService *tagServiceProvider

func CollectionTag() (c *mgo.Collection) {
	c = initialize.ConnectMongo(consts.CollectionTag)
	c.EnsureIndex(mgo.Index{
		Key:        []string{"tag"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	})
	return
}

// Tag represents the tag of article.
type Tag struct {
	Id     bson.ObjectId `bson:"_id,omitempty"`
	Name   string        `bson:"name"`
	Count  int           `bson:"count"`
	Status bool          `bson:"status"`
}

// New insert a new article.
func (*tagServiceProvider) New(tag string) (string, error) {
	c := CollectionTag()
	t := &Tag{
		Id:   bson.NewObjectId(),
		Name: tag,
	}
	err := c.Insert(t)
	return t.Id.Hex(), err
}

// Delete modify tag's status.
func (*tagServiceProvider) Delete(id string) error {
	c := CollectionTag()
	err := c.UpdateId(id, bson.M{"$set": bson.M{"status": false}})
	return err
}

// Enable modify tag's status.
func (*tagServiceProvider) Enable(id string) error {
	c := CollectionTag()
	err := c.UpdateId(id, bson.M{"$set": bson.M{"status": true}})
	return err
}

// Get one tag's info.
func (*tagServiceProvider) Get(id string) (t *Tag, err error) {
	c := CollectionTag()
	err = c.FindId(id).One(t)
	return
}

// All get all tags list.
func (*tagServiceProvider) All(page int) (tags []*Tag, err error) {
	c := CollectionTag()
	err = c.Find(nil).Limit(20).Skip(20 * page).All(tags)
	return
}
