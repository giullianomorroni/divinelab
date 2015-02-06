package controllers

import (
	cb "divinelab/app/controllers/base"
	"github.com/robfig/revel"
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

func (this *Home) Valores() revel.Result {
	return this.Render();
}

func (this *Home) Produtos() revel.Result {
	return this.Render();
}