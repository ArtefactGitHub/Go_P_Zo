package zo_test

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	i "github.com/ArtefactGitHub/Go_P_Zo/internal/api/v2/infrastructure/zo"
	. "github.com/ArtefactGitHub/Go_P_Zo/internal/api/v2/presentation/zo"
	u "github.com/ArtefactGitHub/Go_P_Zo/internal/api/v2/usecase/zo"
	"github.com/ArtefactGitHub/Go_P_Zo/internal/platform/myhttp"
	"github.com/ArtefactGitHub/Go_P_Zo/internal/platform/mytime"
	"github.com/ArtefactGitHub/Go_P_Zo/internal/test_v2"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func Test_create_Create(t *testing.T) {
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
		create u.Create
	}
	type args struct {
		body PostRequest
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		want       PostResponse
		wantStatus int
	}{
		{
			name: "【正常系】指定の情報で成功",
			fields: fields{
				ctx:    test_v2.WithTXAndTokenContext(context.Background(), DB, TX, userID, expiration),
				create: u.NewCreate(i.NewRepository()),
			},
			args: args{
				body: postReqest,
			},
			want: PostResponse{
				ResponseBase: &myhttp.ResponseBase{
					StatusCode: http.StatusOK,
					Error:      nil,
				},
				Zo: Zo{
					AchievementDate: postReqest.AchievementDate,
					Exp:             postReqest.Exp,
					CategoryId:      postReqest.CategoryId,
					Message:         postReqest.Message,
					UserId:          userID,
				},
			},
			wantStatus: http.StatusOK,
		},
		// TODO: zo.UserID検証
		//{
		//	name: "【異常系】存在しないユーザー指定の場合",
		//	fields: fields{
		//		ctx:    test_v2.WithTXAndTokenContext(context.Background(), DB, TX, dummyUserID, expiration),
		//		create: u.NewCreate(i.NewRepository()),
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
			name: "【異常系】トークンが存在しない場合",
			fields: fields{
				ctx:    context.Background(),
				create: u.NewCreate(i.NewRepository()),
			},
			args: args{
				body: postReqest,
			},
			want: PostResponse{
				ResponseBase: &myhttp.ResponseBase{
					StatusCode: http.StatusUnauthorized,
					Error:      nil,
				},
				Zo: Zo{},
			},
			wantStatus: http.StatusUnauthorized,
		},
		{
			name: "【正常系】トークンが有効期限切れの場合",
			fields: fields{
				ctx:    test_v2.WithTXAndTokenContext(context.Background(), DB, TX, userID, expired),
				create: u.NewCreate(i.NewRepository()),
			},
			args: args{
				body: postReqest,
			},
			want: PostResponse{
				ResponseBase: &myhttp.ResponseBase{
					StatusCode: http.StatusUnauthorized,
					Error:      nil,
				},
				Zo: Zo{},
			},
			wantStatus: http.StatusUnauthorized,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			js, err := json.Marshal(tt.args.body)
			if err != nil {
				t.Fatalf("request body error: %v", err.Error())
			}
			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodPost, "/api/v2/zo", bytes.NewBuffer(js))

			// テストケースに応じたcontextをセットする
			req := r.WithContext(tt.fields.ctx)

			h := NewCreate(tt.fields.create)
			h.Create(w, req, nil)

			if diff := cmp.Diff(tt.wantStatus, w.Code); diff != "" {
				t.Errorf("handler.Create mismatch (-wantStatus +got):\n%s", diff)
			}

			// レスポンス値の検証
			var res PostResponse
			if err = json.Unmarshal(w.Body.Bytes(), &res); err != nil {
				t.Fatalf(err.Error())
			}
			opts := cmp.Options{
				cmpopts.IgnoreFields(Zo{}, "CreatedAt", "UpdatedAt"),
				cmpopts.IgnoreFields(myhttp.ResponseBase{}, "Error"),
			}
			if diff := cmp.Diff(tt.want, res, opts); diff != "" {
				t.Errorf("handler.Create mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
