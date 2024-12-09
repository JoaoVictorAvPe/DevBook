package repositories

import (
	"api/src/models"
	"database/sql"
)

type userRepo struct {
	db *sql.DB
}

func NewUserRepositorie(db *sql.DB) *userRepo {
	return &userRepo{db}
}

func (ur userRepo) Add(user models.User) (uint64, error) {
	statement, err := ur.db.Prepare(
		"INSERT INTO user (name, nick, email, password) VALUES(?, ?, ?, ?)",
	)
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	result, err := statement.Exec(user.Name, user.Nick, user.Email, user.Password)
	if err != nil {
		return 0, err
	}

	lasInsertId, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(lasInsertId), nil
}
