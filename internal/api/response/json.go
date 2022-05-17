package response

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type Json struct {
}

func NewJsonResponse() Response {
	return &Json{}
}

func (j *Json) Success(w http.ResponseWriter, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Status", strconv.Itoa(http.StatusOK))

	if payload != nil {
		resp, err := json.Marshal(payload)
		if err != nil {
			//@todo add logger
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		_, _ = w.Write(resp)
	}
}

func (j *Json) Redirect(w http.ResponseWriter, req *http.Request, payload interface{}) {
	http.Redirect(w, req, payload.(string), http.StatusMovedPermanently)
}

func (j *Json) Error(w http.ResponseWriter, message string, errorCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Status", "404 NOT FOUND")
	resp, err := json.Marshal(map[string]string{
		"message": message,
		"code":    strconv.Itoa(errorCode),
	})
	if err != nil {
		//@todo add logger
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, _ = w.Write(resp)
}
