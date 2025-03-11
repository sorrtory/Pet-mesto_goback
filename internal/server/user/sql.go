package user

import (
	"database/sql"
	"log"
	"mesto-goback/internal/db"

	"github.com/google/uuid"
)

// TODO: implement
func GetMe(store *db.Store, u UserMe) (*User, error){
    query := `SELECT * FROM USERS WHERE name = ?;
    `
    r, err := store.Query(query, u.Name)
    if err != nil {
		return nil, err
	}
    user := &User{}
    for r.Next() {
    }
    return user, nil
}

func GetUser(store *db.Store, u UserAuth) (*User, error){
    query := `SELECT id, name, about, avatar FROM users WHERE password=$1`
    user := User{}

    id, err := uuid.Parse(u.Authorization)
    if err != nil {
        return nil, err
    }

    err = store.DB.QueryRow(query, id).Scan(&user.ID, &user.Name, &user.About, &user.Avatar)

    if err != nil {
        if err != sql.ErrNoRows {
            log.Println(err.Error())
        }
        return nil, err
    }
    return &user, nil
}
