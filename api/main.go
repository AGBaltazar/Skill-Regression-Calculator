package main

import {
	"time"
	"net/http"
}

type Skill struct {
	Name        string    `json:"skill"`
	Complexity  float64   `json:"complexity"` 
	LastWorked  time.Time `json:"lastWorked"`
}


func handleData(w http.ResponseWritter, r *http.Request){
	
}