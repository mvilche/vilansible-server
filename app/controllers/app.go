package controllers

import (
	"fmt"
	"vilansible/app/models"

	"github.com/revel/revel"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	return c.Render()
}

func (c App) List() revel.Result {

	t := models.ListExecution()
	return c.Render(t)
}

func (c App) ApiList() revel.Result {

	t := models.ListExecution()
	return c.RenderJSON(t)
}

func (c App) Create() revel.Result {

	var e models.Execution
	var jsonData map[string]interface{}
	c.Params.BindJSON(&jsonData)

	e.Application = fmt.Sprintf("%v", jsonData["application"])
	e.Version = fmt.Sprintf("%v", jsonData["version"])
	e.Date = fmt.Sprintf("%v", jsonData["date"])
	s := fmt.Sprintf("%v", jsonData["status"])
	e.TypeExec = fmt.Sprintf("%v", jsonData["typexec"])
	e.User = fmt.Sprintf("%v", jsonData["user"])

	if e.Application == "<nil>" || e.Date == "<nil>" || s == "" || e.TypeExec == "<nil>" || e.Version == "<nil>" || e.User == "<nil>" {
		c.Response.Status = 403
		return c.RenderJSON("Missing parameters")
	}

	e.Status = revel.ToBool(s)

	if err := models.AddExecution(e); err != nil {
		c.Response.Status = 500
		return c.RenderJSON(err.Error())
	} else {
		return c.RenderJSON("Success")
	}

}
