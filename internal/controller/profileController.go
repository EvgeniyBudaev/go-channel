package controller

import (
	v1 "github.com/EvgeniyBudaev/go-channel/internal/controller/http/api/v1"
	"github.com/EvgeniyBudaev/go-channel/internal/entity"
	"github.com/EvgeniyBudaev/go-channel/internal/logger"
	"github.com/gofiber/fiber/v2"
)

type ProfileController struct {
	logger logger.Logger
	hub    *entity.Hub
}

func NewProfileController(l logger.Logger, h *entity.Hub) *ProfileController {
	return &ProfileController{
		logger: l,
		hub:    h,
	}
}

func (pc *ProfileController) AddLike() fiber.Handler {
	return func(ctf *fiber.Ctx) error {
		pc.logger.Info("POST /gateway/api/v1/profiles/likes")
		go func() {
			pc.hub.Broadcast <- &entity.HubContent{
				Type:         "like",
				Message:      "Hello, World!",
				UserId:       1,
				UserImageUrl: "https://my_image.jpeg",
				Username:     "my_name",
			}
		}()
		return v1.ResponseCreated(ctf, "OK")
	}
}
