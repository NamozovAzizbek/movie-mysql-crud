package utils

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func ParseBody(r *http.Request, m interface{}) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
		return
	}

	if err = json.Unmarshal([]byte(body), m); err != nil {
		log.Fatal(err)
		return
	}
}
