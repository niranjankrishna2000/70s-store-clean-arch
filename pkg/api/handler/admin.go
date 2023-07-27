package handler

import services "main/pkg/usecase/interface"

type AdminHandler struct {
	adminUseCase services.AdminUseCase
}

