package main

import (
	"encoding/json"
	"net/http"
)

type input struct {
	Email string `json:"email"`
	Id    string `json:"id"`
}

func (srv *Server) Add(w http.ResponseWriter, r *http.Request) {
	var body input
	addErr := json.NewDecoder(r.Body).Decode(&body)
	if addErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	srv.CP.AddToCache([]byte(body.Id), []byte(body.Email))
	out := make([]input, 0)
	for k, e := range srv.CP.GetMap() {
		var o input
		o.Id = k
		o.Email = e.Data
		out = append(out, o)
	}
	data, err := json.Marshal(out)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	_, err = w.Write(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	return
}

func (srv *Server) Data(w http.ResponseWriter, r *http.Request) {
	var body input
	addErr := json.NewDecoder(r.Body).Decode(&body)
	if addErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	d, err := srv.CP.GetCache([]byte(body.Id))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	data, err := json.Marshal(d)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	_, err = w.Write(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	return
}
