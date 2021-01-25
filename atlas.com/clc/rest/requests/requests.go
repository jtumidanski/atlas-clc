package requests

import (
	"atlas-clc/rest/attributes"
	"bytes"
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

func Get(l *log.Logger, url string, resp interface{}) error {
	r, err := http.Get(url)
	if err != nil {
		l.Printf("[ERROR] dispatching [GET] to %s", url)
		return errors.New("error dispatching get to url")
	}

	err = ProcessResponse(l, r, resp)
	return err
}

func Post(l *log.Logger, url string, input interface{}) (*http.Response, error) {
	jsonReq, err := json.Marshal(input)
	if err != nil {
		l.Println("[ERROR] marshalling post body.")
		return nil, errors.New("error marshalling post body")
	}

	r, err := http.Post(url, "application/json; charset=utf-8", bytes.NewReader(jsonReq))
	if err != nil {
		l.Printf("[ERROR] dispatching [POST] to %s", url)
		return nil, errors.New("error dispatching post to url")
	}
	return r, nil
}

func ProcessResponse(l *log.Logger, r *http.Response, rb interface{}) error {
	err := attributes.FromJSON(rb, r.Body)
	if err != nil {
		l.Printf("[ERROR] decoding response")
		return err
	}

	return nil
}

func ProcessErrorResponse(l *log.Logger, r *http.Response, eb interface{}) error {
	if r.ContentLength > 0 {
		err := attributes.FromJSON(eb, r.Body)
		if err != nil {
			l.Printf("[ERROR] decoding error response")
			return errors.New("error decoding error response")
		}
		return nil
	} else {
		return nil
	}
}
