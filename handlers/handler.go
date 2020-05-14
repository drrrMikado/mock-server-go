package handlers

import "github.com/drrrMikado/mock-server-go/services"

var (
	s *services.Service
)

func Init(svc *services.Service) {
	s = svc
}
