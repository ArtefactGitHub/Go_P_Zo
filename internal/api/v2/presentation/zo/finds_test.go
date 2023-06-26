package zo_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	i "github.com/ArtefactGitHub/Go_P_Zo/internal/api/v2/infrastructure/zo"
	. "github.com/ArtefactGitHub/Go_P_Zo/internal/api/v2/presentation/zo"
	u "github.com/ArtefactGitHub/Go_P_Zo/internal/api/v2/usecase/zo"
	"github.com/ArtefactGitHub/Go_P_Zo/internal/test_v2"
	"github.com/google/go-cmp/cmp"
)

func Test_finds_Finds(t *testing.T) {
	type fields struct {
		ctx   context.Context
		finds u.Finds
	}
	type args struct {
	}
	var (
		expiration = time.Now().Add(time.Hour * 24)
		expired    = time.Now().Add(time.Hour * -24)
	)
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantStatus int
	}{
		{
			name: "【正常系】リソースが見つかった場合",
			fields: fields{
				ctx:   test_v2.WithDBAndTokenContext(context.Background(), DB, 1, expiration),
				finds: u.NewFinds(i.NewRepository()),
			},
			args:       args{},
			wantStatus: http.StatusOK,
		},
		{
			name: "【異常系】トークンが存在しない場合",
			fields: fields{
				ctx:   context.Background(),
				finds: u.NewFinds(i.NewRepository()),
			},
			args:       args{},
			wantStatus: http.StatusUnauthorized,
		},
		{
			name: "【異常系】非リソース所有の場合",
			fields: fields{
				ctx:   test_v2.WithDBAndTokenContext(context.Background(), DB, 999, expiration),
				finds: u.NewFinds(i.NewRepository()),
			},
			args:       args{},
			wantStatus: http.StatusOK,
		},
		{
			name: "【異常系】トークンが有効期限切れの場合",
			fields: fields{
				ctx:   test_v2.WithDBAndTokenContext(context.Background(), DB, 1, expired),
				finds: u.NewFinds(i.NewRepository()),
			},
			args:       args{},
			wantStatus: http.StatusUnauthorized,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodGet, "/api/v2/zo", nil)

			// テストケースに応じたcontextをセットする
			req := r.WithContext(tt.fields.ctx)

			h := NewFinds(tt.fields.finds)
			h.Finds(w, req, nil)

			if diff := cmp.Diff(tt.wantStatus, w.Code); diff != "" {
				t.Errorf("handler.Find mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
