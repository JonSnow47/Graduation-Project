package controllers

import (
	"encoding/json"
	"log"

	"github.com/astaxie/beego"

	"github.com/JonSnow47/Graduation-Project/blog/consts"
	"github.com/JonSnow47/Graduation-Project/blog/models"
)

type TagController struct{ beego.Controller }

func (c *TagController) New() {
	var (
		req struct {
			Tag string `json:"tag" validate:"required,len<16"`
		}
		err error
	)

	err = json.Unmarshal(c.Ctx.Input.RequestBody, &req)
	if err != nil {
		log.Println(consts.ErrParam, err)
		c.Data["json"] = map[string]interface{}{consts.Status: consts.ErrParam, consts.Data: err}
		goto Finish
	}

	err = models.TagService.New(req.Tag)
	if err != nil {
		log.Println(consts.ErrMongo, err)
		c.Data["json"] = map[string]interface{}{consts.Status: consts.ErrMongo, consts.Data: err}
		goto Finish
	}

	c.Data["json"] = map[string]interface{}{consts.Status: consts.Success}
Finish:
	c.ServeJSON()
}

func (c *TagController) Delete() {
	var req struct {
		Tag string `json:"tag" validate:"len=24"`
	}

	err := json.Unmarshal(c.Ctx.Input.RequestBody, &req)
	if err != nil {
		log.Println(consts.ErrParam, err)
		c.Data["json"] = map[string]interface{}{consts.Status: consts.ErrParam, consts.Data: err}
		goto Finish
	}

	if err = models.TagService.Delete(req.Tag); err != nil {
		log.Println(consts.ErrMongo, err)
		c.Data["json"] = map[string]interface{}{consts.Status: consts.ErrMongo, consts.Data: err}
		goto Finish
	}

	c.Data["json"] = map[string]interface{}{consts.Status: consts.Success}
Finish:
	c.ServeJSON()
}

func (c *TagController) Enable() {
	var req struct {
		Tag string `json:"tag"`
	}

	err := json.Unmarshal(c.Ctx.Input.RequestBody, &req)
	if err != nil {
		log.Println(consts.ErrParam, err)
		c.Data["json"] = map[string]interface{}{consts.Status: consts.ErrParam, consts.Data: err}
		goto Finish
	}

	if err = models.TagService.Enable(req.Tag); err != nil {
		log.Println(consts.ErrMongo, err)
		c.Data["json"] = map[string]interface{}{consts.Status: consts.ErrMongo, consts.Data: err}
		goto Finish
	}

	c.Data["json"] = map[string]interface{}{consts.Status: consts.Success}
Finish:
	c.ServeJSON()
}

// Get a Tag info.
func (c *TagController) Get() {
	var (
		req struct {
			Tag string `json:"tag"`
		}
		t *models.Tag
	)
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &req)
	if err != nil {
		log.Println(consts.ErrParam, err)
		c.Data["json"] = map[string]interface{}{consts.Status: consts.ErrParam, consts.Data: err}
		goto Finish
	}

	t, err = models.TagService.Get(req.Tag)
	if err != nil {
		log.Println(consts.ErrMongo, err)
		c.Data["json"] = map[string]interface{}{consts.Status: consts.ErrMongo, consts.Data: err}
		goto Finish
	}

	c.Data["json"] = map[string]interface{}{consts.Status: consts.Success, consts.Data: t}
Finish:
	c.ServeJSON()
}

// All list of all tags.
func (c *TagController) All() {
	var tags []*models.Tag

	tags, err := models.TagService.All()
	if err != nil {
		log.Println(consts.ErrMongo, err)
		c.Data["json"] = map[string]interface{}{consts.Status: consts.ErrMongo, consts.Data: err}
		goto Finish
	}

	c.Data["json"] = map[string]interface{}{consts.Status: consts.Success, consts.Data: tags}
Finish:
	c.ServeJSON()
}
