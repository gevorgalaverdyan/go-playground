package db

import (
	"fmt"
	"sync"
)

type InMemDB struct {
	mu sync.Mutex
	store map[string]string
}

var instance *InMemDB
var once sync.Once

func NewMemDB() *InMemDB {
	once.Do(func() {
		instance = &InMemDB{
			store: make(map[string]string),
		}
	})

	return instance
}

func (db *InMemDB) LoadFromFile(filename string){
	db.mu.Lock()
	defer db.mu.Unlock()

	// file
}

func (db *InMemDB) List() {
	db.mu.Lock()
	defer db.mu.Unlock()
	fmt.Println("Database contents:")
	for key, value := range db.store {
		fmt.Printf("%s = %s\n", key, value)
	}
}