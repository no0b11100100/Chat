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
		record.Email, record.Password, record.NickName)
	if err != nil {
		panic(err)
	}
}

var CannotFind = errors.New("can not find")

func (db *DataBase) Select(request string) (Record, error) {
	rows, err := db.db.Query("select * from Products where email=" + request)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	record := Record{}
	err = CannotFind
	if rows.Next() {
		err = rows.Scan(&record.Email, &record.Password, &record.NickName)
		if err == nil {
			return record, nil
		}
	}

	return Record{}, err
}

func (db *DataBase) IsEmailUnique(email string) bool {
	_, err := db.Select(email)
	return err == CannotFind
}
