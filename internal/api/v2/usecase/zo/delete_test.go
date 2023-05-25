package zo_test

import (
	"context"
	"testing"

	"github.com/ArtefactGitHub/Go_P_Zo/internal/api/v2/domain/zo"
	i "github.com/ArtefactGitHub/Go_P_Zo/internal/api/v2/infrastructure"
	inZo "github.com/ArtefactGitHub/Go_P_Zo/internal/api/v2/infrastructure/zo"
	. "github.com/ArtefactGitHub/Go_P_Zo/internal/api/v2/usecase/zo"
)

func Test_delete_Do(t *testing.T) {
	type fields struct {
		r zo.Repository
	}
	type args struct {
		target int
	}
	var ()
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:    "正常系：正しく更新が行える",
			fields:  fields{r: inZo.NewRepository()},
			args:    args{target: 1},
			wantErr: false,
		},
		{
			name:    "異常系：存在しないIDを指定する",
			fields:  fields{r: inZo.NewRepository()},
			args:    args{target: 999},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := NewDelete(tt.fields.r)
			c := context.WithValue(context.Background(), i.KeyDB, DB)
			c = context.WithValue(c, i.KeyTX, TX)
			err := u.Do(c, tt.args.target)
			if (err != nil) != tt.wantErr {
				t.Errorf("Do() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
