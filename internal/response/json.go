package response

import (
	"encoding/json"
	"net/http"
)

func JSON(w http.ResponseWriter, status int, data any, pagination ...any) error {
	type envelope struct {
		Data       any `json:"data"`
		Pagination any `json:"pagination,omitempty,omitzero"`
	}
	var pag any
	if len(pagination) > 0 {
		pag = pagination[0]
	}
	return JSONWithHeaders(w, status, &envelope{Data: data, Pagination: pag}, nil)
}

func JSONError(w http.ResponseWriter, status int, data any, headers http.Header) error {
	type envelope struct {
		Error any `json:"error"`
		Code  int `json:"code"`
	}
	return JSONWithHeaders(w, status, &envelope{Error: data, Code: status}, headers)
}

func JSONWithHeaders(w http.ResponseWriter, status int, data any, headers http.Header) error {
	js, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return err
	}

	js = append(js, '\n')

	for key, value := range headers {
		w.Header()[key] = value
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(js)

	return nil
}
