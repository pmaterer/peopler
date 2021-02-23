package repository

import (
	"database/sql"

	"github.com/pmaterer/peopler/user"
)

type Reopository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Reopository {
	return &Reopository{
		db: db,
	}
}

func (r *Reopository) CreateUser(u user.User) (int64, error) {
	var id int64
	query := `INSERT INTO users(first_name, last_name) VALUES (?, ?)`
	statement, err := r.db.Prepare(query)
	if err != nil {
		return id, err
	}
	row, err := statement.Exec(u.FirstName, u.LastName)
	if err != nil {
		return id, err
	}
	id, err = row.LastInsertId()
	if err != nil {
		return id, err
	}
	return id, nil
}

func (r *Reopository) GetAllUsers() ([]user.User, error) {
	var users []user.User
	rows, err := r.db.Query(`SELECT * FROM users`)
	if err != nil {
		return users, err
	}
	defer rows.Close()

	for rows.Next() {
		var user user.User
		err = rows.Scan(&user.ID, &user.FirstName, &user.LastName)
		if err != nil {
			return users, err
		}
		users = append(users, user)
	}
	err = rows.Err()
	if err != nil {
		return users, err
	}
	return users, nil
}

func (r *Reopository) GetUser(id int64) (user.User, error) {
	var user user.User
	err := r.db.QueryRow(`SELECT id, first_name, last_name FROM users WHERE id = ?`, id).Scan(&user.ID, &user.FirstName, &user.LastName)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (r *Reopository) UpdateUser(u user.User) (int64, error) {
	var id int64
	query := `UPDATE users SET first_name=?, last_name=? WHERE id=?`
	statement, err := r.db.Prepare(query)
	if err != nil {
		return id, err
	}
	row, err := statement.Exec(u.FirstName, u.LastName, u.ID)
	if err != nil {
		return id, err
	}
	id, err = row.LastInsertId()
	if err != nil {
		return id, err
	}
	return id, nil
}

func (r *Reopository) DeleteUser(id int64) (int64, error) {
	query := `DELETE FROM users WHERE id=?`
	statement, err := r.db.Prepare(query)
	if err != nil {
		return id, err
	}
	row, err := statement.Exec(id)
	if err != nil {
		return id, err
	}
	id, err = row.LastInsertId()
	if err != nil {
		return id, err
	}
	return id, nil
}
