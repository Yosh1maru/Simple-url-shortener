package response

import "net/http"

type Response interface {
	Success(w http.ResponseWriter, payload interface{})
	Redirect(w http.ResponseWriter, req *http.Request, payload interface{})
	Error(w http.ResponseWriter, message string, errorCode int)
}
