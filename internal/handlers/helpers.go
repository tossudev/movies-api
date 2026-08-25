package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
)

type Pagination struct {
	page   int
	offset int
}

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

func getPagination(r *http.Request) (page int, size int, err error) {
	queries := r.URL.Query()
	pageStr, sizeStr := queries.Get("page"), queries.Get("size")

	// Catch existence
	if len(pageStr) == 0 && len(sizeStr) != 0 || len(sizeStr) == 0 && len(sizeStr) != 0 {
		return -1, -1, errors.New("bad request")
	} else if len(pageStr) == 0 && len(sizeStr) == 0 {
		return 0, 0, nil
	}

	page, err = strconv.Atoi(pageStr)
	if err != nil {
		return -1, -1, errors.New("invalid pagination: poor page value")
	}

	size, err = strconv.Atoi(sizeStr)
	if err != nil {
		return -1, -1, errors.New("invalid pagination: poor size value")
	}

	// Catch invalids
	if size < 1 || page < 0 {
		return -1, -1, errors.New("invalid pagination: too small value")
	}

	return page, size, nil
}
