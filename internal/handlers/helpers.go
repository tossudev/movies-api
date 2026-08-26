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
		return -1, errors.New("invalid id")
	}

	return id, nil
}

func getPagination(r *http.Request) (page int, size int, err error) {
	queries := r.URL.Query()
	pageStr, sizeStr := queries.Get("page"), queries.Get("size")

	// Catch existence
	if (pageStr == "") != (sizeStr == "") {
		return -1, -1, errors.New("page and size must be provided together")
	} else if len(pageStr) == 0 && len(sizeStr) == 0 {
		return 0, 0, nil
	}

	page, err = strconv.Atoi(pageStr)
	if err != nil {
		return -1, -1, errors.New("invalid page")
	}

	size, err = strconv.Atoi(sizeStr)
	if err != nil {
		return -1, -1, errors.New("invalid size")
	}

	// Catch invalids
	if page < 1 {
		return -1, -1, errors.New("page must be greater than zero")
	}

	if size < 1 {
		return -1, -1, errors.New("size must be greater than zero")
	}

	return page, size, nil
}
