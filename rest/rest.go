package rest

import (
	"encoding/json"
	"net/http"
)

type RestMessage struct {
	Status  string      `json:"status"`
	Message interface{} `json:"message"`
}

func MustEncode(w http.ResponseWriter, i interface{}) {
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Content-type", "application/json;charset=utf-8")
	e := json.NewEncoder(w)
	if err := e.Encode(i); err != nil {
		//panic(err)
		e.Encode(err.Error())
	}
}

func MustResEncode(w http.ResponseWriter, r *http.Request, i interface{}) {
	requestID := r.Context().Value("request_id")
	w.Header().Set("RequestID", requestID.(string))
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Content-type", "application/json;charset=utf-8")

	e := json.NewEncoder(w)
	if err := e.Encode(i); err != nil {
		//panic(err)
		e.Encode(err.Error())
	}
}
