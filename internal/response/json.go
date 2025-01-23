package response

import (
	"encoding/json"
	"net/http"
)

func JSON(w http.ResponseWriter, status int, data any) error {
	type envelope struct {
		Data any `json:"data"`
	}
	return JSONWithHeaders(w, status, &envelope{Data: data}, nil)
}

func JSONError(w http.ResponseWriter, status int, data any, headers http.Header) error {
	type envelope struct {
		Error any `json:"error"`
	}
	return JSONWithHeaders(w, status, &envelope{Error: data}, headers)
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
