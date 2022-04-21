package main

import (
	"errors"
	"log"
	"mime"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type Service struct {
	data map[string][]*Config // izigrava bazu podataka
}

func (ts *Service) createConfigHandler(w http.ResponseWriter, req *http.Request) {
	contentType := req.Header.Get("Content-Type")
	mediatype, _, err := mime.ParseMediaType(contentType)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if mediatype != "application/json" {
		err := errors.New("Expect application/json Content-Type")
		http.Error(w, err.Error(), http.StatusUnsupportedMediaType)
		return
	}

	rt, err := decodeBody(req.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id := createId()
	ts.data[id] = rt
	w.Write([]byte(id))
}

func (ts *Service) getAllConfigsHandler(w http.ResponseWriter, req *http.Request) {
	allTasks := [][]*Config{}
	for _, v := range ts.data {
		if len(v) == 1 {
			allTasks = append(allTasks, v)
		}
	}
	renderJSON(w, allTasks)
}

func (ts *Service) getAllGroupssHandler(w http.ResponseWriter, req *http.Request) {
	allTasks := [][]*Config{}
	for _, v := range ts.data {
		if len(v) > 1 {
			allTasks = append(allTasks, v)
		}
	}
	renderJSON(w, allTasks)
}

func (ts *Service) getConfigHandler(w http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	task, ok := ts.data[id]
	if !ok || len(task) > 1 {
		err := errors.New("key not found")
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	renderJSON(w, task)
}

func (ts *Service) getGroupHandler(w http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	task, ok := ts.data[id]
	if !ok || len(task) == 1 {
		err := errors.New("key not found")
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	renderJSON(w, task)
}

func (ts *Service) putConfigHandler(w http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	task, ok := ts.data[id]

	if !ok || len(task) == 1 {
		err := errors.New("key not found")
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	rt, err := decodeBody(req.Body)
	if len(rt) > 1 {
		err := errors.New("Recived invalid JSON format! (confg length > 1)")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err == nil {
		ts.data[id] = append(task, rt[0])
	}
}

func (ts *Service) delConfigHandler(w http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	if v, ok := ts.data[id]; ok || len(v) > 1 {
		delete(ts.data, id)
		renderJSON(w, v)
	} else {
		err := errors.New("key not found")
		http.Error(w, err.Error(), http.StatusNotFound)
	}
}

func (ts *Service) hangHandler(w http.ResponseWriter, req *http.Request) {
	seconds := mux.Vars(req)["seconds"]
	intSeconds, err := strconv.Atoi(seconds)
	if err != nil {
		http.Error(w, "could not convert to seconds", http.StatusBadRequest)
		return
	}

	log.Println("started hanging")
	time.Sleep(time.Second * time.Duration(intSeconds))
}
