package client_test

import (
	"context"
	"database/sql"
	"testing"
	"time"

	domain "github.com/ArtefactGitHub/Go_P_Zo/internal/api/v2/domain/client"
	. "github.com/ArtefactGitHub/Go_P_Zo/internal/api/v2/infrastructure/client"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func Test_repository_Find(t *testing.T) {
	type args struct {
		ctx    context.Context
		id     int
		secret string
	}
	tests := []struct {
		name    string
		args    args
		want    domain.Client
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				ctx:    context.Background(),
				id:     1,
				secret: "secret-1",
			},
			want:    domain.NewClient(1, "secret-1", time.Now(), sql.NullTime{}),
			wantErr: false,
		},
		{
			name: "not found",
			args: args{
				ctx:    context.Background(),
				id:     0,
				secret: "dummy",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRepository()
			got, err := r.Find(tt.args.ctx, tt.args.id, tt.args.secret)
			if (err != nil) != tt.wantErr {
				t.Errorf("Find() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err == nil {
				opts := cmp.Options{
					cmp.AllowUnexported(got),
					cmpopts.IgnoreFields(got, "createdAt", "updatedAt"),
				}
				if diff := cmp.Diff(got, tt.want, opts); diff != "" {
					t.Errorf("Find() value is mismatch: %s", diff)
				}
			}
		})
	}
}
