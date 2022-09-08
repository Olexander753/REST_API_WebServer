package user

import "context"

type Service struct {
	storage Storage
}

func (s *Service) Create(ctx context.Context, dto CreateUser) (User, error) {
	//TODO later
	panic("qwe")
}
