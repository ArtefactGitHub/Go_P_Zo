package auth_test

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/ArtefactGitHub/Go_P_Zo/internal/api/v2/domain/auth"
	infra "github.com/ArtefactGitHub/Go_P_Zo/internal/api/v2/infrastructure"
	. "github.com/ArtefactGitHub/Go_P_Zo/internal/api/v2/infrastructure/auth"
	"github.com/ArtefactGitHub/Go_P_Zo/internal/platform/mycontext"
	"github.com/ArtefactGitHub/Go_P_Zo/internal/platform/mydb"
	"github.com/ArtefactGitHub/Go_P_Zo/internal/platform/mytime"
	"github.com/google/go-cmp/cmp"
)

func Test_repository_Create(t *testing.T) {
	type args struct {
		token auth.UserToken
	}
	var (
		now         = mytime.ToTime("2023-05-01 00:00")
		oneAfterDay = now.Add(24 * time.Hour)
		validToken  = auth.NewUserToken(
			1,
			1,
			"token",
			oneAfterDay,
			now,
			sql.NullTime{},
		)
	)
	tests := []struct {
		name    string
		args    args
		want    auth.UserToken
		wantID  int
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				token: validToken,
			},
			wantID: 10003,
			want: auth.NewUserToken(
				10003,
				1,
				"token",
				oneAfterDay,
				now,
				sql.NullTime{}),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tx, _ := mydb.Db.Begin()
			defer func(tx *sql.Tx) {
				_ = tx.Rollback()
			}(tx)
			ctx := mycontext.NewContext(context.Background(), infra.KeyDB, tx)
			ctx = mycontext.NewContext(ctx, infra.KeyTX, tx)

			r := NewRepository()
			got, err := r.Create(ctx, tt.args.token)
			if (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil {
				return
			}

			if diff := cmp.Diff(got, tt.want, cmp.Options{cmp.AllowUnexported(tt.want)}); diff != "" {
				t.Errorf("Find() value is mismatch: %s", diff)
			}
		})
	}
}
