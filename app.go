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
	go runner()
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func runner() {

	err := client.Login("810540985032900648")
	if err != nil {
		panic(err)
	}

	err = client.SetActivity(client.Activity{
		State:      "Heyy!!!",
		Details:    "I'm running on rich-go :)",
		LargeImage: "largeimageid",
		LargeText:  "This is the large image :D",
		SmallImage: "smallimageid",
		SmallText:  "And this is the small image",
		Party: &client.Party{
			ID:         "-1",
			Players:    15,
			MaxPlayers: 24,
		},
	})

	if err != nil {
		panic(err)
	}
}
