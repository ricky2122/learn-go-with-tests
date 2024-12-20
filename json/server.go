package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const jsonContentType = "application/json"

type PlayerStore interface {
	GetPlayerScore(name string) int
	RecordWin(name string)
	GetLeague() []Player
}

type PlayerServer struct {
	store PlayerStore
	http.Handler
}

func NewPlayerServer(store PlayerStore) *PlayerServer {
	p := new(PlayerServer)

	p.store = store

	router := http.NewServeMux()
	router.Handle("GET /league", http.HandlerFunc(p.leagueHandler))
	router.Handle("GET /players/{name}", http.HandlerFunc(p.getPlayersHandler))
	router.Handle("POST /players/{name}", http.HandlerFunc(p.postPlayerHandler))

	p.Handler = router

	return p
}

func (p *PlayerServer) leagueHandler(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("content-type", jsonContentType)
	_ = json.NewEncoder(w).Encode(p.store.GetLeague())
	w.WriteHeader(http.StatusOK)
}

func (p *PlayerServer) getPlayersHandler(w http.ResponseWriter, r *http.Request) {
	player := r.PathValue("name")
	p.showScore(w, player)
}

func (p *PlayerServer) postPlayerHandler(w http.ResponseWriter, r *http.Request) {
	player := r.PathValue("name")
	p.processWin(w, player)
}

func (p *PlayerServer) showScore(w http.ResponseWriter, player string) {
	score := p.store.GetPlayerScore(player)

	if score == 0 {
		w.WriteHeader(http.StatusNotFound)
	}

	fmt.Fprint(w, score)
}

func (p *PlayerServer) processWin(w http.ResponseWriter, player string) {
	p.store.RecordWin(player)
	w.WriteHeader(http.StatusAccepted)
}
