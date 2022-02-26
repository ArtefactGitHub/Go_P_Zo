package zo

import (
	"context"
	"database/sql"
	"time"
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

func (s *ZoService) Post(ctx context.Context, userId int, rz *requestZo) (int, error) {
	if zo, err := s.newZo(rz, userId); err != nil {
		return -1, err
	} else {
		return s.Zr.Create(ctx, zo)
	}
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

func (s *ZoService) newZo(rz *requestZo, userId int) (*Zo, error) {
	if err := rz.validation(); err != nil {
		return nil, err
	}
	result := NewZo(0, rz.AchievementDate, rz.Exp, rz.CategoryId, rz.Message, time.Now(), sql.NullTime{}, userId)
	return &result, nil
}
