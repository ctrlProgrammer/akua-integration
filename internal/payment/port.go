package payment

import "context"

// I will use the abstraction to maintain the domain and the implementation separated
type Provider interface {
	GetPayments(ctx context.Context) ([]Payment, error)
}

type Service struct {
	provider Provider
}

func NewService(provider Provider) *Service {
	return &Service{provider: provider}
}

func (s *Service) GetPayments(ctx context.Context) ([]Payment, error) {
	return s.provider.GetPayments(ctx)
}
