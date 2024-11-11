package app

import (
	"context"
	"fmt"
	"github.com/EvgeniyBudaev/go-channel/internal/config"
	"github.com/EvgeniyBudaev/go-channel/internal/entity"
	"github.com/EvgeniyBudaev/go-channel/internal/logger"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"sync"
)

const (
	errorFilePathApp = "internal/app/app.go"
)

// App - application structure
type App struct {
	config *config.Config
	Logger logger.Logger
	fiber  *fiber.App
}

// New - create new application
func New() *App {
	// Default logger
	defaultLogger, err := logger.New(logger.GetDefaultLevel())
	if err != nil {
		errorMessage := getErrorMessage("New", "logger.New", errorFilePathApp)
		defaultLogger.Error(errorMessage, zap.Error(err))
	}

	// Config
	cfg, err := config.Load(defaultLogger)
	if err != nil {
		errorMessage := getErrorMessage("New", "config.Load", errorFilePathApp)
		defaultLogger.Error(errorMessage, zap.Error(err))
	}

	// Logger level
	loggerLevel, err := logger.New(cfg.LoggerLevel)
	if err != nil {
		errorMessage := getErrorMessage("New", "logger.New", errorFilePathApp)
		defaultLogger.Error(errorMessage, zap.Error(err))
	}

	// Fiber
	f := fiber.New(fiber.Config{
		ReadBufferSize: 256 << 8,
		BodyLimit:      50 * 1024 * 1024, // 50 MB
	})

	return &App{
		config: cfg,
		Logger: loggerLevel,
		fiber:  f,
	}
}

// Run launches the application
func (app *App) Run(ctx context.Context) {
	// Hub for telegram bot
	hub := entity.NewHub()

	msgChan := make(chan *entity.HubContent, 1) // msgChan - канал для передачи сообщений
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		if err := app.StartHTTPServer(ctx, hub); err != nil {
			errorMessage := getErrorMessage("Run", "StartHTTPServer",
				errorFilePathApp)
			app.Logger.Error(errorMessage, zap.Error(err))
		}
		wg.Done()
	}()
	wg.Add(1)
	go func() {
		if err := app.StartBot(ctx, msgChan); err != nil {
			errorMessage := getErrorMessage("Run", "StartBot", errorFilePathApp)
			app.Logger.Error(errorMessage, zap.Error(err))
		}
		wg.Done()
	}()
	wg.Add(1)
	go func() {
		for {
			select {
			case <-ctx.Done():
				wg.Done()
				return
			case c, ok := <-hub.Broadcast:
				if !ok {
					return
				}
				msgChan <- c
			}
		}
	}()
	wg.Wait()
}

func getErrorMessage(repositoryMethodName, callMethodName, errorFilePath string) string {
	return fmt.Sprintf("error func %s, method %s by path %s", repositoryMethodName, callMethodName,
		errorFilePath)
}
