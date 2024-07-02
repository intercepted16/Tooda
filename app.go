package main

import (
	"context"
	"fmt"

	"golang.org/x/sys/windows/registry"
)

// App struct
type App struct {
	ctx context.Context
}

type Todos map[string]bool

var todos = Todos{}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *App) Test() string {
	return "Hello World"
}

func (a *App) GetTodos() Todos {
	return todos
}

func (a *App) GetWindowsTheme() (string, error) {
	key, err := registry.OpenKey(registry.CURRENT_USER, `Software\Microsoft\Windows\CurrentVersion\Themes\Personalize`, registry.QUERY_VALUE)
	if err != nil {
		return "", err
	}
	defer key.Close()

	lightTheme, _, err := key.GetIntegerValue("AppsUseLightTheme")
	if err != nil {
		return "", err
	}

	if lightTheme == 0 {
		return "dark", nil
	}
	return "light", nil
}

func (a *App) AddTodo(todo string) {
	todos[todo] = false
}

func (a *App) RemoveTodo(todo string) {
	delete(todos, todo)
}

func (a *App) CheckTodo(todo string) {
	todos[todo] = !todos[todo]
}
