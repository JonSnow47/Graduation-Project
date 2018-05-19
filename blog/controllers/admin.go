package controllers

import (
	"encoding/json"
	"log"

	"github.com/astaxie/beego"

	"github.com/JonSnow47/Graduation-Project/blog/consts"
	"github.com/JonSnow47/Graduation-Project/blog/models"
)

type AdminController struct {
	beego.Controller
}

func (c *AdminController) New() {
	var req struct {
		Name string `json:"name" validate:"required"`
		Pwd  string `json:"password" validate:"required"`
	}

	err := json.Unmarshal(c.Ctx.Input.RequestBody, &req)
	if err != nil {
		log.Println(consts.ErrParam, err)
		c.Data["json"] = map[string]interface{}{consts.Status: consts.ErrParam}
	} else {
		id, err := models.AdminService.New(req.Name, req.Pwd)
		if err != nil {
			log.Println(consts.ErrMongo, err)
			c.Data["json"] = map[string]interface{}{consts.Status: err}
		}
		c.Data["json"] = map[string]interface{}{consts.Status: consts.Success, consts.Data: map[string]string{"id": id}}
	}
	c.ServeJSON()
}

func (c *AdminController) Login() {
	var req struct {
		Name string `json:"name" validate:"required"`
		Pwd  string `json:"password" validate:"required"`
	}

	err := json.Unmarshal(c.Ctx.Input.RequestBody, &req)
	if err != nil {
		log.Println(consts.ErrParam, err)
		c.Data["json"] = map[string]interface{}{consts.Status: consts.ErrParam}
	} else {
		if id := models.AdminService.Login(req.Name, req.Pwd); id == "" {
			log.Println(consts.ErrLogin, err)
			c.Data["json"] = map[string]interface{}{consts.Status: consts.ErrLogin}
		} else {
			// token, err := util.NewToken(id)
			c.SetSession(consts.SessionId, id)
			if err != nil {
				log.Println("Session error:", err)
			} else {
				c.Data["json"] = map[string]interface{}{consts.Status: consts.Success}
			}
		}
	}
	c.ServeJSON()
}

func (c *AdminController) Logout() {
	c.DelSession(consts.SessionId)
	c.Data["json"] = "logout success"
	c.ServeJSON()
}
