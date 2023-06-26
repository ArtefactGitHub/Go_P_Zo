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
	"github.com/google/go-cmp/cmp/cmpopts"
)

func Test_find_Do(t *testing.T) {
	type fields struct {
		r domain.Repository
	}
	type args struct {
		id int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    domain.Zo
		wantErr bool
	}{
		{
			name:   "指定の1件が見つかった場合",
			fields: fields{r: infra.NewRepository()},
			args: args{
				id: 1,
			},
			want:    domain.NewZo(1, mytime.ToTime("2023-01-01 01:00"), 100, 1, "メッセージ1", mytime.ToTime("2023-01-01 01:00"), mytime.ToNullTime("2023-01-01 01:00"), 1),
			wantErr: false,
		},
		{
			name:   "指定の1件が見つからない場合",
			fields: fields{r: infra.NewRepository()},
			args: args{
				id: -1,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := NewFind(tt.fields.r)
			c := context.WithValue(context.Background(), i.KeyDB, DB)
			got, err := u.Do(c, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("Do() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil {
				return
			}

			opts := cmp.Options{
				cmp.AllowUnexported(tt.want),
				cmpopts.IgnoreFields(tt.want, "createdAt", "updatedAt"),
			}
			if diff := cmp.Diff(got, tt.want, opts); diff != "" {
				t.Errorf("Find() value is mismatch: %s", diff)
			}
		})
	}
}
