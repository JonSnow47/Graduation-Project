package models

import (
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/JonSnow47/Graduation-Project/blog/conf"
	"github.com/JonSnow47/Graduation-Project/blog/consts"
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
	Status  int8            `bson:"status"`
}

func initMgo() *mgo.Session {
	url := conf.MongoURL + "/" + consts.Database

	s, err := mgo.Dial(url)
	if err != nil {
		panic(err)
	}

	s.SetMode(mgo.Monotonic, true)

	s.DB(consts.Database).C(consts.CollectionArticle).EnsureIndex(mgo.Index{
		Key:        []string{"title"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	})
	return s
}

// New insert a new article.
func (*articleServiceProvider) New(a *Article) (string, error) {
	s := initMgo()
	defer s.Close()
	c := s.DB(consts.Database).C(cname)

	a.Created = time.Now()
	err := c.Insert(a)
	if err != nil {
		return "", err
	}

	var id string
	err = c.Find(a).One(id)
	return id, nil
}

// Delete modify the article's status.
func (*articleServiceProvider) Delete(id string) error {
	s := initMgo()
	defer s.Close()
	c := s.DB(consts.Database).C(cname)

	err := c.UpdateId(id, bson.M{"$set": bson.M{"status": consts.Deleted}})
	return err
}

// Update modify article.
func (*articleServiceProvider) Update(a *Article) (err error) {
	s := initMgo()
	defer s.Close()
	c := s.DB(consts.Database).C(cname)

	doc := bson.M{"$set": bson.M{
		"title":   a.Title,
		"author":  a.Author,
		"tagsid":  a.TagsId,
		"content": a.Content,
	}}
	err = c.Update(bson.M{"_id": a.Id}, doc)
	return
}

// UpdateStatus
func (*articleServiceProvider) UpdateStatus(id string, result int8) error {
	s := initMgo()
	defer s.Close()
	c := s.DB(consts.Database).C(cname)

	err := c.UpdateId(id, bson.M{"$set": bson.M{"status": result}})
	return err
}

func (*articleServiceProvider) GetOne(id string) (a *Article, err error) {
	s := initMgo()
	defer s.Close()
	c := s.DB(consts.Database).C(cname)

	err = c.FindId(id).One(a)
	if err != nil {
		return nil, err
	}
	return a, nil
}

func (*articleServiceProvider) All(page int) (articles []*Article, err error) {
	s := initMgo()
	defer s.Close()
	c := s.DB(consts.Database).C(cname)

	err = c.Find(nil).Limit(5).Skip(5 * page).All(articles)
	if err != nil {
		return nil, err
	}
	return articles, err
}

func (*articleServiceProvider) Approved(page int) (articles []*Article, err error) {
	s := initMgo()
	defer s.Close()
	c := s.DB(consts.Database).C(cname)

	q := bson.M{"status": consts.Approverd}
	err = c.Find(q).Limit(5).Skip(5 * page).All(articles)
	if err != nil {
		return nil, err
	}
	return articles, err
}

func (*articleServiceProvider) ListCreated(page int) (articles []*Article, err error) {
	s := initMgo()
	defer s.Close()
	c := s.DB(consts.Database).C(cname)

	q := bson.M{"status": consts.Created}
	err = c.Find(q).Limit(5).Skip(5 * page).All(articles)
	if err != nil {
		return nil, err
	}
	return articles, err
}
