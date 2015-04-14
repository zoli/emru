package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"

	log "github.com/limetext/log4go"
	"github.com/zoli/emru/list"
)

type ListHandler struct {
	ls   list.Lists
	req  *http.Request
	data []byte
}

func (h *ListHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.req = r
	if err := h.parseReq(); err != nil {
		log.Error(err)
		http.NotFound(w, r)
	}

	w.Header().Set("Content-Type", "application/json")
	if _, err := w.Write(h.data); err != nil {
		log.Error(err)
	}
}

func (h *ListHandler) parseReq() error {
	log.Debug("Parsing %s", h.req.URL.Path)
	url := strings.TrimRight(h.req.URL.Path, "/")

	if url[:6] != "/lists" {
		return errors.New("invalid request")
	}
	if len(url) == 6 {
		switch h.req.Method {
		case "GET":
			h.lists()
		case "POST":
			h.newList()
		default:
			return errors.New("undefined method")
		}
		return nil
	}

	path := strings.Split(url[6:], "/")
	name := path[0]
	l, exist := h.ls[name]
	if !exist {
		return errors.New("list " + name + " not found")
	}
	if len(path) == 1 {
		switch h.req.Method {
		case "GET":
			return h.list(name)
		case "DELETE":
			return h.deleteList(name)
		default:
			return errors.New("undefined method")
		}
	}

	if path[1] != "tasks" {
		return errors.New("invalid request")
	}
	if len(path) == 2 {
		switch h.req.Method {
		case "GET":
			return h.tasks(l)
		case "POST":
			return h.newTask(l)
		default:
			return errors.New("undefined method")
		}
	}

	if len(path) > 3 {
		return errors.New("invalid request")
	}
	id, err := strconv.Atoi(path[2])
	if err != nil {
		return errors.New("invalid task id")
	}
	switch h.req.Method {
	case "GET":
		return h.task(l, id)
	case "PUT":
		return h.updateTask(l, id)
	case "DELETE":
		return h.deleteTask(l, id)
	default:
		return errors.New("undefined method")
	}
}

func (h *ListHandler) lists() (err error) {
	h.data, err = json.Marshal(h.ls)
	return err
}

func (h *ListHandler) list(name string) (err error) {
	h.data, err = json.Marshal(h.ls[name])
	return
}

func (h *ListHandler) newList() (err error) {
	decoder := json.NewDecoder(h.req.Body)
	var nList struct {
		name string
		lst  list.List
	}
	err = decoder.Decode(&nList)
	if err != nil {
		return
	}

	if _, exist := h.ls[nList.name]; exist {
		return errors.New("this name currently exists")
	}
	h.ls[nList.name] = &nList.lst
	return
}

func (h *ListHandler) deleteList(name string) (err error) {
	if _, exist := h.ls[name]; !exist {
		return errors.New("list doesn't exist")
	}
	delete(h.ls, name)
	return
}

func (h *ListHandler) tasks(l *list.List) (err error) {
	h.data, err = json.Marshal(l.Tasks())
	return
}

func (h *ListHandler) task(l *list.List, id int) (err error) {
	task, err := l.Get(id)
	h.data, err = json.Marshal(task)
	return
}

func (h *ListHandler) newTask(l *list.List) (err error) {
	return
}

func (h *ListHandler) updateTask(l *list.List, id int) (err error) {
	return
}

func (h *ListHandler) deleteTask(l *list.List, id int) (err error) {
	return
}