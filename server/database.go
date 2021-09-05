package main

import (
	"errors"
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
}

type InmemoryDB struct {
	data map[string]Record
}

func NewDB() DBInterface {
	return &InmemoryDB{
		data: make(map[string]Record),
	}
}

// TODO: add mutex
func (db *InmemoryDB) AddRecord(record Record) {
	db.data[record.IP] = record
}

// TODO: add mutex
func (db *InmemoryDB) Select(request string) (Record, error) {
	if value, ok := db.data[request]; ok {
		return value, nil
	}
	return Record{}, errors.New("invalid reuqest")
}
