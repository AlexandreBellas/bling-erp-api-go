package bling

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

// Error represents an error response from the API.
type Error struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
}

// errorHandler is a function that checks the response for errors.
func errorHandler(r *http.Response) error {
	switch r.StatusCode {
	case http.StatusOK, http.StatusCreated, http.StatusAccepted, http.StatusNoContent, http.StatusNotModified:
		return nil
	}

	data, err := io.ReadAll(r.Body)
	if err == nil && data != nil {
		raw := new(Error)
		if err := json.Unmarshal(data, raw); err != nil {
			return err
		}
		return errors.New(raw.Msg)
	}
	return err
}
