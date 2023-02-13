package session

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	d "github.com/ArtefactGitHub/Go_P_Zo/internal/api/v2/domain/session"
)

func Test_repository_Find(t *testing.T) {
	type args struct {
		ctx        context.Context
		identifier string
		password   string
	}
	tests := []struct {
		name    string
		args    args
		want    d.SessionData
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				ctx:        context.Background(),
				identifier: "test@com",
				password:   "password",
			},
			want:    d.NewSessionData("hoge", "fuga", "test@com"),
			wantErr: false,
		},
		{
			name: "not found",
			args: args{
				ctx:        context.Background(),
				identifier: "dummy",
				password:   "dummy",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRepository()
			fmt.Printf("pass: %#v \n", tt.args.password)
			got, err := r.Find(tt.args.ctx, tt.args.identifier, tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("Find() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Find() got = %v, want %v", got, tt.want)
			}
		})
	}
}
