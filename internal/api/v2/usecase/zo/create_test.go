package zo_test

import (
	"context"
	"database/sql"
	"testing"
	"time"

	d "github.com/ArtefactGitHub/Go_P_Zo/internal/api/v2/domain/zo"
	i "github.com/ArtefactGitHub/Go_P_Zo/internal/api/v2/infrastructure"
	infra "github.com/ArtefactGitHub/Go_P_Zo/internal/api/v2/infrastructure/zo"
	. "github.com/ArtefactGitHub/Go_P_Zo/internal/api/v2/usecase/zo"
	"github.com/ArtefactGitHub/Go_P_Zo/internal/platform/mytime"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func Test_create_Do(t *testing.T) {
	var (
		zo = d.NewZo(0,
			mytime.ToTime("2023-01-01 01:00"),
			100,
			1,
			"hoge",
			time.Now(),
			sql.NullTime{},
			1)
	)
	type fields struct {
		r d.Repository
	}
	type args struct {
		zo     d.Zo
		userID int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    d.Zo
		wantErr bool
	}{
		{
			name:   "指定の1件が見つかった場合",
			fields: fields{r: infra.NewRepository()},
			args: args{
				zo:     zo,
				userID: 1,
			},
			want:    zo,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := NewCreate(tt.fields.r)
			c := context.WithValue(context.Background(), i.KeyDB, DB)
			c = context.WithValue(c, i.KeyTX, TX)

			got, err := u.Do(c, tt.args.zo)
			if (err != nil) != tt.wantErr {
				t.Errorf("Do() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if diff := cmp.Diff(got, tt.want, cmpopts.IgnoreFields(d.Zo{}, "CreatedAt", "UpdatedAt")); diff != "" {
				t.Errorf("Find() value is mismatch: %s", diff)
			}
		})
	}
}
