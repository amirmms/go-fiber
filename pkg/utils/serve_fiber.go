package utils

import (
	"fiber_test/configs"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
	"os"
	"os/signal"
)

func ServeFiberInProductionMode(app *fiber.App, conf configs.Config) {
	// Create channel for idle connections.
	idleConnsClosed := make(chan struct{})

	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt) // Catch OS signals.
		<-sigint

		// Received an interrupt signal, shutdown.
		if err := app.Shutdown(); err != nil {
			// Error from closing listeners, or context timeout:
			log.Printf("Oops... Server is not shutting down! Reason: %v", err)
		}

		close(idleConnsClosed)
	}()

	// Build Fiber connection URL.
	fiberConnURL := fmt.Sprintf(
		"%s:%s",
		conf.AppHost,
		conf.AppPort,
	)

	// Run server.
	if err := app.Listen(fiberConnURL); err != nil {
		log.Printf("Oops... Server is not running! Reason: %v", err)
	}

	<-idleConnsClosed
}

func ServeFiberInDevelopMode(app *fiber.App, conf configs.Config) {
	// Build Fiber connection URL.
	fiberConnURL := fmt.Sprintf(
		"%s:%s",
		conf.AppHost,
		conf.AppPort,
	)

	// Run server.
	if err := app.Listen(fiberConnURL); err != nil {
		log.Printf("Oops... Server is not running! Reason: %v", err)
	}
}
