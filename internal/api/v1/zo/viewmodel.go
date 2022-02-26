package zo

import (
	"fmt"
	"time"

	"github.com/ArtefactGitHub/Go_P_Zo/internal/platform/myhttp"
)

type requestZo struct {
	AchievementDate time.Time `json:"achievementdate"`
	Exp             int       `json:"exp"`
	CategoryId      int       `json:"category_id"`
	Message         string    `json:"message"`
}

func NewRequestZo(
	achievementDate time.Time,
	exp int,
	categoryId int,
	message string,
) *requestZo {
	return &requestZo{
		AchievementDate: achievementDate,
		Exp:             exp,
		CategoryId:      categoryId,
		Message:         message,
	}
}

func (m *requestZo) validation() error {
	if m.AchievementDate.Unix() > time.Now().Unix() {
		return fmt.Errorf("達成日に未来が設定されています。achievementDate: %s", m.AchievementDate)
	}
	if m.Exp < 0 || m.Exp > 1000 {
		return fmt.Errorf("獲得経験値は0〜1000の範囲で指定してください。exp: %d", m.Exp)
	}
	if len(m.Message) > 30 {
		return fmt.Errorf("メッセージは30文字以内で指定してください。")
	}
	return nil
}

type responseZo struct {
	Id              int       `json:"id"`
	AchievementDate time.Time `json:"achievementdate"`
	Exp             int       `json:"exp"`
	CategoryId      int       `json:"category_id"`
	Message         string    `json:"message"`
}

func NewResponseZo(
	id int,
	achievementDate time.Time,
	exp int,
	categoryId int,
	message string,
) *responseZo {
	return &responseZo{
		Id:              id,
		AchievementDate: achievementDate,
		Exp:             exp,
		CategoryId:      categoryId,
		Message:         message,
	}
}

type GetAllResponse struct {
	myhttp.ResponseBase
	Zos []Zo `json:"zos"`
}

type GetResponse struct {
	myhttp.ResponseBase
	Zo *Zo `json:"zo"`
}

type PostResponse struct {
	myhttp.ResponseBase
	Zo *responseZo `json:"zo"`
}

type PutResponse struct {
	myhttp.ResponseBase
	Zo *Zo `json:"zo"`
}

type DeleteResponse struct {
	myhttp.ResponseBase
}
