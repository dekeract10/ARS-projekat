package main

type Service struct {
	Data map[string]*Config `json:"data"`
}

type Config struct {
	Entries map[string]string `json:"entries"`
}
