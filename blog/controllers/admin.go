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
		Pwd  string `json:"pwd" validate:"required"`
	}

	err := json.Unmarshal(c.Ctx.Input.RequestBody, &req)
	if err != nil {
		log.Println(consts.ErrParam, err)
		c.Data["json"] = map[string]interface{}{"status:": err}
	} else {
		id, err := models.AdminService.New(req.Name, req.Pwd)
		if err != nil {
			log.Println(consts.ErrMongo, err)
			c.Data["json"] = map[string]interface{}{"status:": err}
		}
		c.Data["json"] = map[string]interface{}{"status": consts.Success, "id": id}
	}
	c.ServeJSON()
}

func (c *AdminController) Login() {
	//username := c.GetString("username")
	//password := c.GetString("password")
	var req struct {
		Name string `json:"name" validate:"required"`
		Pwd  string `json:"pwd" validate:"required"`
	}

	err := json.Unmarshal(c.Ctx.Input.RequestBody, &req)
	if err != nil {
		log.Println(consts.ErrParam, err)
		c.Data["json"] = map[string]interface{}{"status": err}
	} else if !models.AdminService.Login(req.Name, req.Pwd) {
		log.Println(consts.ErrMongo, err)
		c.Data["json"] = map[string]interface{}{"status": err}
	} else {
		c.Data["json"] = map[string]interface{}{"status": consts.Success}
	}
	c.ServeJSON()
}

func (u *AdminController) Logout() {
	u.Data["json"] = "logout success"
	u.ServeJSON()
}
