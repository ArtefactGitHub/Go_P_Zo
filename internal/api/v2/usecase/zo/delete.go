package zo

import (
	"context"
	"fmt"

	derr "github.com/ArtefactGitHub/Go_P_Zo/internal/api/v2/domain/error"
	domain "github.com/ArtefactGitHub/Go_P_Zo/internal/api/v2/domain/zo"
	util "github.com/ArtefactGitHub/Go_P_Zo/internal/api/v2/presentation/utils"
)

type (
	Delete interface {
		Do(context.Context, int) error
	}
	deleteUsecase struct {
		r domain.Repository
	}
)

func NewDelete(r domain.Repository) Delete {
	return deleteUsecase{r: r}
}

func (u deleteUsecase) Do(ctx context.Context, id int) error {
	// TODO: zo.UserID検証

	z, err := u.r.Find(ctx, id)
	if err != nil {
		return err
	}
	if z.UserId != id {
		e := util.Wrap(derr.BadRequest, fmt.Sprintf("can not access by userID"))
		return e
	}

	err = u.r.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
