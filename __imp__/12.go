package main

import (
	"errors"
	"fmt"
)

type NameHandler func(name string) error

type App struct {
	handlers map[string]NameHandler
}

func NewApp() *App {
	return &App{
		handlers: make(map[string]NameHandler),
	}
}

func (a *App) handle(action string, handler NameHandler) {
	a.handlers[action] = handler
}

// convenience methods — thin wrappers over handle
// exactly like app.Get, app.Post, app.Put
func (a *App) Greet(handler NameHandler) {
	a.handle("greet", handler)
}

func (a *App) Shout(handler NameHandler) {
	a.handle("shout", handler)

}

func (a *App) Update(handler NameHandler) { a.handle("update", handler) }

// method on App — runs a registered handler by action name

func (a *App) Run(action string, name string) {
	handler, ok := a.handlers[action]
	if !ok {
		fmt.Println("no handler for action:", action)
		return
	}

	if err := handler(name); err != nil {
		fmt.Println("error:", err)
	}
}

func main() {

	app := NewApp()

	app.Greet(func(name string) error {
		fmt.Println("Hello", name)
		return nil
	})

	app.Shout(func(name string) error {
		fmt.Println("HELLO", name)
		return nil
	})

	app.Update(func(name string) error {
		return errors.New("update failed")
	})

	fmt.Println(app.handlers)
	fmt.Println(app.handlers["greet"]("Ram"))

	app.Run("greet", "Ram")
}
