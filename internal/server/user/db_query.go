package user

import "mesto-goback/internal/db"


// TODO: implement
func GetMe(store *db.Store, u UserMe) (*User, error){
    query := `\
    SELECT * FROM USERS WHERE name = ?;
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

// TODO: implement
func GetUser(store *db.Store, u UserAuth) (*User, error){
    query := `\
    SELECT * FROM SECRETS WHERE token = ?;
    `
    r, err := store.Query(query, u.Authorization)
    if err != nil {
		return nil, err
	}
    defer r.Close()

    user := &User{}
    return user, nil


}
