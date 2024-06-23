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

var todos = []string{}

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

func (a *App) GetTodos() []string {
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
	todos = append(todos, todo)
}

func (a *App) RemoveTodo(todo string) {
	for i, t := range todos {
		if t == todo {
			todos = append(todos[:i], todos[i+1:]...)
			break
		}
	}

}
