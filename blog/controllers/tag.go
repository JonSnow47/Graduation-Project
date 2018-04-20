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
			Tag string `json:"tag"`
		}
		id  string
		err error
	)

	err = json.Unmarshal(c.Ctx.Input.RequestBody, &req)
	if err != nil {
		log.Println(consts.ErrParam, err)
		c.Data["json"] = map[string]interface{}{"status:": err}
		goto Finish
	}

	id, err = models.TagService.New(req.Tag)
	if err != nil {
		log.Println(consts.ErrMongo, err)
		c.Data["json"] = map[string]interface{}{"status:": err}
		goto Finish
	}

	c.Data["json"] = map[string]interface{}{"status": consts.Success, "data": id}
Finish:
	c.ServeJSON()
}

func (c *TagController) Delete() {
	var req struct {
		Id string `json:"id"`
	}

	err := json.Unmarshal(c.Ctx.Input.RequestBody, &req)
	if err != nil {
		log.Println(consts.ErrParam, err)
		c.Data["json"] = map[string]interface{}{"status": err}
		goto Finish
	}

	if err = models.TagService.Delete(req.Id); err != nil {
		log.Println(consts.ErrMongo, err)
		c.Data["json"] = map[string]interface{}{"status": err}
		goto Finish
	}

	c.Data["json"] = map[string]interface{}{"status": consts.Success}
Finish:
	c.ServeJSON()
}

func (c *TagController) Enable() {
	var req struct {
		Id string `json:"id"`
	}

	err := json.Unmarshal(c.Ctx.Input.RequestBody, &req)
	if err != nil {
		log.Println(consts.ErrParam, err)
		c.Data["json"] = map[string]interface{}{"status": err}
		goto Finish
	}

	if err = models.TagService.Enable(req.Id); err != nil {
		log.Println(consts.ErrMongo, err)
		c.Data["json"] = map[string]interface{}{"status": err}
		goto Finish
	}

	c.Data["json"] = map[string]interface{}{"status": consts.Success}
Finish:
	c.ServeJSON()
}

// Get a Tag info.
func (c *TagController) Get() {
	var (
		req struct {
			Id string `json:"id"`
		}
		t *models.Tag
	)
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &req)
	if err != nil {
		log.Println(consts.ErrParam, err)
		c.Data["json"] = map[string]interface{}{"status": err}
		goto Finish
	}

	t, err = models.TagService.Get(req.Id)
	if err != nil {
		log.Println(consts.ErrMongo, err)
		c.Data["json"] = map[string]interface{}{"status": err}
		goto Finish
	}

	c.Data["json"] = map[string]interface{}{"status": consts.Success, "data": t}
Finish:
	c.ServeJSON()
}

// All list of all tags.
func (c *TagController) All() {
	var (
		req struct {
			Page int `json:"page"`
		}
		t []*models.Tag
	)

	err := json.Unmarshal(c.Ctx.Input.RequestBody, &req)
	if err != nil {
		log.Println(consts.ErrParam, err)
		c.Data["json"] = map[string]interface{}{"status": err}
		goto Finish
	}

	t, err = models.TagService.All(req.Page)
	if err != nil {
		log.Println(consts.ErrMongo, err)
		c.Data["json"] = map[string]interface{}{"status": err}
		goto Finish
	}

	c.Data["json"] = map[string]interface{}{"status": consts.Success, "data": t}
Finish:
	c.ServeJSON()
}
