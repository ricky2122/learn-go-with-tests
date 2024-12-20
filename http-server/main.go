package main

import (
	"log"
	"net/http"
	"sync"
)

type InMemoryPlayerStore struct {
	store map[string]int
	mu    sync.Mutex
}

func NewInMemoryPlayerStore() *InMemoryPlayerStore {
	return &InMemoryPlayerStore{
		store: map[string]int{},
	}
}

func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return i.store[name]
}

func (i *InMemoryPlayerStore) RecordWin(name string) {
	i.mu.Lock()
	defer i.mu.Unlock()
	i.store[name]++
}

func main() {
	server := &PlayerServer{store: NewInMemoryPlayerStore()}
	log.Fatal(http.ListenAndServe(":5000", server))
}
