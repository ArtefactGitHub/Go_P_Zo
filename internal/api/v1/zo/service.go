package zo

import (
	"context"
)

type ZoService struct {
	Zr ZoRepository
}

func (s *ZoService) GetAll(ctx context.Context, userId int) ([]Zo, error) {
	result, err := s.Zr.FindAllByUserId(ctx, userId)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *ZoService) Get(ctx context.Context, id int) (*Zo, error) {
	result, err := s.Zr.Find(ctx, id)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *ZoService) Post(ctx context.Context, z *Zo) (int, error) {
	result, err := s.Zr.Create(ctx, z)
	if err != nil {
		return -1, err
	}
	return result, nil
}

func (s *ZoService) Update(ctx context.Context, z *Zo) error {
	err := s.Zr.Update(ctx, z)
	if err != nil {
		return err
	}

	return nil
}

func (s *ZoService) Delete(ctx context.Context, id int) error {
	err := s.Zr.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
