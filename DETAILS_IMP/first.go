package main

import (
	"errors"
	"fmt"
)

type TestingHandler func(name string) error

type testing struct {
	name     *string
	handlers map[string]TestingHandler
}

func NewTesting() *testing {
	name := "Shyam"
	return &testing{
		name:     &name,
		handlers: map[string]TestingHandler{},
	}
}

func logger(next func(name string)) func(name string) {
	return func(name string) {
		fmt.Printf("[LOG] about to handle name: %s\n", name)
		next(name)
		fmt.Printf("[LOG] done handling name: %s\n", name)
	}
}

func (t *testing) handle(action string, handler TestingHandler) {
	t.handlers[action] = func(name string) error {
		var handlerErr error
		logger(func(name string) {
			handlerErr = handler(name)
		})(*t.name)
		if handlerErr != nil {
			fmt.Printf("handler error for action [%s]: %v\n", action, handlerErr)
		}
		return handlerErr
	}
}

// Step 5: convenience methods — thin wrappers over handle
// → same as: func (a *App) Get(...)  func (a *App) Post(...)
func (t *testing) Greet(handler TestingHandler)  { t.handle("greet", handler) }
func (t *testing) Shout(handler TestingHandler)  { t.handle("shout", handler) }
func (t *testing) Update(handler TestingHandler) { t.handle("update", handler) }

// Step 6: execute a registered action by name
func (t *testing) Run(action string) {
	handler, ok := t.handlers[action]
	if !ok {
		fmt.Printf("no handler registered for action: %s\n", action)
		return
	}
	handler(*t.name)
}

func main() {
	t := NewTesting()

	// Register handlers using the convenience methods
	t.Greet(func(name string) error {
		fmt.Printf("Hello, %s!\n", name)
		return nil
	})

	t.Shout(func(name string) error {
		fmt.Printf("HELLO, %s!!!\n", name)
		return nil
	})

	t.Update(func(name string) error {
		// simulate an error — same as a DB failure in a real handler
		return errors.New("update failed: name cannot be changed")
	})

	t.Run("greet")
	fmt.Println("---")
	t.Run("shout")
	fmt.Println("---")
	t.Run("update") // this one errors
}
