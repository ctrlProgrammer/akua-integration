package commerce

import "context"

// I will use the abstraction to maintain the domain and the implementation separated
type Provider interface {
	GetOrganizationCommerces(ctx context.Context) ([]Commerce, error)
	CreateCommerce(ctx context.Context, commerce Commerce) (Commerce, error)
	GetCommerceById(ctx context.Context, id string) (Commerce, error)
	DeleteCommerce(ctx context.Context, id string) error
	UpdateCommerce(ctx context.Context, id string, commerce Commerce) (Commerce, error)
}

type Service struct {
	provider Provider
}

func NewService(provider Provider) *Service {
	return &Service{provider: provider}
}

func (s *Service) GetOrganizationCommerces(ctx context.Context) ([]Commerce, error) {
	return s.provider.GetOrganizationCommerces(ctx)
}

func (s *Service) CreateCommerce(ctx context.Context, commerce Commerce) (Commerce, error) {
	return s.provider.CreateCommerce(ctx, commerce)
}

func (s *Service) GetCommerceById(ctx context.Context, id string) (Commerce, error) {
	return s.provider.GetCommerceById(ctx, id)
}

func (s *Service) DeleteCommerce(ctx context.Context, id string) error {
	return s.provider.DeleteCommerce(ctx, id)
}

func (s *Service) UpdateCommerce(ctx context.Context, id string, commerce Commerce) (Commerce, error) {
	return s.provider.UpdateCommerce(ctx, id, commerce)
}
