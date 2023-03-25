package zo_test

import (
	"context"
	"database/sql"
	"testing"
	"time"

	d "github.com/ArtefactGitHub/Go_P_Zo/internal/api/v2/domain/zo"
	infra "github.com/ArtefactGitHub/Go_P_Zo/internal/api/v2/infrastructure"
	. "github.com/ArtefactGitHub/Go_P_Zo/internal/api/v2/infrastructure/zo"
	"github.com/ArtefactGitHub/Go_P_Zo/internal/platform/mycontext"
	"github.com/ArtefactGitHub/Go_P_Zo/internal/platform/mydb"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

const (
	layout = "2006-01-02 15:04:05"
)

var (
	findZo = d.NewZo(1,
		stringToTime("2023-01-01 01:00:00"),
		100,
		1,
		"メッセージ1",
		time.Now(),
		sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
		1)
	findZos = []d.Zo{
		d.NewZo(1,
			stringToTime("2023-01-01 01:00:00"),
			100,
			1,
			"メッセージ1",
			time.Now(),
			sql.NullTime{
				Time:  time.Now(),
				Valid: true,
			},
			1),
		d.NewZo(2,
			stringToTime("2023-01-01 02:00:00"),
			200,
			1,
			"メッセージ2",
			time.Now(),
			sql.NullTime{
				Time:  time.Now(),
				Valid: true,
			},
			1),
		d.NewZo(3,
			stringToTime("2023-01-01 03:00:00"),
			300,
			1,
			"メッセージ3",
			time.Now(),
			sql.NullTime{
				Time:  time.Now(),
				Valid: true,
			},
			1),
	}
	createZo = d.NewZo(4,
		stringToTime("2023-02-13 22:48:03"),
		100,
		1,
		"dummy category",
		stringToTime("2023-02-13 22:48:03"),
		sql.NullTime{
			Time:  stringToTime("2023-02-13 22:48:03"),
			Valid: true,
		},
		1)
	updateZo = d.NewZo(2,
		stringToTime("2023-01-01 02:00:00"),
		200,
		1,
		"メッセージ2updated",
		time.Now(),
		sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
		1)
)

func Test_repository_Create(t *testing.T) {
	type args struct {
		z d.Zo
	}

	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				z: createZo,
			},
			want:    4,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tx, _ := mydb.Db.Begin()
			defer tx.Rollback()
			ctx := mycontext.NewContext(context.Background(), infra.KeyTX, tx)

			id, err := NewRepository().Create(ctx, tt.args.z)
			if (err != nil) != tt.wantErr {
				t.Errorf("Create() \nerror = %v\nwantErr %v", err, tt.wantErr)
				return
			}
			if id != tt.want {
				t.Errorf("Create() \nerror = %v\nwantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func Test_repository_Delete(t *testing.T) {
	type args struct {
		id int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				id: 3,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tx, _ := mydb.Db.Begin()
			defer tx.Rollback()
			ctx := mycontext.NewContext(context.Background(), infra.KeyTX, tx)

			r := NewRepository()
			if err := r.Delete(ctx, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_repository_Find(t *testing.T) {
	type args struct {
		id int
	}
	tests := []struct {
		name    string
		args    args
		want    d.Zo
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				id: 1,
			},
			want:    findZo,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := mycontext.NewContext(context.Background(), infra.KeyDB, mydb.Db)

			got, err := NewRepository().Find(ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("Find() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if got == nil {
				t.Errorf("got is nil")
			}

			if diff := cmp.Diff(*got, tt.want, cmpopts.IgnoreFields(d.Zo{}, "CreatedAt", "UpdatedAt")); diff != "" {
				t.Errorf("Find() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func Test_repository_FindAll(t *testing.T) {
	type args struct {
	}
	tests := []struct {
		name    string
		args    args
		want    []d.Zo
		wantErr bool
	}{
		{
			name:    "success",
			args:    args{},
			want:    findZos,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := mycontext.NewContext(context.Background(), infra.KeyDB, mydb.Db)

			got, err := NewRepository().FindAll(ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("Find() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if diff := cmp.Diff(got, tt.want, cmpopts.IgnoreFields(d.Zo{}, "CreatedAt", "UpdatedAt")); diff != "" {
				t.Errorf("Find() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func Test_repository_FindAllByUserId(t *testing.T) {
	type args struct {
		userId int
	}
	tests := []struct {
		name    string
		args    args
		want    []d.Zo
		wantErr bool
	}{
		{
			name:    "success",
			args:    args{},
			want:    findZos,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := mycontext.NewContext(context.Background(), infra.KeyDB, mydb.Db)

			got, err := NewRepository().FindAll(ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("Find() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if diff := cmp.Diff(got, tt.want, cmpopts.IgnoreFields(d.Zo{}, "CreatedAt", "UpdatedAt")); diff != "" {
				t.Errorf("Find() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func Test_repository_Update(t *testing.T) {
	type args struct {
		z d.Zo
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				z: updateZo,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tx, _ := mydb.Db.Begin()
			defer tx.Rollback()
			ctx := mycontext.NewContext(context.Background(), infra.KeyTX, tx)

			if err := NewRepository().Update(ctx, &updateZo); (err != nil) != tt.wantErr {
				t.Errorf("Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func stringToTime(str string) time.Time {
	tz, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		panic(err)
	}
	//t, _ := time.Parse(layout, str)
	t, _ := time.ParseInLocation(layout, str, tz)
	return t
}
