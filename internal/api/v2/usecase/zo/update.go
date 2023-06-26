package zo

import (
	"context"
	"fmt"

	derr "github.com/ArtefactGitHub/Go_P_Zo/internal/api/v2/domain/error"
	domain "github.com/ArtefactGitHub/Go_P_Zo/internal/api/v2/domain/zo"
	util "github.com/ArtefactGitHub/Go_P_Zo/internal/api/v2/presentation/utils"
)

type (
	Update interface {
		Do(context.Context, domain.Zo) (domain.Zo, error)
	}
	update struct {
		r domain.Repository
	}
)

func NewUpdate(r domain.Repository) Update {
	return update{r: r}
}

func (u update) Do(ctx context.Context, target domain.Zo) (domain.Zo, error) {
	z, err := u.r.Find(ctx, target.ID())
	if err != nil {
		return nil, err
	}
	if z.UserID() != target.UserID() {
		e := util.Wrap(derr.BadRequest, fmt.Sprintf("can not access by userID"))
		return nil, e
	}

	updateModel := toUpdateModel(z, target)
	err = u.r.Update(ctx, updateModel)
	if err != nil {
		return nil, err
	}

	return updateModel, nil
}

func toUpdateModel(z domain.Zo, target domain.Zo) domain.Zo {
	return domain.NewZo(
		z.ID(),
		target.AchievementDate(),
		target.Exp(),
		target.CategoryID(),
		target.Message(),
		z.CreatedAt(),
		target.UpdatedAt(),
		z.UserID(),
	)
}
