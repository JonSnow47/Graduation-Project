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
	Id      bson.ObjectId `bson:"_id,omitempty"` // 文章 Id
	Title   string        `bson:"Title"`         // 标题
	Author  string        `bson:"Author"`        // 作者
	Brief   string        `bson:"Brief"`         // 简介
	Content string        `bson:"Content"`       // 内容
	Tags    []string      `bson:"Tags"`          // 标签Id
	Img     string        `bson:"Img"`           // 图片位置
	Views   int64         `bson:"Views"`         // 浏览次数
	Created time.Time     `bson:"Created"`       // 创建时间
	State   int8          `bson:"State"`         // 状态: 删除(-2),不可浏览(-1),未审查(0),可浏览(1),热门(2)
}

// connect to mongodb.
func CollectionArticle() mongo.Mongodb {
	m := mongo.ConnectMongo(consts.CollectionArticle)
	m.C.EnsureIndex(mgo.Index{
		Key:        []string{"Title"},
		Unique:     false,
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

	// 生成 ObjectId
	a.Id = bson.NewObjectId()
	// 匿名作者
	if a.Author == "" {
		a.Author = "Unknown"
	}
	// 将 base64 内容译码
	//content, err := base64.StdEncoding.DecodeString(a.Content)
	//a.Content = string(content)
	a.Created = time.Now()
	a.State = consts.Approverd

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

	err := m.C.UpdateId(bson.ObjectIdHex(id), bson.M{"$set": bson.M{"State": consts.Deleted}})
	return err

}

// Update modify article.
func (*articleServiceProvider) Update(a *Article) (err error) {
	m := CollectionArticle()
	defer m.S.Close()

	doc := bson.M{"$set": bson.M{
		"Title":   a.Title,
		"Author":  a.Author,
		"Brief":   a.Brief,
		"Content": a.Content,
		"Tags":    a.Tags,
	}}
	err = m.C.UpdateId(a.Id, doc)
	return
}

// ModifyStatus modify article's status.
func (*articleServiceProvider) ModifyState(id string, status int) error {
	m := CollectionArticle()
	defer m.S.Close()

	err := m.C.UpdateId(bson.ObjectIdHex(id), bson.M{"$set": bson.M{"State": status}})
	return err
}

// Get use id to get a article's detail.
func (*articleServiceProvider) Get(id string) (a Article, err error) {
	m := CollectionArticle()
	defer m.S.Close()

	err = m.C.FindId(bson.ObjectIdHex(id)).One(&a)
	if err != nil {
		return
	}

	err = m.C.UpdateId(bson.ObjectIdHex(id), bson.M{"$set": bson.M{"Views": a.Views + 1}})
	return
}

// All only admin can use this model to view all articles.
func (*articleServiceProvider) All(page int) (articles []Article, err error) {
	m := CollectionArticle()
	defer m.S.Close()

	err = m.C.Find(nil).Limit(5).Skip(5 * page).Sort("-Created").All(&articles)
	if err != nil {
		return nil, err
	}
	return articles, err
}

// Approved return all approved articles.
func (*articleServiceProvider) Approved() (articles []Article, err error) {
	m := CollectionArticle()
	defer m.S.Close()

	q := bson.M{"State": consts.Approverd}
	err = m.C.Find(q).All(&articles)
	// 按页查找，每页5个
	// err = m.C.Find(q).Limit(5).Skip(5 * page).Sort("-Created").All(&articles)
	if err != nil {
		return nil, err
	}
	return articles, err
}

// ListCreated admin use this model to audit articles.
func (*articleServiceProvider) ListCreated(page int) (articles []Article, err error) {
	m := CollectionArticle()
	defer m.S.Close()

	q := bson.M{"State": consts.Created}
	err = m.C.Find(q).Limit(5).Skip(5 * page).Sort("-Created").All(&articles)
	if err != nil {
		return nil, err
	}
	return articles, err
}
