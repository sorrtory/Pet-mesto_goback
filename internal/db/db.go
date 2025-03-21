package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type Store struct {
    DB *sql.DB
}

func NewConnection(host string, user string, password string) (*Store, error){
    connStr := fmt.Sprintf("host=%s user=%s dbname=postgres password=%s sslmode=disable", host, user, password)
	db, err := sql.Open("postgres", connStr)
    if err != nil {
        return nil, err
    }
    if err := db.Ping(); err != nil {
		return nil, err
	}
    return &Store{db}, nil
}


// func (s Store) Read()  {
//     r, err := s.db.Query("SELECT * FROM users5")
//     if err != nil {
//         return
//     }
//     fmt.Println(r.Columns())
// }


func (s Store) Query(query string, args ...any) (*sql.Rows, error){
    r, err := s.DB.Query(query, args...)
    if err != nil {
        log.Printf("Can't query: %v\n", query)
        return nil, err
    }
    return r, nil
}
