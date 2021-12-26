package database

import (
	"database/sql"
	"fmt"
	"sync"

	_ "github.com/mattn/go-sqlite3"
)

type DataBase struct {
	db *sql.DB
	m  sync.RWMutex
}

const tableName = "users"

func NewDataBase() *DataBase {
	db, err := sql.Open("sqlite3", tableName+".db")
	if err != nil {
		panic(err)
	}
	return &DataBase{
		db: db,
	}
}

func (db *DataBase) Close() {
	db.db.Close()
}

// https://metanit.com/go/tutorial/10.4.php
func (db *DataBase) AddRecord(record Record) {
	db.m.Lock()
	defer db.m.Unlock()
	_, err := db.db.Exec("insert into "+tableName+" (email, password, nickName) values ($1, $2, $3)",
		record.Email, record.Password, record.NickName)
	if err != nil {
		fmt.Println("AddRecord error", err)
	}
}

func (db *DataBase) Select(email string) (Record, error) {
	db.m.RLock()
	defer db.m.RUnlock()
	record := Record{}
	err := db.db.QueryRow("select * from users where email=$1", email).Scan(&record.Email, &record.Password, &record.NickName)
	switch err {
	case sql.ErrNoRows:
		return Record{}, err
	case nil:
		return record, nil
	default:
		fmt.Println("Select", err)
		return Record{}, err
	}
}

func (db *DataBase) IsEmailUnique(email string) bool {
	db.m.RLock()
	defer db.m.RUnlock()
	row := db.db.QueryRow("select * from users where email=$1", email)
	record := Record{}
	switch err := row.Scan(&record.Email); err {
	case sql.ErrNoRows:
		return true
	case nil:
		return false
	default:
		fmt.Println("IsEmailUnique", err)
		return false
	}
}
