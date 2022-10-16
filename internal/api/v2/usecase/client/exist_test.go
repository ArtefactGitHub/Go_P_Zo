package client_test

import (
	"context"
	domain "github.com/ArtefactGitHub/Go_P_Zo/internal/api/v2/domain/client"
	infra "github.com/ArtefactGitHub/Go_P_Zo/internal/api/v2/infrastructure/client"
	. "github.com/ArtefactGitHub/Go_P_Zo/internal/api/v2/usecase/client"
	"testing"
)

func Test_exist_Do(t *testing.T) {
	type fields struct {
		r domain.Repository
	}
	type args struct {
		ctx    context.Context
		id     int
		secret string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bool
		wantErr bool
	}{
		{
			name:   "success",
			fields: fields{r: infra.NewRepository()},
			args: args{
				ctx:    context.Background(),
				id:     1,
				secret: "secret-1",
			},
			want:    true,
			wantErr: false,
		},
		{
			name:   "not found",
			fields: fields{r: infra.NewRepository()},
			args: args{
				ctx:    context.Background(),
				id:     0,
				secret: "dummy",
			},
			want:    false,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := NewExist(tt.fields.r)
			got, err := u.Do(tt.args.ctx, tt.args.id, tt.args.secret)
			if (err != nil) != tt.wantErr {
				t.Errorf("Do() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Do() got = %v, want %v", got, tt.want)
			}
		})
	}
}
