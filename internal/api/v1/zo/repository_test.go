package zo

import (
	"context"
	"testing"

	"github.com/ArtefactGitHub/Go_P_Zo/internal/platform/mydb"
	"github.com/ArtefactGitHub/Go_P_Zo/internal/test"
)

var r_tests = map[string]func(t *testing.T){
	"test_r_findall":           test_r_findall,
	"test_r_findall_by_userid": test_r_findall_by_userid,
	"test_r_find":              test_r_find,
	"test_r_create":            test_r_create,
	"test_r_update":            test_r_update,
	"test_r_delete":            test_r_delete}

func Test_repository(t *testing.T) {
	test.Run(t, r_tests, nil, nil, test_seed)
}

// findall()のテスト
func test_r_findall(t *testing.T) {
	r := ZoRepository{}
	zos, err := r.FindAll(context.Background())
	if err != nil {
		t.Errorf("findall() has error: %v", err)
	}

	want := len(seeds)
	if len(zos) != want {
		t.Errorf("len() = %d, want %d", len(zos), want)
	}
}

func test_r_findall_by_userid(t *testing.T) {
	r := ZoRepository{}
	userId := userId_2
	zos, err := r.FindAllByUserId(context.Background(), userId)
	if err != nil {
		t.Errorf("FindAllByUserId() has error: %v", err)
	}

	want := 1
	if len(zos) != want {
		t.Errorf("len() = %d, want %d", len(zos), want)
	}
}

// find()のテスト
func test_r_find(t *testing.T) {
	r := ZoRepository{}
	z, err := r.Find(context.Background(), 1)
	if err != nil {
		t.Errorf("findall() has error: %v", err)
	}

	if z.Exp != 100 {
		t.Errorf("exp = %d, want %d", z.Exp, 100)
	}

	if z.Message != "test-1" {
		t.Errorf("exp = %s, want %s", z.Message, "test-1")
	}
}

// create()のテスト
func test_r_create(t *testing.T) {
	r := ZoRepository{}
	z := seeds[0]
	z.Message = "created by test"
	id, err := r.Create(context.Background(), &z)
	if err != nil {
		t.Fatalf("create() has error: %v", err)
	}

	want := "created by test"
	var message string
	err = mydb.Db.QueryRow("SELECT * FROM zos WHERE id = ?", id).Scan(
		&test.TrashScanner{},
		&test.TrashScanner{},
		&test.TrashScanner{},
		&test.TrashScanner{},
		&message,
		&test.TrashScanner{},
		&test.TrashScanner{},
		&test.TrashScanner{})
	if err != nil {
		t.Fatalf("create() has error: %v", err)
	}

	if message != want {
		t.Errorf("z.Message = %s, want %s", z.Message, want)
	}

	var count int
	err = mydb.Db.QueryRow("SELECT COUNT(*) FROM zos").Scan(&count)
	if err != nil {
		t.Fatalf("create() has error: %v", err)
	}

	if count != len(seeds)+1 {
		t.Errorf("count = %d, want %d", count, len(seeds)+1)
	}
}

// update()のテスト
func test_r_update(t *testing.T) {
	r := ZoRepository{}
	z := seeds[0]
	z.Exp = 500
	err := r.Update(context.Background(), &z)
	if err != nil {
		t.Fatalf("update() has error: %v", err)
	}

	var want int
	err = mydb.Db.QueryRow("SELECT * FROM zos WHERE id = ?", z.Id).Scan(
		&test.TrashScanner{},
		&test.TrashScanner{},
		&want,
		&test.TrashScanner{},
		&test.TrashScanner{},
		&test.TrashScanner{},
		&test.TrashScanner{},
		&test.TrashScanner{})
	if err != nil {
		t.Fatalf("update() has error: %v", err)
	}

	if z.Exp != want {
		t.Errorf("z.Exp = %d, want %d", z.Exp, want)
	}
}

// delete()のテスト
func test_r_delete(t *testing.T) {
	r := ZoRepository{}
	z := seeds[0]
	err := r.Delete(context.Background(), z.Id)
	if err != nil {
		t.Fatalf("delete() has error: %v", err)
	}

	var count int
	want := len(seeds) - 1
	err = mydb.Db.QueryRow("SELECT COUNT(*) FROM zos").Scan(&count)
	if err != nil {
		t.Fatalf("delete() has error: %v", err)
	}

	if count != want {
		t.Errorf("count = %d, want %d", count, want)
	}
}
