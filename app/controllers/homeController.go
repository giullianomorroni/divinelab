package controllers

import (
	cb "divinelab/app/controllers/base"
	"github.com/robfig/revel"
	srv "divinelab/app/services"
)

type (
	Home struct {
		cb.BaseController
	}
)

func init() {}

func (this *Home) Index() revel.Result {
	return this.Render();
}
