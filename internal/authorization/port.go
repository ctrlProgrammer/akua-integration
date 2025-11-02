package authorization

import "context"

type Provider interface {
	Authorize(ctx context.Context, authorization Authorization) (Authorization, error)
}

type Service struct {
	provider Provider
}

func NewService(provider Provider) *Service {
	return &Service{provider: provider}
}

func (s *Service) Authorize(ctx context.Context, authorization Authorization) (Authorization, error) {
	return s.provider.Authorize(ctx, authorization)
}
