package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

func ParseBody(r *http.Request, m interface{}){
	err := json.NewDecoder(r.Body).Decode(&m)
	if err != nil {
		log.Fatal(err)
		return
	}
}