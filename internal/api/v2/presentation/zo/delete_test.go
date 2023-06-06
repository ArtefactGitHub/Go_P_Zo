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
	"github.com/ArtefactGitHub/Go_P_Zo/internal/platform/mytime"
	"github.com/ArtefactGitHub/Go_P_Zo/internal/test_v2"
	"github.com/ArtefactGitHub/Go_P_Zo/pkg/common"
	"github.com/google/go-cmp/cmp"
)

func Test_delete_Delete(t *testing.T) {
	var (
		userID = 1
		//dummyUserID = 999
		expiration = time.Now().Add(time.Hour * 24)
		expired    = time.Now().Add(time.Hour * -24)
		postReqest = PostRequest{
			AchievementDate: mytime.ToTime("2023-01-01 01:00"),
			Exp:             100,
			CategoryId:      1,
			Message:         "hoge",
		}
	)
	type fields struct {
		ctx    context.Context
		delete u.Delete
	}
	type args struct {
		body   PostRequest
		params common.QueryMap
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantStatus int
	}{
		{
			name: "正常系：指定の情報で成功",
			fields: fields{
				ctx:    test_v2.WithTXAndTokenContext(context.Background(), DB, TX, userID, expiration),
				delete: u.NewDelete(i.NewRepository()),
			},
			args: args{
				body:   postReqest,
				params: common.QueryMap{resourceKey: "1"},
			},
			wantStatus: http.StatusOK,
		},
		// TODO: zo.UserID検証
		//{
		//	name: "異常系：存在しないユーザー指定の場合",
		//	fields: fields{
		//		ctx:    test_v2.WithTXAndTokenContext(context.Background(), DB, TX, dummyUserID, expiration),
		//		delete: u.NewDelete(i.NewRepository()),
		//	},
		//	args: args{
		//		body: postReqest,
		//	},
		//	want: PostResponse{
		//		ResponseBase: &myhttp.ResponseBase{
		//			StatusCode: http.StatusBadRequest,
		//			Error:      nil,
		//		},
		//		Zo: Zo{},
		//	},
		//	wantStatus: http.StatusBadRequest,
		//},
		{
			name: "異常系：リソースが存在しない場合",
			fields: fields{
				ctx:    context.Background(),
				delete: u.NewDelete(i.NewRepository()),
			},
			args: args{
				body:   postReqest,
				params: common.QueryMap{resourceKey: "999"},
			},
			wantStatus: http.StatusUnauthorized,
		},
		{
			name: "異常系：トークンが存在しない場合",
			fields: fields{
				ctx:    context.Background(),
				delete: u.NewDelete(i.NewRepository()),
			},
			args: args{
				body: postReqest,
			},
			wantStatus: http.StatusUnauthorized,
		},
		{
			name: "正常系：トークンが有効期限切れの場合",
			fields: fields{
				ctx:    test_v2.WithTXAndTokenContext(context.Background(), DB, TX, userID, expired),
				delete: u.NewDelete(i.NewRepository()),
			},
			args: args{
				body: postReqest,
			},
			wantStatus: http.StatusUnauthorized,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodDelete, "/api/v2/zo"+test_v2.GetResourceIdStr(tt.args.params, resourceKey), nil)

			// テストケースに応じたcontextをセットする
			req := r.WithContext(tt.fields.ctx)

			h := NewDelete(tt.fields.delete)
			h.Delete(w, req, tt.args.params)

			if diff := cmp.Diff(tt.wantStatus, w.Code); diff != "" {
				t.Errorf("handler.Delete mismatch (-wantStatus +got):\n%s", diff)
			}
		})
	}
}
