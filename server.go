package main

import (
	"errors"
	cs "github.com/dekeract10/ARS-projekat/configstore"
	"github.com/gorilla/mux"
	"mime"
	"net/http"
)

type Service struct {
	store *cs.ConfigStore
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

	rt, err := decodeConfigBody(req.Body)
	if err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	config, err := ts.store.CreateConfig(rt)
	// make version if it doesn't exist
	// id := createId()
	w.Write([]byte(config.ID))
}

func (ts *Service) putNewVersion(w http.ResponseWriter, req *http.Request) {
	contentType := req.Header.Get("Content-Type")
	mediatype, _, err := mime.ParseMediaType(contentType)
	id := mux.Vars(req)["id"]

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if mediatype != "application/json" {
		err := errors.New("Expect application/json Content-Type")
		http.Error(w, err.Error(), http.StatusUnsupportedMediaType)
		return
	}

	rt, err := decodeConfigBody(req.Body)
	if err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	rt.ID = id
	config, err := ts.store.UpdateConfigVersion(rt)

	if err != nil {
		http.Error(w, "Given config version already exists! ", http.StatusBadRequest)
		return
	}

	w.Write([]byte(config.ID))
}

func (ts *Service) getConfigHandler(w http.ResponseWriter, req *http.Request) {
	ver := mux.Vars(req)["ver"]
	id := mux.Vars(req)["id"]
	task, ok := ts.store.FindConf(id, ver)
	if ok != nil {
		err := errors.New("key not found")
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	renderJSON(w, task)
}

func (ts *Service) getConfigGroupsHandler(w http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	task, ok := ts.store.FindConfVersions(id)
	if ok != nil {
		err := errors.New("key not found")
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	renderJSON(w, task)
}

// func (ts *Service) createGroupHandler(w http.ResponseWriter, req *http.Request) {
// 	ver := mux.Vars(req)["ver"]

// 	contentType := req.Header.Get("Content-Type")
// 	mediatype, _, err := mime.ParseMediaType(contentType)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}

// 	if mediatype != "application/json" {
// 		err := errors.New("Expect application/json Content-Type")
// 		http.Error(w, err.Error(), http.StatusUnsupportedMediaType)
// 		return
// 	}

// 	rt, err := decodeBody(req.Body)
// 	if err != nil || len(rt) == 1 {
// 		http.Error(w, "Invalid JSON format (config len = 1)", http.StatusBadRequest)
// 		return
// 	}
// 	// make version if it doesn't exist
// 	if _, ok := ts.versions[ver]; !ok {
// 		ts.versions[ver] = make(map[string][]*Config)
// 	}

// 	id := createId()
// 	ts.versions[ver][id] = rt
// 	w.Write([]byte(id))
// }

// func (ts *Service) getAllConfigsHandler(w http.ResponseWriter, req *http.Request) {
// 	ver := mux.Vars(req)["ver"]
// 	allTasks := [][]*Config{}
// 	for _, v := range ts.versions[ver] {
// 		if len(v) == 1 {
// 			allTasks = append(allTasks, v)
// 		}
// 	}
// 	renderJSON(w, allTasks)
// }

// func (ts *Service) getAllGroupsHandler(w http.ResponseWriter, req *http.Request) {
// 	ver := mux.Vars(req)["ver"]
// 	allTasks := [][]*Config{}
// 	for _, v := range ts.versions[ver] {
// 		if len(v) > 1 {
// 			allTasks = append(allTasks, v)
// 		}
// 	}
// 	renderJSON(w, allTasks)
// }

// func (ts *Service) getGroupHandler(w http.ResponseWriter, req *http.Request) {
// 	id := mux.Vars(req)["id"]
// 	ver := mux.Vars(req)["ver"]
// 	task, ok := ts.versions[ver][id]
// 	if !ok || len(task) == 1 {
// 		err := errors.New("key not found")
// 		http.Error(w, err.Error(), http.StatusNotFound)
// 		return
// 	}
// 	renderJSON(w, task)
// }

// func (ts *Service) putConfigHandler(w http.ResponseWriter, req *http.Request) {
// 	id := mux.Vars(req)["id"]
// 	ver := mux.Vars(req)["ver"]
// 	task, ok := ts.versions[ver][id]

// 	if !ok || len(task) == 1 {
// 		err := errors.New("key not found")
// 		http.Error(w, err.Error(), http.StatusNotFound)
// 		return
// 	}

// 	rt, err := decodeBody(req.Body)
// 	if len(rt) > 1 {
// 		err := errors.New("Recived invalid JSON format! (confg length > 1)")
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}

// 	if err == nil {
// 		ts.versions[ver][id] = append(task, rt[0])
// 	}
// }

// func (ts *Service) delConfigHandler(writer http.ResponseWriter, request *http.Request) {
// 	id := mux.Vars(request)["id"]
// 	ver := mux.Vars(request)["ver"]
// 	if value, ok := ts.versions[ver][id]; ok && len(value) == 1 {
// 		delete(ts.versions[ver], id)
// 		renderJSON(writer, value)
// 	} else {
// 		err := errors.New("key not found")
// 		http.Error(writer, err.Error(), http.StatusNotFound)
// 	}
// }

// func (ts *Service) delGroupHandler(writer http.ResponseWriter, request *http.Request) {
// 	id := mux.Vars(request)["id"]
// 	ver := mux.Vars(request)["ver"]
// 	if value, ok := ts.versions[ver][id]; ok && len(value) > 1 {
// 		delete(ts.versions[ver], id)
// 		renderJSON(writer, value)
// 	} else {
// 		err := errors.New("key not found")
// 		http.Error(writer, err.Error(), http.StatusNotFound)
// 	}
// }
