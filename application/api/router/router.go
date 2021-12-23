package router

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type HandlerFunc interface {
	Process(w http.ResponseWriter, r *http.Request) (interface{}, error)
}

type ErrorHandlerFunc func(r *http.Request, w http.ResponseWriter, err error)

type handler struct {
	handler      HandlerFunc
	errorHandler ErrorHandlerFunc
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	res, err := h.handler.Process(w, r)
	if err != nil {
		h.errorHandler(r, w, err)
	} else {
		writeResponse(r, w, http.StatusOK, res)
	}
}

func NewHandler(h HandlerFunc) http.HandlerFunc {
	han := handler{
		handler:      h,
		errorHandler: DefaultErrorHandler,
	}
	return han.ServeHTTP
}

func NewHandlerWithErrorHandler(h HandlerFunc, errorHandler ErrorHandlerFunc) http.HandlerFunc {
	han := handler{
		handler:      h,
		errorHandler: errorHandler,
	}
	return han.ServeHTTP
}

func DefaultErrorHandler(r *http.Request, w http.ResponseWriter, err error) {
	// TODO log, response
	writeResponse(r, w, http.StatusInternalServerError, nil)
}

func ReadRequest(r *http.Request, out interface{}) error {
	requestBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(requestBytes, out)
	if err != nil {
		return err
	}
	return nil
}

func writeResponse(r *http.Request, w http.ResponseWriter, code int, data interface{}) {
	var responseBytes []byte
	if data != nil {
		w.Header().Set("Content-Type", "application/json")
		responseBytes, _ = json.Marshal(data)
	}
	w.WriteHeader(code)
	if len(responseBytes) > 0 {
		_, _ = w.Write(responseBytes)
	}
}
