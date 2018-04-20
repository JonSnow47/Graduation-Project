package controllers

import (
	"encoding/json"
	"log"

	"github.com/astaxie/beego"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/JonSnow47/Graduation-Project/blog/consts"
	"github.com/JonSnow47/Graduation-Project/blog/models"
)

type ArticleController struct {
	beego.Controller
}

// New create a new article.
func (c *ArticleController) New() {
	var req struct {
		Title   string          `json:"title" validate:"required"`
		Author  string          `json:"author"`
		Content string          `json:"content" validate:"required"`
		TagsId  []bson.ObjectId `json:"tagsid"`
		Img     string          `json:"img"`
	}

	err := json.Unmarshal(c.Ctx.Input.RequestBody, &req)
	if err != nil {
		log.Println(consts.ErrParam, err)
		c.Data["json"] = map[string]interface{}{consts.ErrParam: err}
		return
	}

	a := &models.Article{
		Title:   req.Title,
		Author:  req.Author,
		Content: req.Content,
		TagsId:  req.TagsId,
	}

	id, err := models.ArticleService.New(a)
	if err != nil {
		log.Println("Create article failed:", err)
		c.Data["json"] = map[string]interface{}{"Create article failed:": err}
		return
	}

	c.Data["json"] = map[string]string{"ObjectId": id}
	c.ServeJSON()
}

// Delete
func (c *ArticleController) Delete() {
	var req struct {
		Id string `json:"id"`
	}
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &req)
	if err != nil {
		log.Println(consts.ErrParam, err)
		c.Data["json"] = map[string]interface{}{consts.ErrParam: err}
		return
	}

	if err = models.ArticleService.Delete(req.Id); err != nil {
		log.Println(consts.ErrMongo, err)
		c.Data["json"] = map[string]interface{}{"status": err}
	}

	c.Data["json"] = map[string]interface{}{"status": err}
	c.ServeJSON()
}

func (c *ArticleController) Update() {
	var req struct {
		Id      string          `json:"id" validate:"required"`
		Title   string          `json:"title" validate:"required"`
		Author  string          `json:"author"`
		Content string          `json:"content" validate:"required"`
		TagsId  []bson.ObjectId `json:"tagsid"`
		Img     string          `json:"img"`
	}

	err := json.Unmarshal(c.Ctx.Input.RequestBody, &req)
	if err != nil {
		log.Println(consts.ErrParam, err)
		c.Data["json"] = map[string]interface{}{consts.ErrParam: err}
		return
	}

	a := &models.Article{
		Id:      bson.ObjectIdHex(req.Id),
		Title:   req.Title,
		Author:  req.Author,
		Content: req.Content,
		TagsId:  req.TagsId,
		Img:     req.Img,
	}
	err = models.ArticleService.Update(a)
	if err != nil {
		log.Println()
		c.Data["json"] = err.Error()
		return
	}
	c.Data["json"] = map[string]string{"status": "update success!"}
	c.ServeJSON()
}

func (c *ArticleController) Get() {
	var req struct {
		Id string `json:"id" validate:"required"`
	}
	// objectId := c.Ctx.Input.Param(":objectId")
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &req)
	if err != nil {
		log.Println(consts.ErrParam, err)
		c.Data["json"] = map[string]interface{}{consts.ErrParam: err}
		return
	}

	a, err := models.ArticleService.GetOne(req.Id)
	if err != nil {
		if err == mgo.ErrNotFound {
			c.Data["json"] = map[string]interface{}{"status": consts.Success, "data": ""}
		} else {
			log.Println("Mongodb error:", err)
			c.Data["json"] = err.Error()
		}
	}

	c.Data["json"] = map[string]interface{}{"status": consts.Success, "data": *a}
	c.ServeJSON()
}

func (c *ArticleController) All() {
	var req struct {
		Page int `json:"page"`
	}
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &req)
	if err != nil {
		log.Println(consts.ErrParam, err)
		c.Data["json"] = map[string]interface{}{consts.ErrParam: err}
		return
	}

	articles, err := models.ArticleService.Approved(req.Page)
	if err != nil {
		log.Println(consts.ErrMongo, err)
		c.Data["json"] = map[string]interface{}{"status": err}
		return
	}

	c.Data["json"] = map[string]interface{}{"status": consts.Success, "data": articles}
	c.ServeJSON()
}

func (c *ArticleController) Approved() {
	var req struct {
		Page int `json:"page"`
	}
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &req)
	if err != nil {
		log.Println(consts.ErrParam, err)
		c.Data["json"] = map[string]interface{}{consts.ErrParam: err}
		return
	}

	articles, err := models.ArticleService.Approved(req.Page)
	if err != nil {
		if err == mgo.ErrNotFound {
			c.Data["json"] = map[string]interface{}{"status": err}
			return
		}
		log.Println("Mongodb error:", err)
		return
	}
	c.Data["json"] = map[string]interface{}{"status": consts.Success, "data": articles}
	c.ServeJSON()
}
