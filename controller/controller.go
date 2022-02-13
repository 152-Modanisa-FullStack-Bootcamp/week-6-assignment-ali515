package controller

import (
	"bootcamp/handler"
	"net/http"
	"strings"
)

type IController interface {
	Routing(w http.ResponseWriter, r *http.Request)
}

type controller struct {
	handler handler.IHandler
}

/*
type ControllerError struct {
	Message string
}

func (e *ControllerError) Error() string {
	return e.Message
}
*/

func (c *controller) Routing(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		c.handler.GetAll(w, r)
		return
	}
	paths := strings.Split(r.URL.Path, "/")
	if len(paths) != 2 {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	switch r.Method {
	case http.MethodGet:
		c.handler.Get(w, r)
		return
	case http.MethodPost:
		c.handler.Update(w, r)
		return
	case http.MethodPut:
		c.handler.Create(w, r)
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
}

func NewController(handler handler.IHandler) IController {
	return &controller{handler: handler}
}
