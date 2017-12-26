package controllers

import (
	"encoding/json"
	"fmt"
	"revel-modules/xcontroller/app/utils"
	"revel-modules/xcontroller/app/web"

	"github.com/revel/revel"
)

// Controller definition for database transaction
// This controller is only useful if you intend to use the database instance
// defined in github.com/revel/modules/orm/gorp/app.Db
type XController struct {
	*revel.Controller
}

func (x XController) Validate() revel.Result {
	fmt.Println("dasd")
	if x.Validation.HasErrors() {
		return x.SendCustomError(x.Validation.Errors[0].Message)
	}
	return nil
}

func (x XController) MustDecodeBody(res interface{}) {
	var body = x.Request.GetBody()
	if err := json.NewDecoder(body).Decode(res); err != nil {
		web.AssertNil(err)
	}
}

func (x XController) SendData(data interface{}) revel.Result {
	x.Response.Status = 200
	return x.RenderJSON(utils.Reponse{
		Data:   data,
		Status: "success",
	})
}

func (x XController) Success() revel.Result {
	x.Response.Status = 200
	return x.RenderJSON(utils.Reponse{
		Data:   nil,
		Status: "success",
	})
}
func (x XController) SendCustomError(err string, code ...int) revel.Result {
	if len(code) > 0 {
		x.Response.Status = code[0]
	}
	x.Response.Status = 400
	return x.RenderJSON(utils.ErrorResponse{
		Error:  err,
		Status: "error",
	})
}
func (x XController) SendError(err web.IWebError) revel.Result {
	x.Response.Status = err.StatusCode()
	return x.RenderJSON(utils.ErrorResponse{
		Error:  err,
		Status: "error",
	})
}
