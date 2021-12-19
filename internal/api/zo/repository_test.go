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
	"test_r_update":  test_r_update}

var ac, _ = time.Parse(test.TimeLayout, "2021-12-18")
var seeds []zo = []zo{
	newZo(1, ac, 100, 0, "test-1", time.Now(), sql.NullTime{}),
	newZo(2, ac, 200, 0, "test-2", time.Now(), sql.NullTime{}),
	newZo(3, ac, 300, 0, "test-3", time.Now(), sql.NullTime{})}

func Test_repository(t *testing.T) {
	test.Run(t, r_tests, test_r_before, nil)
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

func test_r_before() {
	_, err := mydb.Db.Exec("TRUNCATE zos")
	if err != nil {
		test.Failuer(err)
	}

	test_r_seed()
}

func test_r_seed() {
	ctx := context.Background()
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
