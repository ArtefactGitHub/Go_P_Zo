package session

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	i "github.com/ArtefactGitHub/Go_P_Zo/internal/api/v2/infrastructure/session"
	"github.com/ArtefactGitHub/Go_P_Zo/internal/api/v2/usecase/session"
	u "github.com/ArtefactGitHub/Go_P_Zo/internal/api/v2/usecase/session"
)

func Test_login_Create(t *testing.T) {
	type fields struct {
		login session.Login
	}
	type args struct {
		body PostRequest
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantCode int
	}{
		{
			name: "success",
			fields: fields{
				login: u.NewLogin(i.NewRepository()),
			},
			args: args{
				body: PostRequest{
					Identifier: "test@com",
					Secret:     "password",
				},
			},
			wantCode: http.StatusOK,
		},
		{
			name: "failed with wrong ID",
			fields: fields{
				login: u.NewLogin(i.NewRepository()),
			},
			args: args{
				body: PostRequest{
					Identifier: "dummy",
					Secret:     "password",
				},
			},
			wantCode: http.StatusNotFound,
		},
		{
			name: "failed with wrong secret",
			fields: fields{
				login: u.NewLogin(i.NewRepository()),
			},
			args: args{
				body: PostRequest{
					Identifier: "test@com",
					Secret:     "dummy",
				},
			},
			wantCode: http.StatusBadRequest,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			js, err := json.Marshal(tt.args.body)
			if err != nil {
				t.Fatalf("request body error: %v", err.Error())
			}
			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodPost, "/api/v2/login", bytes.NewBuffer(js))

			h := NewLogin(tt.fields.login)
			h.Create(w, r, nil)

			if w.Code != tt.wantCode {
				t.Errorf("code = %v, wantCode = %v", w.Code, tt.wantCode)
			}

			var j PostResponse
			if err = json.Unmarshal(w.Body.Bytes(), &j); err != nil {
				t.Fatalf(err.Error())
			}
		})
	}
}
