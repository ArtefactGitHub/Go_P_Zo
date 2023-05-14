package zo_test

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/ArtefactGitHub/Go_P_Zo/internal/api/v2/domain/zo"
	i "github.com/ArtefactGitHub/Go_P_Zo/internal/api/v2/infrastructure"
	inZo "github.com/ArtefactGitHub/Go_P_Zo/internal/api/v2/infrastructure/zo"
	. "github.com/ArtefactGitHub/Go_P_Zo/internal/api/v2/usecase/zo"
	"github.com/ArtefactGitHub/Go_P_Zo/internal/platform/mytime"
	"github.com/google/go-cmp/cmp"
)

func Test_update_Do(t *testing.T) {
	type fields struct {
		r zo.Repository
	}
	type args struct {
		target zo.Zo
	}
	var (
		want = zo.NewZo(
			1,
			mytime.ToTime("2023-01-01 01:00"),
			100,
			1,
			"メッセージ1-updated",
			mytime.ToTime("2023-01-01 01:00"),
			mytime.ToNullTime("2023-01-01 01:00"),
			1)
		notFound      = zo.NewZo(999, time.Time{}, 0, 0, "", time.Time{}, sql.NullTime{}, 0)
		invalidUserID = zo.NewZo(
			1,
			mytime.ToTime("2023-01-01 01:00"),
			100,
			1,
			"メッセージ1-updated",
			mytime.ToTime("2023-01-01 01:00"),
			mytime.ToNullTime("2023-01-01 01:00"),
			999)
	)
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    zo.Zo
		wantErr bool
	}{
		{
			name:    "正常系：正しく更新が行える",
			fields:  fields{r: inZo.NewRepository()},
			args:    args{target: want},
			want:    want,
			wantErr: false,
		},
		{
			name:    "異常系：存在しないIDを指定する",
			fields:  fields{r: inZo.NewRepository()},
			args:    args{target: notFound},
			wantErr: true,
		},
		{
			name:    "異常系：存在しないIDを指定する",
			fields:  fields{r: inZo.NewRepository()},
			args:    args{target: invalidUserID},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := NewUpdate(tt.fields.r)
			c := context.WithValue(context.Background(), i.KeyDB, DB)
			c = context.WithValue(c, i.KeyTX, TX)
			got, err := u.Do(c, tt.args.target)
			if (err != nil) != tt.wantErr {
				t.Errorf("Do() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("Do() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
