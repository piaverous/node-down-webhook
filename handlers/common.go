package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type responseJSON struct {
	Status  int
	Message string
}

func asJson(w http.ResponseWriter, status int, message string) {
	data := responseJSON{
		Status:  status,
		Message: message,
	}
	bytes, _ := json.Marshal(data)
	json := string(bytes[:])

	w.WriteHeader(status)
	fmt.Fprint(w, json)
}
