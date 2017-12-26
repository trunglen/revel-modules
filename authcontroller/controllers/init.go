package controllers

import (
	"github.com/revel/revel"
)

func authenticate(c *revel.Controller) revel.Result {
	return nil
}
func init() {
	revel.InterceptFunc(authenticate, revel.BEFORE, &AuthController{})
}
