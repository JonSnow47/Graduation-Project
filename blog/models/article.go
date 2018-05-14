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
	Id      bson.ObjectId   `bson:"_id,omitempty"` // 文章 Id
	Title   string          `bson:"title"`         // 标题
	Author  string          `bson:"author"`        // 作者
	Content string          `bson:"content"`       // 内容
	TagsId  []bson.ObjectId `bson:"tagsid"`        // 标签Id
	Img     string          `bson:"img"`           // 图片位置
	Views   int64           `bson:"views"`         // 浏览次数
	Created time.Time       `bson:"created"`       // 创建时间
	State   int8            `bson:"state"`         // 状态: 删除(-2),不可浏览(-1),未审查(0),可浏览(1),热门(2)
}

// connect to mongodb.
func CollectionArticle() mongo.Mongodb {
	m := mongo.ConnectMongo(consts.CollectionArticle)
	m.C.EnsureIndex(mgo.Index{
		Key:        []string{"title"},
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
		a.Author = "Unknow"
	}
	// 将 base64 内容译码
	// content,err := base64.StdEncoding.DecodeString(a.Content)
	// a.Content = string(content)
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
	if err != nil {
		return
	}

	err = m.C.UpdateId(bson.ObjectIdHex(id), bson.M{"$set": bson.M{"views": a.Views + 1}})
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
func (*articleServiceProvider) Approved() (articles []Article, err error) {
	m := CollectionArticle()
	defer m.S.Close()

	q := bson.M{"state": consts.Approverd}
	err = m.C.Find(q).All(&articles)
	// 按页查找，每页5个
	// err = m.C.Find(q).Limit(5).Skip(5 * page).All(&articles)
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
