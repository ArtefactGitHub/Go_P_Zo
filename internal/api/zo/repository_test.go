package zo

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/ArtefactGitHub/Go_P_Zo/internal/platform/mydb"
	"github.com/ArtefactGitHub/Go_P_Zo/internal/test"
)

var r_tests = map[string]func(t *testing.T){
	"test_r_findall": test_r_findall,
	"test_r_find":    test_r_find,
	"test_r_create":  test_r_create,
	"test_r_update":  test_r_update,
	"test_r_delete":  test_r_delete}

var ac, _ = time.Parse(test.TimeLayout, "2021-12-18")
var seeds []zo = []zo{
	newZo(1, ac, 100, 0, "test-1", time.Now(), sql.NullTime{}),
	newZo(2, ac, 200, 0, "test-2", time.Now(), sql.NullTime{}),
	newZo(3, ac, 300, 0, "test-3", time.Now(), sql.NullTime{})}

func Test_repository(t *testing.T) {
	test.Run(t, r_tests, nil, nil, test_seed)
}

// findall()のテスト
func test_r_findall(t *testing.T) {
	r := zoRepository{}
	zos, err := r.findall()
	if err != nil {
		t.Errorf("findall() has error: %v", err)
	}

	want := 3
	if len(zos) != want {
		t.Errorf("len() = %d, want %d", len(zos), want)
	}
}

// find()のテスト
func test_r_find(t *testing.T) {
	r := zoRepository{}
	z, err := r.find(1)
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
	r := zoRepository{}
	z := seeds[0]
	z.Message = "created by test"
	id, err := r.create(&z)
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

	if count != cap(seeds)+1 {
		t.Errorf("count = %d, want %d", count, cap(seeds)+1)
	}
}

// update()のテスト
func test_r_update(t *testing.T) {
	r := zoRepository{}
	z := seeds[0]
	z.Exp = 500
	err := r.update(&z)
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
	r := zoRepository{}
	z := seeds[0]
	err := r.delete(z.Id)
	if err != nil {
		t.Fatalf("delete() has error: %v", err)
	}

	var count int
	want := 2
	err = mydb.Db.QueryRow("SELECT COUNT(*) FROM zos").Scan(&count)
	if err != nil {
		t.Fatalf("delete() has error: %v", err)
	}

	if count != want {
		t.Errorf("count = %d, want %d", count, want)
	}
}

func test_seed(ctx context.Context) {
	tx, err := mydb.Db.BeginTx(ctx, nil)
	if err != nil {
		test.Failuer(err)
	}
	// Defer a rollback in case anything fails.
	// https://go.dev/doc/database/execute-transactions
	defer tx.Rollback()

	for _, z := range seeds {
		_, err := mydb.Db.ExecContext(
			ctx,
			`INSERT INTO zos(id, achievementDate, exp, categoryId, message, createdAt, updatedAt)
		            values(?, ?, ?, ?, ?, ?, ?)`,
			nil,
			z.AchievementDate,
			z.Exp,
			z.CategoryId,
			z.Message,
			z.CreatedAt,
			z.UpdatedAt)
		if err != nil {
			test.Failuer(err)
		}
	}

	// Commit the transaction.
	if err = tx.Commit(); err != nil {
		test.Failuer(err)
	}
}
