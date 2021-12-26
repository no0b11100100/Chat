package database

import (
	"errors"
	"fmt"
	"sync"
)

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
