package main

import (
	"context"
	"fmt"

	"github.com/hugolgst/rich-go/client"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx

}

func modify(id string, activity client.Activity) bool {
	err := client.Login(id)
	if err != nil {
		return false
	}

	err = client.SetActivity(activity)
	return err == nil
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}
