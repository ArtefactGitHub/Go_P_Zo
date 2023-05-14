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
	// TODO: zo.UserID検証

	z, err := u.r.Find(ctx, target.Id)
	if err != nil {
		return domain.Zo{}, err
	}
	if z.UserId != target.UserId {
		e := util.Wrap(derr.BadRequest, fmt.Sprintf("can not access by userID"))
		return domain.Zo{}, e
	}

	updateModel := toUpdateModel(z, target)
	err = u.r.Update(ctx, updateModel)
	if err != nil {
		return domain.Zo{}, err
	}

	return updateModel, nil
}

func toUpdateModel(z domain.Zo, target domain.Zo) domain.Zo {
	return domain.Zo{
		Id:              z.Id,
		AchievementDate: target.AchievementDate,
		Exp:             target.Exp,
		CategoryId:      target.CategoryId,
		Message:         target.Message,
		CreatedAt:       z.CreatedAt,
		UpdatedAt:       target.UpdatedAt,
		UserId:          z.UserId,
	}
}
