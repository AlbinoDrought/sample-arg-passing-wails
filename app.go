package main

import (
	"context"
	"fmt"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx              context.Context
	instanceMessages <-chan string
}

// NewApp creates a new App application struct
func NewApp(instanceMessages <-chan string) *App {
	return &App{
		instanceMessages: instanceMessages,
	}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	go func() {
		for message := range a.instanceMessages {
			println("New instance message received:", message)
			runtime.EventsEmit(a.ctx, "instance-message", message)
		}
	}()
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}
