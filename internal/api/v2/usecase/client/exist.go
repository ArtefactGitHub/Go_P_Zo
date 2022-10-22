package client

import (
	"context"

	d "github.com/ArtefactGitHub/Go_P_Zo/internal/api/v2/domain/client"
	e "github.com/ArtefactGitHub/Go_P_Zo/internal/api/v2/domain/error"
)

type Exist interface {
	Do(ctx context.Context, id int, secret string) (bool, error)
}

type exist struct {
	r d.Repository
}

func NewExist(r d.Repository) Exist {
	return exist{r: r}
}

func (u exist) Do(ctx context.Context, id int, secret string) (bool, error) {
	_, err := u.r.Find(ctx, id, secret)
	if err != nil && err != e.NotFound {
		return false, err
	}

	return err != e.NotFound, nil
}
