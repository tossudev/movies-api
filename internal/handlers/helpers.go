package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
)

func decodeRequest[T any](r *http.Request) (req T, err error) {
	return req, json.NewDecoder(r.Body).Decode(&req)
}

func getID(r *http.Request) (int, error) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id <= 0 {
		return -1, errors.New("Invalid ID")
	}

	return id, nil
}
