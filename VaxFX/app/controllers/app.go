package controllers

import (
	"github.com/revel/revel"
)

type App struct {
	*revel.Controller
}

func (c App) Submit() revel.Result {
	return c.Render()
}

func (c App) DataSecFX() revel.Result {
	return c.Render()
}

func (c App) DataCovid() revel.Result {
	return c.Render()
}