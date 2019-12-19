package transport

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"gitlab.com/phypass_server/go/commons/dto"
)

func DecodeListRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := dto.ListRequest{
		Offset: 0,
		Limit:  10,
	}
	query := r.URL.Query()

	oStr := query.Get("offset")
	oVal, oErr := strconv.Atoi(oStr)
	if oErr == nil {
		req.Offset = oVal
	}

	lStr := query.Get("limit")
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
