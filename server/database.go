package main

import (
	"database/sql"
	"errors"
	"sync"

	_ "github.com/mattn/go-sqlite3"
)

type Record struct {
	IP       string
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
	db.data[record.IP] = record
}

func (db *InmemoryDB) Close() {
}

func (db *InmemoryDB) Select(request string) (Record, error) {
	db.m.Lock()
	defer db.m.Unlock()
	if value, ok := db.data[request]; ok {
		return value, nil
	}
	return Record{}, errors.New("invalid reuqest")
}

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
	_, err := db.db.Exec("insert into "+tableName+" (IP, Email, Password, NickName) values ($1, $2, $3, $4)",
		record.IP, record.Email, record.Password, record.NickName)
	if err != nil {
		panic(err)
	}
}

func (db *DataBase) Select(request string) (Record, error) {
	return Record{}, errors.New("invalid reuqest")
}

func (db *DataBase) IsEmailUnique(email string) bool {
	return false
}
