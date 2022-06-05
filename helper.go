package main

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	cs "github.com/dekeract10/ARS-projekat/configstore"
	tracer "github.com/dekeract10/ARS-projekat/tracer"
	"github.com/google/uuid"
)

func decodeConfigBody(ctx context.Context, r io.Reader) (*cs.Config, error) {
	span := tracer.StartSpanFromContext(ctx, "decodeConfigBody")
	defer span.Finish()

	dec := json.NewDecoder(r)
	dec.DisallowUnknownFields()

	var config *cs.Config
	if err := dec.Decode(&config); err != nil {
		tracer.LogError(span, err)
		return nil, err
	}
	return config, nil
}

func decodeGroupBody(r io.Reader) (*cs.Group, error) {
	dec := json.NewDecoder(r)
	dec.DisallowUnknownFields()

	var group *cs.Group
	if err := dec.Decode(&group); err != nil {
		return nil, err
	}
	return group, nil
}

func renderJSON(w http.ResponseWriter, v interface{}, id string) {
	js, err := json.Marshal(v)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)

}

func createId() string {
	return uuid.New().String()
}
