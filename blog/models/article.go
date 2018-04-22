package models

import (
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/JonSnow47/Graduation-Project/blog/consts"
	"github.com/JonSnow47/Graduation-Project/blog/mongo"
)

const cname = "article"

type articleServiceProvider struct{}

var ArticleService *articleServiceProvider

// Article represent the article information.
type Article struct {
	Id      bson.ObjectId   `bson:"_id,omitempty"`
	Title   string          `bson:"title"`
	Author  string          `bson:"author"`
	Content string          `bson:"content"`
	TagsId  []bson.ObjectId `bson:"tagsid"`
	Img     string          `bson:"img"`
	Created time.Time       `bson:"created"`
	State   int8            `bson:"state"`
}

// connect to mongodb.
func CollectionArticle() mongo.Mongodb {
	m := mongo.ConnectMongo(consts.CollectionArticle)
	m.C.EnsureIndex(mgo.Index{
		Key:        []string{"title"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	})
	return m
}

// New insert a new article.
func (*articleServiceProvider) New(a *Article) (string, error) {
	m := CollectionArticle()
	defer m.S.Close()

	a.Id = bson.NewObjectId()
	a.Created = time.Now()
	err := m.C.Insert(a)
	if err != nil {
		return "", err
	}

	return a.Id.Hex(), nil
}

// Delete modify the article's status.
func (*articleServiceProvider) Delete(id string) error {
	m := CollectionArticle()
	defer m.S.Close()

	err := m.C.UpdateId(bson.ObjectIdHex(id), bson.M{"$set": bson.M{"state": consts.Deleted}})
	return err
}

// Update modify article.
func (*articleServiceProvider) Update(a *Article) (err error) {
	m := CollectionArticle()
	defer m.S.Close()

	doc := bson.M{"$set": bson.M{
		"title":   a.Title,
		"author":  a.Author,
		"tagsid":  a.TagsId,
		"content": a.Content,
	}}
	err = m.C.UpdateId(a.Id, doc)
	return
}

// ModifyStatus modify article's status.
func (*articleServiceProvider) ModifyState(id string, status int) error {
	m := CollectionArticle()
	defer m.S.Close()

	err := m.C.UpdateId(bson.ObjectIdHex(id), bson.M{"$set": bson.M{"state": status}})
	return err
}

// GetOne use id to get a article.
func (*articleServiceProvider) Get(id string) (a Article, err error) {
	m := CollectionArticle()
	defer m.S.Close()

	err = m.C.FindId(bson.ObjectIdHex(id)).One(&a)
	return
}

// All only admin can use this model to view all articles.
func (*articleServiceProvider) All(page int) (articles []Article, err error) {
	m := CollectionArticle()
	defer m.S.Close()

	err = m.C.Find(nil).Limit(5).Skip(5 * page).All(&articles)
	if err != nil {
		return nil, err
	}
	return articles, err
}

// Approved return all approved articles.
func (*articleServiceProvider) Approved(page int) (articles []Article, err error) {
	m := CollectionArticle()
	defer m.S.Close()

	q := bson.M{"state": consts.Approverd}
	err = m.C.Find(q).Limit(5).Skip(5 * page).All(&articles)
	if err != nil {
		return nil, err
	}
	return articles, err
}

// ListCreated admin use this model to audit articles.
func (*articleServiceProvider) ListCreated(page int) (articles []Article, err error) {
	m := CollectionArticle()
	defer m.S.Close()

	q := bson.M{"state": consts.Created}
	err = m.C.Find(q).Limit(5).Skip(5 * page).All(&articles)
	if err != nil {
		return nil, err
	}
	return articles, err
}
