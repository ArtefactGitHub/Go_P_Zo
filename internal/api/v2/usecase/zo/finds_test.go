package zo_test

import (
	"context"
	"testing"

	domain "github.com/ArtefactGitHub/Go_P_Zo/internal/api/v2/domain/zo"
	i "github.com/ArtefactGitHub/Go_P_Zo/internal/api/v2/infrastructure"
	infra "github.com/ArtefactGitHub/Go_P_Zo/internal/api/v2/infrastructure/zo"
	. "github.com/ArtefactGitHub/Go_P_Zo/internal/api/v2/usecase/zo"
	"github.com/ArtefactGitHub/Go_P_Zo/internal/platform/mytime"
	"github.com/google/go-cmp/cmp"
)

func Test_finds_Do(t *testing.T) {
	type fields struct {
		r domain.Repository
	}
	type args struct {
		ctx context.Context
		id  int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []domain.Zo
		wantErr bool
	}{
		{
			name:   "指定のリソース群が見つかった場合",
			fields: fields{r: infra.NewRepository()},
			args: args{
				id: 1,
			},
			want: []domain.Zo{
				domain.NewZo(1, mytime.ToTime("2023-01-01 01:00"), 100, 1, "メッセージ1", mytime.ToTime("2023-01-01 01:00"), mytime.ToNullTime("2023-01-01 01:00"), 1),
				domain.NewZo(2, mytime.ToTime("2023-01-01 02:00"), 200, 1, "メッセージ2", mytime.ToTime("2023-01-01 02:00"), mytime.ToNullTime("2023-01-01 02:00"), 1),
				domain.NewZo(3, mytime.ToTime("2023-01-01 03:00"), 300, 1, "メッセージ3", mytime.ToTime("2023-01-01 03:00"), mytime.ToNullTime("2023-01-01 03:00"), 1),
			},
			wantErr: false,
		},
		{
			name:   "指定のリソース群が見つかった場合",
			fields: fields{r: infra.NewRepository()},
			args: args{
				id: -1,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := NewFinds(tt.fields.r)
			c := context.WithValue(context.Background(), i.KeyDB, DB)
			got, err := u.Do(c, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("Do() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("Find() value is mismatch: %s", diff)
			}
		})
	}
}
