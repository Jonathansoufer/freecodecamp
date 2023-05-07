package utils

import (
	"encoding/json"
	"net/http"
	"io/ioutil"
)

func ParseBody(r *http.Request, x interface{}) error {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, x)
	if err != nil {
		return err
	}
	return nil
}

