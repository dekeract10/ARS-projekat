package main

import (
	"errors"
	"github.com/gorilla/mux"
	"mime"
	"net/http"
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

func (ts *Service) delConfigHandler(writer http.ResponseWriter, request *http.Request) {
	id := mux.Vars(request)["id"]
	if value, ok := ts.data[id]; ok && len(value) == 1 {
		delete(ts.data, id)
		renderJSON(writer, value)
	} else {
		err := errors.New("key not found")
		http.Error(writer, err.Error(), http.StatusNotFound)
	}
}

func (ts *Service) delGroupHandler(writer http.ResponseWriter, request *http.Request) {
	id := mux.Vars(request)["id"]
	if value, ok := ts.data[id]; ok && len(value) > 1 {
		delete(ts.data, id)
		renderJSON(writer, value)
	} else {
		err := errors.New("key not found")
		http.Error(writer, err.Error(), http.StatusNotFound)
	}
}
