package app

import (
	"context"
	"github.com/EvgeniyBudaev/go-channel/internal/controller"
	"github.com/EvgeniyBudaev/go-channel/internal/entity"
	"go.uber.org/zap"
)

const (
	errorFilePathHttp = "internal/app/http.go"
	prefix            = "/gateway/api/v1"
)

func (app *App) StartHTTPServer(ctx context.Context, hub *entity.Hub) error {
	app.fiber.Static("/static", "./static")
	router := app.fiber.Group(prefix)
	done := make(chan struct{})
	profileController := controller.NewProfileController(app.Logger, hub)
	router.Post("/profiles/likes", profileController.AddLike())
	go func() {
		port := ":" + app.config.Port
		if err := app.fiber.Listen(port); err != nil {
			errorMessage := getErrorMessage("StartHTTPServer", "Listen",
				errorFilePathHttp)
			app.Logger.Error(errorMessage, zap.Error(err))
		}
		close(done)
	}()
	select {
	case <-ctx.Done():
		if err := app.fiber.Shutdown(); err != nil {
			errorMessage := getErrorMessage("StartHTTPServer", "Shutdown",
				errorFilePathHttp)
			app.Logger.Error(errorMessage, zap.Error(err))
		}
	case <-done:
		app.Logger.Info("server finished successfully")
	}
	return nil
}
