package user

import (
	"context"
	"database/sql"
	"fmt"

	"golang.org/x/crypto/bcrypt"

	"github.com/ArtefactGitHub/Go_P_Zo/internal/platform/mydb"
)

type UserRepository struct {
}

type UserTokenRepository struct {
}

type userCategoryRepository struct {
}

func (r *UserRepository) FindAll(ctx context.Context) ([]User, error) {
	rows, err := mydb.Db.QueryContext(ctx, "SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var u User
	var result []User
	for rows.Next() {
		err := rows.Scan(
			&u.Id,
			&u.GivenName,
			&u.FamilyName,
			&u.Email,
			&u.Password,
			&u.CreatedAt,
			&u.UpdatedAt)
		if err != nil {
			return nil, err
		}

		result = append(result, u)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return result, nil
}

func (r *UserRepository) Find(ctx context.Context, id int) (*User, error) {
	u := User{}
	err := mydb.Db.QueryRowContext(ctx, "SELECT * FROM users WHERE id = ?", id).Scan(
		&u.Id,
		&u.GivenName,
		&u.FamilyName,
		&u.Email,
		&u.Password,
		&u.CreatedAt,
		&u.UpdatedAt)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return &u, nil
}

func (r *UserRepository) FindByEmail(ctx context.Context, email string) (*User, error) {
	u := User{}
	err := mydb.Db.QueryRowContext(ctx, "SELECT * FROM users WHERE email = ?", email).Scan(
		&u.Id,
		&u.GivenName,
		&u.FamilyName,
		&u.Email,
		&u.Password,
		&u.CreatedAt,
		&u.UpdatedAt)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return &u, nil
}

func (r *UserRepository) FindByEmailTx(ctx context.Context, tx *sql.Tx, email string) (*User, error) {
	u := User{}
	err := tx.QueryRowContext(ctx, "SELECT * FROM users WHERE email = ?", email).Scan(
		&u.Id,
		&u.GivenName,
		&u.FamilyName,
		&u.Email,
		&u.Password,
		&u.CreatedAt,
		&u.UpdatedAt)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return &u, nil
}

func (r *UserRepository) FindByIdentifier(ctx context.Context, identifier string, password string) (*User, error) {
	result := &User{}
	err := mydb.Db.QueryRowContext(ctx,
		"SELECT * FROM users WHERE email = ?", identifier).
		Scan(
			&result.Id,
			&result.GivenName,
			&result.FamilyName,
			&result.Email,
			&result.Password,
			&result.CreatedAt,
			&result.UpdatedAt)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(password))
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *UserRepository) Create(ctx context.Context, u *User) (int, error) {
	err := r.checkNotExist(ctx, nil, u)
	if err != nil {
		return -1, err
	}

	password := []byte(u.Password)
	hashed, err := bcrypt.GenerateFromPassword(password, 12)
	if err != nil {
		return -1, err
	}

	result, err := mydb.Db.ExecContext(ctx, `
			INSERT INTO users(id, given_name, family_name, email, password, createdAt, updatedAt)
			values(?, ?, ?, ?, ?, ?, ?)`,
		nil,
		&u.GivenName,
		&u.FamilyName,
		&u.Email,
		&hashed,
		&u.CreatedAt,
		&u.UpdatedAt)
	if err != nil {
		return -1, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return -1, err
	}

	u.Id = int(id)
	return u.Id, nil
}

func (r *UserRepository) CreateTx(ctx context.Context, tx *sql.Tx, u *User) (int, error) {
	err := r.checkNotExist(ctx, tx, u)
	if err != nil {
		return -1, err
	}

	password := []byte(u.Password)
	hashed, err := bcrypt.GenerateFromPassword(password, 12)
	if err != nil {
		return -1, err
	}

	result, err := tx.ExecContext(ctx, `
			INSERT INTO users(id, given_name, family_name, email, password, createdAt, updatedAt)
			values(?, ?, ?, ?, ?, ?)`,
		nil,
		&u.GivenName,
		&u.FamilyName,
		&u.Email,
		&hashed,
		&u.CreatedAt,
		&u.UpdatedAt)
	if err != nil {
		return -1, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return -1, err
	}

	u.Id = int(id)
	return u.Id, nil
}

func (r *UserRepository) Update(ctx context.Context, u *User) error {
	_, err := mydb.Db.ExecContext(ctx, `
		UPDATE users
		SET given_name = ?,
				family_name = ?,
				email = ?,
		    updatedAt = ?
		WHERE id = ?`,
		u.GivenName,
		u.FamilyName,
		u.Email,
		u.UpdatedAt,
		u.Id)
	if err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) Delete(ctx context.Context, id int) error {
	_, err := mydb.Db.ExecContext(ctx, `
		DELETE FROM users
		WHERE id = ?`,
		id)
	if err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) checkNotExist(ctx context.Context, tx *sql.Tx, u *User) error {
	var exist *User
	var err error
	if tx == nil {
		exist, err = r.FindByEmail(ctx, u.Email)
	} else {
		exist, err = r.FindByEmailTx(ctx, tx, u.Email)
	}

	if err != nil {
		return err
	} else if exist != nil {
		return fmt.Errorf("email address has already been registered: %s", u.Email)
	}

	return nil
}

// UserTokenRepository
func (r *UserTokenRepository) Create(ctx context.Context, m *UserToken) (int, error) {
	result, err := mydb.Db.ExecContext(ctx, `
		INSERT INTO UserTokens(id, user_id, token, expiredAt, createdAt, updatedAt)
			VALUES(?,?,?,?,?,?)`,
		nil,
		&m.UserId,
		&m.Token,
		&m.ExpiredAt,
		&m.CreatedAt,
		&m.UpdatedAt)
	if err != nil {
		return -1, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return -1, err
	}

	m.Id = int(id)
	return m.Id, nil
}

// userCategoryRepository
func (r *userCategoryRepository) FindAllByUserId(ctx context.Context, userId int) ([]UserCategory, error) {
	rows, err := mydb.Db.QueryContext(ctx, "SELECT * FROM usercategories WHERE user_id = ? order by id", userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var m UserCategory
	var result []UserCategory
	for rows.Next() {
		err := rows.Scan(
			&m.Id,
			&m.Number,
			&m.Name,
			&m.ColorId,
			&m.UserId,
			&m.CreatedAt,
			&m.UpdatedAt)
		if err != nil {
			return nil, err
		}

		result = append(result, m)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return result, nil
}

func (r *userCategoryRepository) Create(ctx context.Context, m *UserCategory) (int, error) {
	result, err := mydb.Db.ExecContext(ctx, `
			INSERT INTO UserCategories(id, number, name, color_id, user_id, createdAt, updatedAt)
			VALUES(?, ?, ?, ?, ?, ?, ?)`,
		nil,
		&m.Number,
		&m.Name,
		&m.ColorId,
		&m.UserId,
		&m.CreatedAt,
		&m.UpdatedAt)
	if err != nil {
		return -1, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return -1, err
	}

	m.Id = int(id)
	return m.Id, nil
}
