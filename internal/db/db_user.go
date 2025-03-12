package db

import (
	"database/sql"
	"log"
	mestoTypes "mesto-goback/internal/types"

	"github.com/google/uuid"
)

// Update name and about of User
func UserRenameMe(store *Store, u mestoTypes.User, ua mestoTypes.UserMe) error {
	query := `UPDATE users SET name=$1, about=$2 WHERE id=$3`

	_, err := store.DB.Exec(query, ua.Name, ua.About, u.ID)
	if err != nil {
		return err
	}
	return nil
}

// Update User's avatar
func UserUpdateAvatar(store *Store, u mestoTypes.User, ua mestoTypes.UserAvatar) error  {
	query := `UPDATE users SET avatar=$1 WHERE id=$2`

	_, err := store.DB.Exec(query, ua.Avatar, u.ID)
	if err != nil {
		return err
	}
    return nil
}

// Get one user by password
func UserGetByPassword(store *Store, u mestoTypes.UserAuth) (*mestoTypes.User, error) {
	query := `SELECT id, name, about, avatar, cohort FROM users WHERE password=$1`

	password, err := uuid.Parse(u.Authorization)
	if err != nil {
		return nil, err
	}

	user, err := UserGet(store, query, password)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// Get one user by id
func UserGetByID(store *Store, id int) (*mestoTypes.User, error) {
	query := `SELECT id, name, about, avatar, cohort FROM users WHERE id=$1`
	user, err := UserGet(store, query, id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// Get one user by SQL query
func UserGet(store *Store, query string, args ...any) (*mestoTypes.User, error) {
	user := mestoTypes.User{}
	err := store.DB.QueryRow(query, args...).Scan(&user.ID, &user.Name, &user.About, &user.Avatar, &user.Cohort)
	if err != nil {
		if err != sql.ErrNoRows {
			log.Println(err.Error())
		}
		return nil, err
	}
	return &user, nil
}
