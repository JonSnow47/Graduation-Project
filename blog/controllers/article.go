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
		c.Data["json"] = map[string]interface{}{consts.Stauts: consts.Failure, consts.Data: err}
		c.ServeJSON()
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
		c.Data["json"] = map[string]interface{}{consts.Stauts: consts.Failure, consts.Data: err}
		c.ServeJSON()
		return
	}

	c.Data["json"] = map[string]interface{}{consts.Stauts: consts.Success, consts.Data: map[string]string{"id": id}}
	c.ServeJSON()
}

// Delete modify article's state to -2.
func (c *ArticleController) Delete() {
	var req struct {
		Id string `json:"id"`
	}

	err := json.Unmarshal(c.Ctx.Input.RequestBody, &req)
	if err != nil {
		log.Println(consts.ErrParam, err)
		c.Data["json"] = map[string]interface{}{consts.Stauts: consts.Failure, consts.Data: err}
		c.ServeJSON()
		return
	}

	if err = models.ArticleService.Delete(req.Id); err != nil {
		log.Println(consts.ErrMongo, err)
		c.Data["json"] = map[string]interface{}{consts.Stauts: consts.Failure, consts.Data: err}
		c.ServeJSON()
		return
	}

	c.Data["json"] = map[string]interface{}{consts.Stauts: consts.Success}
	c.ServeJSON()
}

// Update update article.
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
		c.Data["json"] = map[string]interface{}{consts.Stauts: consts.Failure, consts.Data: err}
		c.ServeJSON()
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
		c.Data["json"] = map[string]interface{}{consts.Stauts: consts.Failure, consts.Data: err}
		c.ServeJSON()
		return
	}

	c.Data["json"] = map[string]interface{}{consts.Stauts: consts.Success}
	c.ServeJSON()
}

// ModifyState
func (c *ArticleController) ModifyState() {
	var req struct {
		Id    string `json:"id" validate:"required"`
		State int    `json:"state"`
	}

	err := json.Unmarshal(c.Ctx.Input.RequestBody, &req)
	if err != nil {
		log.Println(consts.ErrParam, err)
		c.Data["json"] = map[string]interface{}{consts.Stauts: consts.Failure, consts.Data: err}
		c.ServeJSON()
		return
	}

	err = models.ArticleService.ModifyState(req.Id, req.State)
	if err != nil {
		log.Println()
		c.Data["json"] = map[string]interface{}{consts.Stauts: consts.Failure, consts.Data: err}
		c.ServeJSON()
		return
	}

	c.Data["json"] = map[string]interface{}{consts.Stauts: consts.Success}
	c.ServeJSON()
}

// Get get a article by id.
func (c *ArticleController) Get() {
	var req struct {
		Id string `json:"id" validate:"required"`
	}
	// objectId := c.Ctx.Input.Param(":objectId")
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &req)
	if err != nil {
		log.Println(consts.ErrParam, err)
		c.Data["json"] = map[string]interface{}{consts.Stauts: consts.Failure, consts.Data: err}
		c.ServeJSON()
		return
	}

	// JWT validate.
	/*	tokenString := c.Ctx.Input.Header("Authorization")
		if ok, err := util.ValidateToken(tokenString); err != nil {
			log.Println(err)
			c.Data["json"] = map[string]interface{}{consts.Stauts: consts.Failure,consts.Data: util.ErrExpired}
			c.ServeJSON()
			return
		} else if !ok {
			c.Data["json"] = map[string]interface{}{consts.Stauts: consts.ErrLoginRequired}
			c.ServeJSON()
			return
		}*/

	a, err := models.ArticleService.Get(req.Id)
	if err != nil {
		if err == mgo.ErrNotFound {
			c.Data["json"] = map[string]interface{}{consts.Stauts: consts.Failure, consts.Data: err}
			c.ServeJSON()
		} else {
			log.Println("Mongodb error:", err)
			c.Data["json"] = err.Error()
		}
	}

	c.Data["json"] = map[string]interface{}{consts.Stauts: consts.Success, consts.Data: a}
	c.ServeJSON()
}

// All admin can use this model view all articles.
func (c *ArticleController) All() {
	var req struct {
		Page int `json:"page"`
	}

	err := json.Unmarshal(c.Ctx.Input.RequestBody, &req)
	if err != nil {
		log.Println(consts.ErrParam, err)
		c.Data["json"] = map[string]interface{}{consts.Stauts: consts.Failure, consts.Data: err}
		c.ServeJSON()
		return
	}

	articles, err := models.ArticleService.All(req.Page)
	if err != nil {
		log.Println(consts.ErrMongo, err)
		c.Data["json"] = map[string]interface{}{consts.Stauts: consts.Failure, consts.Data: err}
		c.ServeJSON()
		return
	}

	c.Data["json"] = map[string]interface{}{consts.Stauts: consts.Success, consts.Data: articles}
	c.ServeJSON()
}

// Approved view all approved articles.
func (c *ArticleController) Approved() {
	var req struct {
		Page int `json:"page"`
	}
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &req)
	if err != nil {
		log.Println(consts.ErrParam, err)
		c.Data["json"] = map[string]interface{}{consts.Stauts: consts.Failure, consts.Data: err}
		c.ServeJSON()
		return
	}

	articles, err := models.ArticleService.Approved(req.Page)
	if err != nil {
		if err == mgo.ErrNotFound {
			c.Data["json"] = map[string]interface{}{consts.Stauts: consts.Failure, consts.Data: err}
			c.ServeJSON()
			return
		}
		log.Println("Mongodb error:", err)
		return
	}

	c.Data["json"] = map[string]interface{}{consts.Stauts: consts.Success, consts.Data: articles}
	c.ServeJSON()
}

// ListCreated admin use this model to audit articles.
func (c *ArticleController) ListCreated() {
	var req struct {
		Page int `json:"page"`
	}

	err := json.Unmarshal(c.Ctx.Input.RequestBody, &req)
	if err != nil {
		log.Println(consts.ErrParam, err)
		c.Data["json"] = map[string]interface{}{consts.Stauts: consts.Failure, consts.Data: err}
		c.ServeJSON()
		return
	}

	articles, err := models.ArticleService.ListCreated(req.Page)
	if err != nil {
		if err == mgo.ErrNotFound {
			c.Data["json"] = map[string]interface{}{consts.Stauts: consts.Failure, consts.Data: err}
			c.ServeJSON()
			return
		}
		log.Println("Mongodb error:", err)
		return
	}

	c.Data["json"] = map[string]interface{}{consts.Stauts: consts.Success, consts.Data: articles}
	c.ServeJSON()
}
