package session

import (
	"context"
	"errors"
	"fmt"

	e "github.com/ArtefactGitHub/Go_P_Zo/internal/api/v2/domain/error"
	d "github.com/ArtefactGitHub/Go_P_Zo/internal/api/v2/domain/session"
)

type Login interface {
	Do(ctx context.Context, identifier string, secret string) (d.SessionData, error)
}
type login struct {
	r d.Repository
}

func NewLogin(r d.Repository) Login {
	return login{r: r}
}

func (s login) Do(ctx context.Context, identifier string, secret string) (d.SessionData, error) {
	result, err := s.r.Find(ctx, identifier, secret)
	if err != nil {
		if errors.Is(e.NotFound, err) {
			return d.SessionData{}, err
		}
		return d.SessionData{}, fmt.Errorf("%s: %w", err.Error(), e.BadRequest)
	}

	fmt.Printf("res: %#v \n", result)
	return result, nil
}
