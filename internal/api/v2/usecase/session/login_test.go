package session_test

import (
	"context"
	"reflect"
	"testing"

	d "github.com/ArtefactGitHub/Go_P_Zo/internal/api/v2/domain/session"
	infra "github.com/ArtefactGitHub/Go_P_Zo/internal/api/v2/infrastructure/session"
	. "github.com/ArtefactGitHub/Go_P_Zo/internal/api/v2/usecase/session"
)

func Test_login_Do(t *testing.T) {
	type fields struct {
		r d.Repository
	}
	type args struct {
		ctx        context.Context
		identifier string
		secret     string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    d.SessionData
		wantErr bool
	}{
		{
			name:   "success",
			fields: fields{r: infra.NewRepository()},
			args: args{
				ctx:        context.Background(),
				identifier: "test@com",
				secret:     "password",
			},
			want:    d.NewSessionData("hoge", "fuga", "test@com"),
			wantErr: false,
		},
		{
			name:   "not found",
			fields: fields{r: infra.NewRepository()},
			args: args{
				ctx:        context.Background(),
				identifier: "dummy",
				secret:     "dummy",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewLogin(tt.fields.r)
			got, err := s.Do(tt.args.ctx, tt.args.identifier, tt.args.secret)
			if (err != nil) != tt.wantErr {
				t.Errorf("Do() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Do() got = %v, want %v", got, tt.want)
			}
		})
	}
}
