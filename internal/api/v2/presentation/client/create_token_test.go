package client_test

import (
	"bytes"
	"encoding/json"
	i "github.com/ArtefactGitHub/Go_P_Zo/internal/api/v2/infrastructure/client"
	. "github.com/ArtefactGitHub/Go_P_Zo/internal/api/v2/presentation/client"
	u "github.com/ArtefactGitHub/Go_P_Zo/internal/api/v2/usecase/client"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_createToken_Create(t *testing.T) {
	type fields struct {
		exist       u.Exist
		createToken u.CreateToken
	}
	type args struct {
		body PostRequest
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "success",
			fields: fields{
				exist:       u.NewExist(i.NewRepository()),
				createToken: u.NewCreateToken(),
			},
			args: args{
				body: PostRequest{
					Id:     1,
					Secret: "secret-1",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			js, err := json.Marshal(tt.args.body)
			if err != nil {
				t.Fatalf("request body error: %v", err.Error())
			}
			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodPost, "/api/v2/client/token", bytes.NewBuffer(js))

			h := NewCreateToken(tt.fields.exist, tt.fields.createToken)
			h.Create(w, r, nil)

			var j PostResponse
			if err = json.Unmarshal(w.Body.Bytes(), &j); err != nil {
				t.Fatalf(err.Error())
			}
		})
	}
}
