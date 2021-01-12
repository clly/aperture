package helloworld

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func HelloWorld() http.Handler {
	return &hello{
		msg: "Hello!",
	}
}

type hello struct {
	msg string
}

func (h *hello) Get(r *http.Request) (statusCode int, content []byte) {
	return http.StatusOK, []byte(h.msg)
}

func (h *hello) Post(r *http.Request) (statusCode int, content []byte) {
	if r.ContentLength > 100 {
		return http.StatusRequestEntityTooLarge, []byte{}
	}
	body, err := r.GetBody()
	if err != nil {
		return http.StatusInternalServerError, []byte{}
	}
	msg, err := ioutil.ReadAll(body)
	if err != nil {
		return http.StatusInternalServerError, []byte{}
	}

	h.msg = string(msg)
	return http.StatusOK, []byte{}
}

func (h *hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var code int
	var resp []byte
	switch r.Method {
	case http.MethodGet:
		code, resp = h.Get(r)
	case http.MethodPost:
		code, resp = h.Post(r)
	default:
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(code)
	_, err := w.Write(resp)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}
