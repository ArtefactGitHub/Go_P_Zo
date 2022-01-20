package user

import (
	"context"
	"database/sql"

	"golang.org/x/crypto/bcrypt"

	"github.com/ArtefactGitHub/Go_P_Zo/internal/platform/mydb"
)

type UserRepository struct {
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

func (r *UserRepository) Create(ctx context.Context, u *User) (int, error) {
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
