package handlers

import "rest-api-redis/pkg/repository"

type Handler struct {
	UserHandler *UserHandler
}

func InitHandler(repository *repository.Repository) *Handler {
	return &Handler{
		UserHandler: InitUserHandler(repository.UserRepository),
	}
}
