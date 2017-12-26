package controllers

import (
	"encoding/json"
	"revel-modules/xcontroller/app/web"

	"github.com/revel/revel"
)

// Controller definition for database transaction
// This controller is only useful if you intend to use the database instance
// defined in github.com/revel/modules/orm/gorp/app.Db
type XController struct {
	*revel.Controller
}
type Reponse struct {
	Data   interface{} `json:"data"`
	Status string      `json:"status"`
}
type ErrorResponse struct {
	Error  interface{} `json:"error"`
	Status string      `json:"status"`
}

func (x XController) MustDecodeBody(res interface{}) {
	var body = x.Request.GetBody()
	if err := json.NewDecoder(body).Decode(res); err != nil {
		web.AssertNil(err)
	}
}

func (x XController) SendData(data interface{}) revel.Result {
	x.Response.Status = 200
	return x.RenderJSON(Reponse{
		Data:   data,
		Status: "success",
	})
}

func (x XController) Success() revel.Result {
	x.Response.Status = 200
	return x.RenderJSON(Reponse{
		Data:   nil,
		Status: "success",
	})
}

func (x XController) SendError(err web.IWebError) revel.Result {
	x.Response.Status = err.StatusCode()
	return x.RenderJSON(ErrorResponse{
		Error:  err,
		Status: "error",
	})
}
