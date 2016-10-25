package main

import (
	"sync"
)

type DB struct {
	data  map[string]string
	mutex sync.RWMutex
}

func NewDb() *DB {
	return &DB{
		data: make(map[string]string),
	}
}

func (db *DB) Get(key string) string {
	db.mutex.RLock()
	defer db.mutex.RUnlock()

	return db.data[key]
}

func (db *DB) Put(key, value string) {
	db.mutex.RLock()
	defer db.mutex.RUnlock()

	db.data[key] = value
	return
}
