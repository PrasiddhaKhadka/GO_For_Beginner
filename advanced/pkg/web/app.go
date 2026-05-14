package web

import "fmt"

type App struct {
}

type HandleFunc func(ctx *Ctx) (*Result, error)

func (a *App) AddResource(resourceName string, controller ResourceController) {
	a.handle("GET", resourceName, controller.Index)
	a.handle("GET", fmt.Sprintf("%s/create", resourceName), controller.Create)
	a.handle("POST", resourceName, controller.Store)
	a.handle("GET", fmt.Sprintf("%s/:id", resourceName), controller.Show)
	a.handle("GET", fmt.Sprintf("%s/edit", resourceName), controller.Edit)
	a.handle("PUT", fmt.Sprintf("%s/:id/edit", resourceName), controller.Update)
	a.handle("DELETE", fmt.Sprintf("%s/:id", resourceName), controller.Delete)
}

func (a *App) handle(method string, path string, handler HandleFunc) {}
