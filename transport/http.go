package transport

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/tiennv147/mazti-commons/endpoints"
	"github.com/tiennv147/mazti-commons/errors"
)

func codeFrom(err error) int {
	switch err {
	case errors.ErrNotFound:
		return http.StatusNotFound
	case errors.ErrAlreadyExists, errors.ErrInconsistentIDs:
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}
}

func EncodeError(_ context.Context, err error, w http.ResponseWriter) {
	if err == nil {
		panic("encodeError with nil error")
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(codeFrom(err))
	_ = json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}

func EncodeCommonResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	if response != nil {
		return json.NewEncoder(w).Encode(response)
	}
	return nil
}

func DecodeRequestCommonWithID(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	idStr, ok := vars["id"]
	if !ok {
		return nil, errors.ErrBadRouting
	}
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return nil, errors.ErrInconsistentIDs
	}
	return endpoints.RequestWithID{ID: id}, nil
}

func DecodeListCommonRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req = endpoints.ListRequest{
		Offset: 0,
		Limit:  10,
	}
	values := r.URL.Query()

	oStr := values.Get("offset")
	oVal, oErr := strconv.Atoi(oStr)
	if oErr == nil {
		req.Offset = oVal
	}

	lStr := values.Get("limit")
	lVal, lErr := strconv.Atoi(lStr)
	if lErr == nil {
		req.Limit = lVal
	}

	return req, nil
}

func EncodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	if response != nil {
		return json.NewEncoder(w).Encode(response)
	}
	return nil
}
