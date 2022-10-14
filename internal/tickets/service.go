package tickets

import (
        "context"
        "desafio-goweb-juanurbano/internal/domain"
)


type Service interface {
	GetAll(ctx context.Context) ([]domain.Ticket, error)
	GetTotalTickets(ctx context.Context, destination string) ([]domain.Ticket, error)
        AverageDestination(ctx context.Context, destination string) (float64, error)
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{repository: r}
}

func (s *service) GetAll(ctx context.Context) ([]domain.Ticket, error) {
	return s.repository.GetAll(ctx)
}

func (s *service) GetTotalTickets(ctx context.Context, destination string) ([]domain.Ticket, error) {
	return s.repository.GetTicketByDestination(ctx, destination) 
}

func (s *service) AverageDestination(ctx context.Context, destination string) (float64, error) {
	return s.repository.AverageDestination(ctx, destination) 
}
