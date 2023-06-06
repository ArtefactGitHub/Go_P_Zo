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
	"github.com/ArtefactGitHub/Go_P_Zo/pkg/common"
	"github.com/google/go-cmp/cmp"
)

const resourceKey = "zo_id"

func Test_find_Find(t *testing.T) {
	var (
		expiration = time.Now().Add(time.Hour * 24)
		expired    = time.Now().Add(time.Hour * -24)
	)
	type fields struct {
		ctx  context.Context
		find u.Find
	}
	type args struct {
		params common.QueryMap
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantStatus int
	}{
		{
			name: "【正常系】指定の1件が見つかった場合",
			fields: fields{
				ctx:  test_v2.WithDBAndTokenContext(context.Background(), DB, 1, expiration),
				find: u.NewFind(i.NewRepository()),
			},
			args: args{
				params: common.QueryMap{resourceKey: "1"},
			},
			wantStatus: http.StatusOK,
		},
		{
			name: "【異常系】指定の1件が見つからない場合",
			fields: fields{
				ctx:  test_v2.WithDBAndTokenContext(context.Background(), DB, 1, expiration),
				find: u.NewFind(i.NewRepository()),
			},
			args: args{
				params: common.QueryMap{resourceKey: "999"},
			},
			wantStatus: http.StatusNotFound,
		},
		{
			name: "【異常系】トークンが存在しない場合",
			fields: fields{
				ctx:  context.Background(),
				find: u.NewFind(i.NewRepository()),
			},
			args: args{
				params: common.QueryMap{resourceKey: "999"},
			},
			wantStatus: http.StatusUnauthorized,
		},
		{
			name: "【異常系】非リソース所有者の場合",
			fields: fields{
				ctx:  test_v2.WithDBAndTokenContext(context.Background(), DB, 999, expiration),
				find: u.NewFind(i.NewRepository()),
			},
			args: args{
				params: common.QueryMap{resourceKey: "1"},
			},
			wantStatus: http.StatusUnauthorized,
		},
		{
			name: "【異常系】トークンが有効期限切れの場合",
			fields: fields{
				ctx:  test_v2.WithDBAndTokenContext(context.Background(), DB, 1, expired),
				find: u.NewFind(i.NewRepository()),
			},
			args: args{
				params: common.QueryMap{resourceKey: "1"},
			},
			wantStatus: http.StatusUnauthorized,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodGet, "/api/v2/zo/"+test_v2.GetResourceIdStr(tt.args.params, resourceKey), nil)

			// テストケースに応じたcontextをセットする
			req := r.WithContext(tt.fields.ctx)

			h := NewFind(tt.fields.find)
			h.Find(w, req, tt.args.params)

			if diff := cmp.Diff(tt.wantStatus, w.Code); diff != "" {
				t.Errorf("handler.Find mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
