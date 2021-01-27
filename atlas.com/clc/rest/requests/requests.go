package requests

import (
	"atlas-clc/rest/attributes"
	"bytes"
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

const (
	BaseRequest string = "http://atlas-nginx:80"
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

func Patch(l *log.Logger, url string, input interface{}) (*http.Response, error) {
	jsonReq, err := json.Marshal(input)
	if err != nil {
		l.Println("[ERROR] marshalling patch body.")
		return nil, errors.New("error marshalling patch body")
	}

	client := &http.Client{}
	r, err := http.NewRequest(http.MethodPatch, url, bytes.NewReader(jsonReq))
	if err != nil {
		l.Printf("[ERROR] dispatching [PATCH] to %s", url)
		return nil, errors.New("error dispatching patch to url")
	}
	r.Header.Set("Content-Type", "application/json")

	return client.Do(r)
}

func Delete(l *log.Logger, url string) (*http.Response, error) {
	client := &http.Client{}
	r, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		l.Printf("[ERROR] dispatching [DELETE] to %s", url)
		return nil, errors.New("error dispatching delete to url")
	}
	r.Header.Set("Content-Type", "application/json")

	return client.Do(r)
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
			l.Println(err)
			return errors.New("error decoding error response")
		}
		return nil
	} else {
		return nil
	}
}
