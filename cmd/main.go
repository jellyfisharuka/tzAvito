package main

import (
	"context"
	"log"
	"tzAvito/internal/app"
)

func main() { 
	ctx := context.Background()

	a, err := app.NewApp(ctx)
	if err != nil {
		log.Fatalf("Error creating app: %v", err)
	}

	if err := a.Run(); err != nil {
		log.Fatalf("Error running app: %v", err)
	}
}
