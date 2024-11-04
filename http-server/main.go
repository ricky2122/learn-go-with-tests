package main

import (
	"log"
	"net/http"
)

type InMemoryPlayerStore struct{}

func (i *InMemoryPlayerStore) GetPlayerScore(_ string) int {
	return 123
}

func (i *InMemoryPlayerStore) RecordWin(_ string) {}

func main() {
	server := &PlayerServer{store: &InMemoryPlayerStore{}}
	log.Fatal(http.ListenAndServe(":5000", server))
}
