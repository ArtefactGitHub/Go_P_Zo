package zo

import (
	"context"
	"database/sql"
	"time"

	"github.com/ArtefactGitHub/Go_P_Zo/internal/platform/mydb"
)

type ZoService struct {
	Zr  ZoRepository
	Uzr UserZosRepository
}

func (s *ZoService) GetAll(ctx context.Context) ([]Zo, error) {
	result, err := s.Zr.Findall(ctx)
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
	// ZosとUserZosテーブルへレコードを追加するためトランザクション処理を行う
	result, err := mydb.Tran(ctx, func(ctx context.Context, tx *sql.Tx) (interface{}, error) {
		z.CreatedAt = time.Now()
		result, err := s.Zr.CreateTx(ctx, tx, z)
		if err != nil {
			return -1, err
		}

		// TODO：仮実装
		uz := UserZos{UserId: z.UserId, ZoId: z.Id, CreatedAt: time.Now()}
		_, err = s.Uzr.CreateTx(ctx, tx, &uz)
		if err != nil {
			return -1, err
		}

		return result, nil
	})

	if err != nil {
		return -1, err
	}

	return result.(int), err
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
