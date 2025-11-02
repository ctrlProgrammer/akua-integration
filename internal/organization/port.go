package organization

import "context"

// I will use the abstraction to maintain the domain and the implementation separated
type Provider interface {
	GetOrganizations(ctx context.Context) ([]Organization, error)
}

type Service struct {
	provider Provider
}

func NewService(provider Provider) *Service {
	return &Service{provider: provider}
}

func (s *Service) GetOrganizations(ctx context.Context) ([]Organization, error) {
	return s.provider.GetOrganizations(ctx)
}
