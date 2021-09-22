package main

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"sync"

	_ "github.com/mattn/go-sqlite3"
)

type Record struct {
	Email    string
	Password string
	NickName string
}

type DBInterface interface {
	AddRecord(record Record)
	Select(request string) (Record, error)
	// RemoveRecord(record Record)
	IsEmailUnique(string) bool
	Close()
}

type InmemoryDB struct {
	data map[string]Record
	m    sync.Mutex
}

func NewDB() DBInterface {
	return &InmemoryDB{
		data: make(map[string]Record),
	}
}

func (db *InmemoryDB) IsEmailUnique(email string) bool {
	db.m.Lock()
	defer db.m.Unlock()
	for _, record := range db.data {
		if record.Email == email {
			return false
		}
	}

	return true
}

func (db *InmemoryDB) AddRecord(record Record) {
	db.m.Lock()
	defer db.m.Unlock()
	db.data[record.Email] = record
	fmt.Println("add record", db.data)
}

func (db *InmemoryDB) Close() {
}

func (db *InmemoryDB) Select(request string) (Record, error) {
	db.m.Lock()
	defer db.m.Unlock()
	fmt.Println("select", db.data)
	if value, ok := db.data[request]; ok {
		return value, nil
	}
	return Record{}, errors.New("invalid reuqest")
}

type DataBase struct {
	db *sql.DB
	// m  sync.RWMutex
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
	_, err := db.db.Exec("insert into "+tableName+" (email, password, nickName) values ($1, $2, $3)",
		strings.Replace(record.Email, "@", "a", -1), record.Password, record.NickName)
	if err != nil {
		fmt.Println("AddRecord error", err)
	}
}

func (db *DataBase) Select(email string) (Record, error) {
	record := Record{}
	err := db.db.QueryRow("select * from users where email=$1", strings.Replace(email, "@", "a", -1)).Scan(&record.Email, &record.Password, &record.NickName)
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
	row := db.db.QueryRow("select * from users where email=$1", strings.Replace(email, "@", "a", -1))
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
